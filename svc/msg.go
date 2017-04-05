package svc

import (
	"fmt"
	"github.com/zyfdegh/wx-devdemo/types"
)

func HandleTextMsg(msg types.TextMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, %s", msg.Content)

	return
}

func HandleImageMsg(msg types.ImageMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleVoiceMsg(msg types.VoiceMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleVideoMsg(msg types.VideoMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleShortVideoMsg(msg types.ShortVideoMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleLocationMsg(msg types.LocationMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleLinkMsg(msg types.LinkMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg.Msg)
	reply.Content = fmt.Sprintf("well, you send me a %s", msg.MsgType)
	return
}

func HandleUnknownMsg(msg types.Msg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(msg)
	reply.Content = fmt.Sprintf("well, you send me a unknown type of msg: %s", msg.MsgType)
	return
}
