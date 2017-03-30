package env

import (
	"errors"
	"log"
	"os"
	"strconv"
)

var (
	// keys of environment variables

	// TOKEN is token of WeChat API, this is the developer-defined TOKEN, not access_token
	// For more, see https://mp.weixin.qq.com/wiki/8/f9a0b8382e0b77d87b3bcc1ce6fbc104.html
	TOKEN = Env("TOKEN")

	// APPID is the query param 'appid' to fetch access_token
	APPID = Env("APPID")

	// SECRET is the query param 'secret' to fetch access_token
	SECRET = Env("SECRET")

	// POLLING_SEC is the time period that token will be refreshed
	// interger number, in second
	POLLING_SEC = Env("POLLING_SEC")

	// ErrNotSet returned when the key of env is not set
	ErrNotSet = errors.New("env not set")
)

// Env is struct of environment variable
type Env string

func init() {
	mustSet(TOKEN)
	mustSet(APPID)
	mustSet(SECRET)
}

// ToInt parse value to int
func (e Env) ToInt() (int, error) {
	val, found := lookup(e)
	if !found {
		return 0, ErrNotSet
	}

	i, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// ToBool parse value to bool
func (e Env) ToBool() (bool, error) {
	val, found := lookup(e)
	if !found {
		return false, ErrNotSet
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (e Env) ToString() (string, error) {
	val, found := lookup(e)
	if !found {
		return "", ErrNotSet
	}
	return val, nil
}

// String convert Env to string
func (e Env) String() string {
	return string(e)
}

func lookup(e Env) (val string, found bool) {
	return os.LookupEnv(e.String())
}

func mustSet(e Env) {
	_, found := lookup(e)
	if !found {
		log.Fatalf("env %s not set\n", e)
	}
}
