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

	token, err := env.TOKEN.ToString()
	if err != nil {
		log.Printf("get env %s error: %v\n", env.TOKEN.String(), err)
		return
	}

	log.Printf("signature: %s\n", signature)
	log.Printf("timestamp: %s\n", timestamp)
	log.Printf("nonce: %s\n", nonce)
	log.Printf("token: %s\n", token)

	ok := svc.CheckSig(token, timestamp, nonce, signature)
	if !ok {
		log.Printf("signature not ok\n")
	}
}
