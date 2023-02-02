/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/2 15:31
 * @desc   utils.go
 */

package utils

import (
	"encoding/json"
	"github.com/go-utils-module/utils/global"
)

func Success(status int) bool {
	return global.ErrCode(status) == global.Success
}
func JsonString(params any) string {
	b, _ := json.Marshal(params)
	return string(b)
}
