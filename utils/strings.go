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
	"unicode"
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

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Camel2Case 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
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
