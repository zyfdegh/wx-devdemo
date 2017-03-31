package svc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/zyfdegh/wx-devdemo/tokensvc"
	"github.com/zyfdegh/wx-devdemo/types"
	"github.com/zyfdegh/wx-devdemo/util"
)

const (
	menuBaseURL = "https://api.weixin.qq.com/cgi-bin/menu/create"
)

var (
	// ErrCreateMenu returned when creating menu error
	ErrCreateMenu = errors.New("create menu error")
)

func CreateMenu(menuFile string) (err error) {
	token, err := tokensvc.D.GetToken()
	if err != nil {
		log.Printf("get token error: %v\n", err)
		return
	}

	url, err := util.AppendToken(menuBaseURL, token)
	if err != nil {
		log.Printf("append token to url error: %v\n", err)
		return
	}

	menu := &types.Menu{}
	err = util.DecodeJSONFile(menuFile, menu)
	if err != nil {
		log.Printf("read menu file error: %v\n", err)
		return
	}

	return createMenu(url, menu)
}

func createMenu(url string, menu *types.Menu) (err error) {
	resp, err := util.PostJSON(url, menu)
	if err != nil {
		log.Printf("post json error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %v\n", err)
		return
	}

	errResp := &types.WechatErrResp{}
	err = json.Unmarshal(data, errResp)
	if err != nil {
		log.Printf("unmarshal json error: %v\n", err)
		return
	}

	// {"errcode":0,"errmsg":"ok"}
	if errResp.ErrCode == 0 && errResp.ErrMsg == "ok" {
		return nil
	}

	// {"errcode":40018,"errmsg":"invalid button name size"}
	err = ErrCreateMenu
	log.Printf("create menu error, errcode: %d, errmsg: %s\n",
		errResp.ErrCode, errResp.ErrMsg)
	return err
}
