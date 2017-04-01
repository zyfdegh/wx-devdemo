package api

import (
	"gopkg.in/kataras/iris.v6"
)

func dumbReply(ctx *iris.Context) {
	ctx.WriteString("success")
	return
}
