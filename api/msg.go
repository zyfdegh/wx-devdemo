package api

import (
	"encoding/xml"
	"fmt"
	"github.com/zyfdegh/wx-devdemo/svc"
	"github.com/zyfdegh/wx-devdemo/types"
	"gopkg.in/kataras/iris.v6"
	"io/ioutil"
	"log"
)

// ReceiveMsg handles POST /msg
// Example msg(TextMsg):
//
// <xml>
//  <ToUserName><![CDATA[toUser]]></ToUserName>
//  <FromUserName><![CDATA[fromUser]]></FromUserName>
//  <CreateTime>1348831860</CreateTime>
//  <MsgType><![CDATA[text]]></MsgType>
//  <Content><![CDATA[this is a test]]></Content>
//  <MsgId>1234567890123456</MsgId>
//  </xml>
//
// Example reply(TextReply):
//
// <xml>
// <ToUserName><![CDATA[toUser]]></ToUserName>
// <FromUserName><![CDATA[fromUser]]></FromUserName>
// <CreateTime>12345678</CreateTime>
// <MsgType><![CDATA[text]]></MsgType>
// <Content><![CDATA[hello, user]]></Content>
// </xml>
func ReceiveMsg(ctx *iris.Context) {
	// log.Printf("url params: %+v\n", ctx.URLParams())
	fmt.Println("\n")

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("read body error: %v\n", err)
		dumbReply(ctx)
		return
	}

	var msgType string
	m := &types.Msg{}
	err = xml.Unmarshal(reqBody, m)
	if err != nil {
		log.Printf("parse body to msg error: %v\n", err)
		dumbReply(ctx)
		return
	}
	msgType = m.MsgType

	log.Printf("got msg, type: %s\n", msgType)

	var textReply = &types.TextReply{}

	switch msgType {
	case types.Text:
		textMsg := types.TextMsg{}
		err = xml.Unmarshal(reqBody, &textMsg)
		if err != nil {
			log.Printf("read xml to textMsg error: %v\n", err)
			break
		}
		log.Printf("textMsg: %+v\n", textMsg)
		textReply, err = svc.HandleTextMsg(textMsg)
	case types.Image:
		imageMsg := types.ImageMsg{}
		err = xml.Unmarshal(reqBody, &imageMsg)
		if err != nil {
			log.Printf("read xml to imageMsg error: %v\n", err)
			break
		}
		log.Printf("imageMsg: %+v\n", imageMsg)
		textReply, err = svc.HandleImageMsg(imageMsg)
	case types.Voice:
		voiceMsg := types.VoiceMsg{}
		err = xml.Unmarshal(reqBody, &voiceMsg)
		if err != nil {
			log.Printf("read xml to voiceMsg error: %v\n", err)
			break
		}
		log.Printf("voiceMsg: %+v\n", voiceMsg)
		textReply, err = svc.HandleVoiceMsg(voiceMsg)
	case types.Video:
		videoMsg := types.VideoMsg{}
		err = xml.Unmarshal(reqBody, &videoMsg)
		if err != nil {
			log.Printf("read xml to videoMsg error: %v\n", err)
			break
		}
		log.Printf("videoMsg: %+v\n", videoMsg)
		textReply, err = svc.HandleVideoMsg(videoMsg)
	case types.ShortVideo:
		shortVideoMsg := types.ShortVideoMsg{}
		err = xml.Unmarshal(reqBody, &shortVideoMsg)
		if err != nil {
			log.Printf("read xml to shortVideoMsg error: %v\n", err)
			break
		}
		log.Printf("shortVideoMsg: %+v\n", shortVideoMsg)
		textReply, err = svc.HandleShortVideoMsg(shortVideoMsg)
	case types.Location:
		locationMsg := types.LocationMsg{}
		err = xml.Unmarshal(reqBody, &locationMsg)
		if err != nil {
			log.Printf("read xml to locationMsg error: %v\n", err)
			break
		}
		log.Printf("locationMsg: %+v\n", locationMsg)
		textReply, err = svc.HandleLocationMsg(locationMsg)
	case types.Link:
		linkMsg := types.LinkMsg{}
		err = xml.Unmarshal(reqBody, &linkMsg)
		if err != nil {
			log.Printf("read xml to linkMsg error: %v\n", err)
			break
		}
		log.Printf("linkMsg: %+v\n", linkMsg)
		textReply, err = svc.HandleLinkMsg(linkMsg)
	default:
		msg := types.Msg{}
		err = xml.Unmarshal(reqBody, &msg)
		if err != nil {
			log.Printf("read xml to msg error: %v\n", err)
			break
		}
		log.Printf("msg: %+v\n", msg)
		textReply, err = svc.HandleUnknownMsg(msg)
	}

	if err != nil {
		log.Printf("handle %s msg error: %v\n", msgType, err)
		dumbReply(ctx)
		return
	}

	log.Printf("reply: %+v\n", textReply)

	ctx.XML(iris.StatusOK, textReply)
	return
}

func dumbReply(ctx *iris.Context) {
	ctx.WriteString("success")
	return
}
