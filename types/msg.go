package types

import (
	"github.com/zyfdegh/wx-devdemo/util"
)

const (
	Text       = "text"
	Image      = "image"
	Voice      = "voice"
	Video      = "video"
	ShortVideo = "shortvideo"
	Location   = "location"
	Link       = "link"

	// KeyMsgType is the key name of field MsgType
	KeyMsgType = "MsgType"
)

// Msg
type Msg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MsgId        int64  `xml:"MsgId"`
}

// TextMsg
// MsgType "text"
type TextMsg struct {
	Msg
	Content string `xml:"Content"`
}

// ImageMsg
// MsgType "image"
type ImageMsg struct {
	Msg
	PicUrl  string `xml:"PicUrl"`
	MediaId string `xml:"MediaId"`
}

// VoiceMsg
// MsgType "voice"
type VoiceMsg struct {
	Msg
	Format      string `xml:"Format"`
	MediaId     string `xml:"MediaId"`
	Recognition string `xml:"Recognition"`
}

// VideoMsg
// MsgType "video"
type VideoMsg struct {
	Msg
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

// ShortVideoMsg
// MsgType "shortvideo"
type ShortVideoMsg struct {
	Msg
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

// LocationMsg
// MsgType "location"
type LocationMsg struct {
	Msg
	Location_X string `xml:"Location_X"`
	Location_Y string `xml:"Location_Y"`
	Scale      string `xml:"Scale"`
	Label      string `xml:"Label"`
}

// LinkMsg
// MsgType "link"
type LinkMsg struct {
	Msg
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	Url         string `xml:"Url"`
}

// Reply
type Reply struct {
	XMLName      string `xml:"xml"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
}

// TextReply
// MsgType "text"
type TextReply struct {
	Reply
	Content string `xml:"Content"`
}

// VoiceReply
// MsgType "voice"
type VoiceReply struct {
	Reply
	MediaId string `xml:"MediaId"`
}

// VideoReply
// MsgType "video"
type VideoReply struct {
	Reply
	MediaId     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

// MusicReply
// MsgType "music"
type MusicReply struct {
	Reply
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	MusicURL     string `xml:"MusicURL"`
	HQMusicUrl   string `xml:"HQMusicUrl"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

// NewsReply contains images, links of articles
// MsgType "news"
type NewsReply struct {
	// 8 at most
	ArticleCount int        `xml:"ArticleCount"`
	Articles     []*Article `xml:"Articles"`
}

// Article of news
type Article struct {
	Item Item `xml:"item"`
}

// Item of article
type Item struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	// JPG/PNG file
	// 360*200 for 1st article
	// 200*200 for other articles
	PicUrl string `xml:"PicUrl"`
	Url    string `xml:"Url"`
}

// NewReply construct a Reply
// with CreateTime set to now
func NewReply() (reply *Reply) {
	reply = &Reply{}
	reply.CreateTime = util.Now()
	return
}

// NewReplyToMsg construct a Reply using input Msg
// with CreateTime set to now
// with FromUserName set to msg.ToUserName
// with ToUserName set to msg.FromUserName
func NewReplyToMsg(msg Msg) (reply *Reply) {
	reply = NewReply()
	reply.FromUserName = msg.ToUserName
	reply.ToUserName = msg.FromUserName
	return
}

// NewTextReply construct a TextReply
// with CreateTime set to now
func NewTextReply() (textReply *TextReply) {
	textReply = &TextReply{}
	textReply.Reply = *NewReply()
	textReply.MsgType = Text
	return
}

// NewTextReplyToMsg construct a TextReply using input Msg
// with CreateTime set to now
// with FromUserName set to msg.ToUserName
// with ToUserName set to msg.FromUserName
func NewTextReplyToMsg(msg Msg) (textReply *TextReply) {
	textReply = NewTextReply()
	textReply.FromUserName = msg.ToUserName
	textReply.ToUserName = msg.FromUserName
	return
}
