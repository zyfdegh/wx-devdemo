package main

import (
	"log"

	"github.com/zyfdegh/wx-devdemo/api"
	"github.com/zyfdegh/wx-devdemo/tokensvc"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

const (
	defaultMenuFile = "./conf/menu.json"
)

func main() {
	app := iris.New()

	app.Adapt(
		iris.DevLogger(),
		httprouter.New())

	app.Get("/", api.GetRoot)
	app.Get("/ping", api.GetPing)
	app.Post("/msg", api.ReceiveMsg)

	// start access token service
	if err := tokensvc.Start(); err != nil {
		log.Fatalf("start token service error: %v\n", err)
	}

	// init menu buttons
	// 未通过微信认证的个人号，不支持创建自定义菜单
	// 返回 errcode: 48001, errmsg: api unauthorized hint: [TDGZKA0472vr29!]
	// 获得条件：
	// 订阅号必须通过微信认证
	// 服务号自动获得
	// https://mp.weixin.qq.com/merchant/store?action=detail&t=wxverify/detail&info=verify&lang=zh_CN
	//
	// if err := svc.CreateMenu(defaultMenuFile); err != nil {
	// 	log.Fatalf("create menu failed: %v\n", err)
	// }

	app.Listen(":80")
}
