package api

import (
	"github.com/zyfdegh/wx-devdemo/svc"
	"github.com/zyfdegh/wx-devdemo/types"
	"gopkg.in/kataras/iris.v6"
	"log"
)

// ReceiveMsg handles POST /msg
func ReceiveMsg(ctx *iris.Context) {
	log.Printf("url params: %+v\n", ctx.URLParams())

	var msg interface{}
	err := ctx.ReadXML(msg)
	if err != nil {
		log.Printf("read xml msg error: %v\n", err)
		return
	}

	log.Printf("msg: %+v\n", msg)

	var textReply = &types.TextReply{}

	switch msg.(types.Msg).MsgType {
	case types.Text:
		textReply, err = svc.HandleTextMsg(msg.(types.TextMsg))
	case types.Image:
		textReply, err = svc.HandleImageMsg(msg.(types.ImageMsg))
	case types.Voice:
		textReply, err = svc.HandleVoiceMsg(msg.(types.VoiceMsg))
	case types.Video:
		textReply, err = svc.HandleVideoMsg(msg.(types.VideoMsg))
	case types.ShortVideo:
		textReply, err = svc.HandleShortVideoMsg(msg.(types.ShortVideoMsg))
	case types.Location:
		textReply, err = svc.HandleLocationMsg(msg.(types.LocationMsg))
	case types.Link:
		textReply, err = svc.HandleLinkMsg(msg.(types.LinkMsg))
	default:
		textReply, err = svc.HandleUnknownMsg(msg.(types.Msg))
	}

	if err != nil {
		log.Printf("handle msg error: %v\n", err)
		ctx.WriteString("success")
		return
	}

	ctx.XML(iris.StatusOK, textReply)
	return
}
