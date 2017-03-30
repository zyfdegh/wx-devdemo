package main

import (
	"github.com/zyfdegh/wx-devdemo/api"

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

	app.Listen(":80")
}
