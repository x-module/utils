/**
 * Created by PhpStorm.
 * @file   base.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/4/23 20:13
 * @desc   base.go
 */

package tools

import (
	"github.com/go-xmodule/utils/utils/fileutil"
	"github.com/go-xmodule/utils/utils/xlog"
	"io"
	"log"
	"os"
	"strings"
)

type Field struct {
	ColumnName    string `json:"column_name"`
	DataType      string `json:"data_type"`
	ColumnComment string `json:"column_comment"`
}
type Table struct {
	Name    string `json:"Name,omitempty"`
	Comment string `json:"Comment,omitempty"`
}

func WriteFile(fileName string, code string) {
	/******************* 使用 io.WriteString 写入文件 **********************/
	f1, err1 := OpenFile(fileName)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	defer f1.Close()
	n, err1 := io.WriteString(f1, code) // 写入文件(字符串)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	xlog.Logger.Infof("file: %s", fileName)
	xlog.Logger.Infof("写入 %d 个字节\n", n)
}

// OpenFile 判断文件是否存在  存在则OpenFile 不存在则Create
func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		xlog.Logger.Debug("文件不存在")
		return os.Create(filename) // 创建文件
	}
	fileutil.RemoveFile(filename)
	xlog.Logger.Debug("文件存在")
	return os.Create(filename) // 创建文件
}

func FormatStr(str string) string {
	out := ""
	for _, item := range strings.Split(str, "_") {
		out += strings.ToUpper(item[:1]) + item[1:]
	}
	return out
}

func GetTableComment(comment string) string {
	return "请输入" + strings.ReplaceAll(strings.TrimSpace(comment), "\n", "")
}

func getType(fieldType string) string {
	switch fieldType {
	case "int":
		return "number"
	case "tinyint":
		return "number"
	case "bigint":
		return "number"
	default:
		return "string"
	}
}
