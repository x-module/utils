/**
* Created by Goland
* @file utils.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2021/11/11 2:49 下午
* @desc
 */

package utils

import (
	"math/rand"
	"strings"
	"time"
)

func ToLine(str string) string {
	out := ""
	for _, v := range str {
		if int(v) > 64 && int(v) < 91 {
			if out == "" {
				out += string(rune(int(v) + 32))
			} else {
				out += "_" + string(rune(int(v)+32))
			}
		} else {
			out += string(v)
		}
	}
	return out
}
func ToHump(str string) string {
	strs := strings.Split(str, "_")
	var strRes []string

	for _, str := range strs {
		str = UcFirst(str)
		strRes = append(strRes, str)
	}
	return strings.Join(strRes, "")
}
func UcFirst(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr

}
func RandString(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}
