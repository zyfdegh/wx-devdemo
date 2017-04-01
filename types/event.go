package types

const (
	EventSubscribe   = "subscribe"
	EventUnsubscribe = "unsubscribe"
	EventScan        = "SCAN"
	EventLocation    = "LOCATION"
	EventClick       = "CLICK"
	EventView        = "VIEW"
	// dup
	// EventScanSubscribe = "subscribe"
)

// EventMsg is a special type of Msg
// MsgType "event"
// Event type of event
type EventMsg struct {
	Msg
	Event string `xml:"Event"`
}

// SubscribeEvent
// Event "subscribe"
type SubscribeEvent struct {
	EventMsg
}

// SubscribeEvent
// Event "unsubscribe"
type UnsubscribeEvent struct {
	EventMsg
}

// ScanSubscribeEvent 用户未关注时，扫码关注后的事件推送
// Event "subscribe"
// EventKey "qrscene_xxxxxx"
// Ticket ticket of QRCode image
type ScanSubscribeEvent struct {
	EventMsg
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

// ScanEvent 用户已关注时，扫码关注后的事件推送
// Event "SCAN"
// EventKey scene_id "xxxxxx"
// Ticket ticket of QRCode image
type ScanEvent struct {
	EventMsg
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

// LocationEvent
// Event "LOCATION"
// Latitude 纬度
// Longitude 经度
// Precision 精度
type LocationEvent struct {
	EventMsg
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
}

// ClickEvent
// Event "CLICK"
// EventKey the key in menu
type ClickEvent struct {
	EventMsg
	EventKey string `xml:"EventKey"`
}

// ViewEvent
// Event "VIEW"
// EventKey the key of view, commonly a URL
type ViewEvent struct {
	EventMsg
	EventKey string `xml:"EventKey"`
}
