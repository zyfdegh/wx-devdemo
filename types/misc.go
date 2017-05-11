package types

type Token string

// AccessToken is WeChat returned access token structure
type AccessToken struct {
	Token    Token `json:"access_token"`
	ExpireIn int   `json:"expires_in"`
}

func (t Token) String() string {
	if len(string(t)) == 0 {
		return ""
	}
	return "***"
}
