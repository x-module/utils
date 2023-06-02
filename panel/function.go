/**
* Created by GoLand
* @file template_funtion.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/11 9:48 上午
* @desc 定义模板方法
 */

package panel

import (
	"fmt"
	"github.com/x-module/utils/utils/convertor"
	"github.com/x-module/utils/utils/cryptor"
	"github.com/x-module/utils/utils/datetime"
	"html/template"
	"reflect"
	"strings"
	"time"
)

// StaticPath 静态文件路由地址
const StaticPath = "/static/application/server/resource"

// InitializeTemplateFun 定义模板方法
func InitializeTemplateFun() template.FuncMap {
	return template.FuncMap{
		"staticPath": func() string {
			return StaticPath
		},
		"lang": func(a any) any {
			return a
		},
		"interfaceChek": func(a any) bool {
			res := fmt.Sprint(a)
			if res == "" {
				return false
			} else {
				return true
			}
		},
		"showDate": func(time time.Time) string {
			return time.Format(datetime.DateTemplate)
		},
		"date": func(a any) any {
			switch a.(type) {
			case time.Time:
				return a.(time.Time).Format(datetime.DateTimeTemplate)
			case string:
				t, _ := time.ParseInLocation(datetime.DateTimeTemplate, a.(string), time.Local)
				return t.Format(datetime.DateTimeTemplate)
			}
			return a
		},
		"transDate": func(a any) any {
			switch a.(type) {
			case time.Time:
				return a.(time.Time).Format(datetime.DateTemplate)
			case string:
				t, _ := time.ParseInLocation(datetime.ParseTimeTemplate, a.(string), time.Local)
				return t.Format(datetime.DateTemplate)
			}
			return a
		},
		"add": func(one int, two int) int {
			return one + two
		},

		"sub": func(one int, two float64) int {
			return int(float64(one) * two)
		},

		"subtraction": func(one int, two int) int {
			return one - two
		},

		"index": func(option any, num int) string {
			return ""
		},
		"baseName": func(path any) string {
			if path != nil {
				temp := strings.Split(path.(string), "/")
				return temp[len(temp)-1]
			}
			return ""
		},
		"isArray": func(params any) bool {
			kind := reflect.TypeOf(params).Kind().String()
			return kind == "slice"
		},
		"isHidden": func(elementType any) bool {
			elementType = fmt.Sprint(elementType)
			return elementType == "hidden"
		},
		"option": func(editAble, detailAble, deleteAble bool) bool {
			return editAble || detailAble || deleteAble
		},
		"langHtml": func(a any) any {
			return a
		},
		"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
			if cdnUrl == "" {
				return prefixUrl + assetsUrl
			}
			return cdnUrl + assetsUrl
		},
		"isLinkUrl": func(s string) bool {
			return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
		},
		"render": func(s, old, repl template.HTML) template.HTML {
			return template.HTML(strings.ReplaceAll(string(s), string(old), string(repl)))
		},
		"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
			return template.JS(strings.ReplaceAll(string(s), string(old), string(repl)))
		},
		"divide": func(a, b int) int {
			return a / b
		},
		"renderRowDataHTML": func(id, content template.HTML, value ...any) template.HTML {
			return template.HTML("")
		},
		// "renderRowDataJS": func(id template.HTML, content template.JS, value ...map[string]types.InfoItem) template.JS {
		//	return template.JS(types.ParseTableDataTmplWithID(id, string(content), value...))
		// },
		"attr": func(s template.HTML) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"js": func(s any) template.JS {
			if ss, ok := s.(string); ok {
				return template.JS(ss)
			}
			if ss, ok := s.(template.HTML); ok {
				return template.JS(ss)
			}
			return ""
		},
		"md5": func(params string) string {
			md5 := cryptor.Md5String(params)
			return md5
		},

		"checkStatus": func(status string, target string) bool {
			return status == target
		},
		"convertUTC": func(t string) string {
			t1, _ := time.Parse("2006-01-02T15:04:05Z", t)
			var cstSh, _ = time.LoadLocation("Asia/Shanghai") // 上海
			return t1.In(cstSh).Format("2006-01-02 15:04:05")
		},

		"formatFileSize": func(fileSize float64) (size string) {
			return convertor.FormatFileSize(fileSize)
		},
		"html": func(str string) any {
			return template.HTML(str)
		},
		"strEq": func(one string, tow string) bool {
			return one == tow
		},
	}
}
