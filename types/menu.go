package types

// Menu is the whole buttons of wechat app
type Menu struct {
	Button []*Button `json:"button"`
}

// Button is the struct of menu,
// including 1st level buttons(or buttons) and 2nd level sub buttons
type Button struct {
	// Type marks up function of button
	// must set
	// e.g. "view", "click", "miniprogram",
	// "scancode_waitmsg", "scancode_push",
	// "pic_sysphoto", "pic_photo_or_album", "pic_weixin",
	// "location_select", "media_id", "view_limited"
	Type string `json:"type"`
	// Name is the displayed name of button
	// must set
	// 4 Chinese characters at most on 1st level menu
	// 7 Chinese characters at most on 2nd level menu
	Name string `json:"name"`
	// Key is the ID of button
	// must set when Type is "click"
	// 128 bytes at most
	Key string `json:"key"`
	// SubButton is the nested 2nd level buttons
	// 5 sub buttons at most
	SubButton []*Button `json:"sub_button"`
	// URL is the redirect link of button
	// must set when Type is "view" or "miniprogram"
	// 1024 bytes at most
	URL string `json:"url"`
	// MediaID is the ID of picture(or other types of file) resource
	// must set when Type is "media_id" or "view_limited"
	MediaID string `json:"media_id"`
	// AppID is the ID of miniprogram
	// must set when Type is "miniprogram"
	AppID string `json:"appid"`
	// PagePath is the path of miniprogram
	// must set when Type is "miniprogram"
	PagePath string `json:"pagepath"`
}
