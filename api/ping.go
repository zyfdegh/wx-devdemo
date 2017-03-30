package api

import (
	"gopkg.in/kataras/iris.v6"
)

// GetPing handles GET /ping
func GetPing(ctx *iris.Context) {
	ctx.WriteString("pong")
	return
}
