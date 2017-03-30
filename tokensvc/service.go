package tokensvc

import (
	"log"

	"github.com/zyfdegh/wx-devdemo/env"
)

var (
	DaemonInstance *TokenDaemon
)

func Start()(err error){
	appID, _ := env.APPID.ToString()
	secret, _ := env.SECRET.ToString()
	pollingSec, _ := env.POLLING_SEC.ToInt()

	log.Printf("APPID: %s\n", appID)
	// log.Printf("SECRET: %s\n", secret)
	log.Printf("POLLING_SEC: %d\n", pollingSec)

	config := DaemonConfig{
		AppID:  appID,
		Secret: secret,
		PollingSec: pollingSec,
	}

	daemon, err := NewTokenDaemon(config)
	if err != nil {
		log.Printf("new token daemon error: %v\n", err)
		return
	}
	err = daemon.Start()
	if err!=nil{
		log.Printf("start token daemon error: %v\n", err)
		return
	}
	DaemonInstance = daemon
	return
}
