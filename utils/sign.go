/**
 * Created by PhpStorm.
 * @file   sign.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/1 18:01
 * @desc   sign.go
 */

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-xmodule/utils/utils/cryptor"
	"github.com/golang-module/carbon"
	"sort"
	"strings"
	"time"
)

// Sign 接口请求签名
func SignParams(paramsStruct any, secret string) map[string]any {
	var params map[string]any
	b, _ := json.Marshal(paramsStruct)
	json.Unmarshal(b, &params)
	timeStamp := carbon.Now().Timestamp()
	params["timestamp"] = timeStamp

	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	for _, k := range keys {
		dataParams = dataParams + strings.TrimSpace(k) + "=" + strings.TrimSpace(fmt.Sprint(params[k])) + "&"
	}
	dataParams = fmt.Sprintf("%s@%s@%s", secret, dataParams, secret)
	// 对字符串进行sha1哈希
	params["sign"] = cryptor.Sha1(dataParams)
	return params
}

// Sign 接口请求签名
func Sign(params map[string]any, secret string) string {
	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	for _, k := range keys {
		dataParams = dataParams + strings.TrimSpace(k) + "=" + strings.TrimSpace(fmt.Sprint(params[k])) + "&"
	}
	dataParams = fmt.Sprintf("%s@%s@%s", secret, dataParams, secret)
	// 对字符串进行sha1哈希
	sign := cryptor.Sha1(dataParams)
	return sign
}

func ApiSign(url string, secret string) string {
	ts := fmt.Sprint(carbon.Now().Timestamp())
	signStr := fmt.Sprintf("%s@%s@%s", secret, ts, secret)
	// 对字符串进行sha1哈希
	sign := cryptor.Sha1(signStr)
	if strings.Contains(url, "?") {
		url += fmt.Sprintf("&ts=%s&sign=%s", ts, sign)
	} else {
		url += fmt.Sprintf("?ts=%s&sign=%s", ts, sign)
	}
	return url
}

func RequestSign(ts string, secret string) string {
	signStr := fmt.Sprintf("%s@%s@%s", secret, ts, secret)
	return cryptor.Sha1(signStr)
}

func RpcSign(secret string) (string, int64) {
	ts := time.Now().Unix()
	signStr := fmt.Sprintf("%s@%d@%s", secret, ts, secret)
	return cryptor.Sha1(signStr), ts
}
