package main

import (
	"log"

	"github.com/zyfdegh/wx-devdemo/api"
	"github.com/zyfdegh/wx-devdemo/tokensvc"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()

	app.Adapt(
		iris.DevLogger(),
		httprouter.New())

	app.Get("/", api.GetRoot)
	app.Get("/ping", api.GetPing)

	// start access token service
	if err := tokensvc.Start(); err != nil {
		log.Fatalf("start token service error: %v\n", err)
	}

	app.Listen(":80")
}
