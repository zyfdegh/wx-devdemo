package tokensvc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	defaultAPIBaseURL = "https://api.weixin.qq.com/cgi-bin/token"
	defaultGrantType  = "client_credential"
	// refresh token every 7190s(a little less than 2 hour) by default
	defaultPullingSec = 7190
)

var (
	// ErrAlreadyRunning returned when token daemon is already running
	ErrDaemonAlreadyRunning = errors.New("token daemon is already running")
	// ErrEmptyToken returned when token is empty
	ErrEmptyToken = errors.New("token empty")
	// ErrFetchToken returned when access_token fetch failed
	ErrFetchToken = errors.New("fetch access token failed")
)

type AccessToken struct {
	Token    string `json:"access_token"`
	ExpireIn int    `json:"expires_in"`
}

type TokenDaemon struct {
	IsRunning bool
	token     string
	ticker    *time.Ticker
	Config    DaemonConfig
	// Start() error
	// GetToken() string
	// RefreshToken() error
	// Stop() error
}

type DaemonConfig struct {
	AppID  string
	Secret string
	// optinal
	PollingSec int
	// optinal
	APIBaseURL string
	// optinal
	GrantType string
}

func NewTokenDaemon(config DaemonConfig) (daemon *TokenDaemon, err error) {
	daemon = &TokenDaemon{}
	daemon.Config = config
	return
}

func (d *TokenDaemon) Start() error {
	if d.IsRunning {
		return ErrDaemonAlreadyRunning
	}
	return d.start()
}

// start timer and continiously call API then update token
func (d *TokenDaemon) start() error {
	// get token immediately at the beginning
	d.refreshToken()

	pollingSec := defaultPullingSec
	if d.Config.PollingSec > 0 {
		pollingSec = d.Config.PollingSec
	}

	log.Printf("will refresh access token every %d sec\n", pollingSec)

	d.ticker = time.NewTicker(time.Duration(pollingSec) * time.Second)
	go func() {
		for range d.ticker.C {
			d.refreshToken()
		}
	}()

	d.IsRunning = true
	return nil
}

// clear token
// stop timer
func (d *TokenDaemon) Stop() error {
	d.token = ""
	d.IsRunning = false
	d.ticker.Stop()
	return nil
}

func (d *TokenDaemon) RefreshToken() error {
	return d.refreshToken()
}

func (d *TokenDaemon) refreshToken() error {
	c := d.Config
	accessToken, err := fetchAccessToken(c.APIBaseURL, c.GrantType, c.AppID, c.Secret)
	if err != nil {
		log.Printf("fetch access token error: %v\n", err)
		return ErrFetchToken
	}

	d.token = accessToken.Token
	// record last refreshed timestamp
	// or log timestamp
	log.Printf("access token refreshed, will expire in %d sec\n", accessToken.ExpireIn)

	return nil
}

func (d *TokenDaemon) GetToken() string {
	return string(d.token)
}

// call WeChat token API and fetch token
// GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
// {"access_token":"ACCESS_TOKEN","expires_in":7200}
func fetchAccessToken(apiBaseURL string, grantType, appID, secret string) (accessToken *AccessToken, err error) {
	if len(apiBaseURL) == 0 {
		apiBaseURL = defaultAPIBaseURL
	}
	if len(grantType) == 0 {
		grantType = defaultGrantType
	}
	url := fmt.Sprintf("%s?grant_type=%s&appid=%s&secret=%s",
		apiBaseURL, grantType, appID, secret)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("call WeChat API to get token error: %v\n", err)
		return
	}

	defer resp.Body.Close()

	// parse body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read response body error: %v\n", err)
		return
	}

	accessToken = &AccessToken{}
	err = json.Unmarshal(data, accessToken)
	if err != nil {
		log.Printf("unmarshal json error: %v\n", err)
		return
	}

	return
}
