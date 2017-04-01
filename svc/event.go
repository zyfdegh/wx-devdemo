package svc

import (
	"fmt"
	"github.com/zyfdegh/wx-devdemo/types"
)

func HandleSubscribeEvent(event types.SubscribeEvent) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToEvent(event.Msg)
	reply.Content = fmt.Sprintf("%s", "蟹蟹关注这个测试号，回复1有惊喜")
	return
}

func HandleUnsubscribeEvent(event types.UnsubscribeEvent) (err error) {
	return
}

func HandleScanEvent(event types.ScanEvent) (err error) {
	return
}

func HandleLocationEvent(event types.LocationEvent) (err error) {
	return
}

func HandleClickEvent(event types.ClickEvent) (err error) {
	return
}

func HandleViewEvent(event types.ViewEvent) (err error) {
	return
}

func HandleUnknownEvent(event types.EventMsg) (reply *types.TextReply, err error) {
	reply = types.NewTextReplyToMsg(event.Msg)
	reply.Content = fmt.Sprintf("well, you send me a unknown type of event: %s", event.Event)
	return
}
