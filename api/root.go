package api

import (
	"github.com/zyfdegh/wx-devdemo/env"
	"github.com/zyfdegh/wx-devdemo/svc"
	"gopkg.in/kataras/iris.v6"
	"log"
)

// GetRoot handles GET /
func GetRoot(ctx *iris.Context) {
	signature := ctx.URLParam("signature")
	timestamp := ctx.URLParam("timestamp")
	nonce := ctx.URLParam("nonce")
	echostr := ctx.URLParam("echostr")

	token, err := env.TOKEN.ToString()
	if err != nil {
		log.Printf("get env %s error: %v\n", env.TOKEN.String(), err)
		return
	}

	// log.Printf("signature: %s\n", signature)
	// log.Printf("timestamp: %s\n", timestamp)
	// log.Printf("nonce: %s\n", nonce)
	// log.Printf("echostr: %s\n", echostr)
	// log.Printf("token: %s\n", token)

	if len(signature) == 0 || len(timestamp) == 0 ||
		len(nonce) == 0 || len(echostr) == 0 {
		log.Println("refuse to check, lack of query params")
		return
	}

	ok := svc.CheckSig(token, timestamp, nonce, signature)
	if !ok {
		log.Printf("signature not ok\n")
		return
	}

	ctx.WriteString(echostr)
	return
}
