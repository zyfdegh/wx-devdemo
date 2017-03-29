package svc

import (
	"fmt"
	"github.com/zyfdegh/wx-devdemo/util"
)

func CheckSig(token, timestamp, nonce string, signature string) bool {
	arr := util.SortByDict([]string{token, timestamp, nonce})
	if util.SHA1(fmt.Sprintf("%s%s%s", arr[0], arr[1], arr[2])) == signature {
		return true
	}
	return false
}
