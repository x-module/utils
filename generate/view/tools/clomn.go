/**
 * Created by PhpStorm.
 * @file   view.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/4/23 20:12
 * @desc   view.go
 */

package tools

import (
	"github.com/x-module/utils/utils/fileutil"
	"github.com/x-module/utils/utils/strutil"
	"strings"
)

const ColumnTemplate = "application/modules/generate/source/columns.ts"
const OutDir = "frontend/src/element/columns"

func GenerateColumn(table Table, fields []Field) {
	code, err := fileutil.ReadFileToString(ColumnTemplate)
	if err != nil {
		panic(err)
	}
	temp := strings.Split(table.Name, "_")
	tableName := strings.Join(temp[1:], "_")
	tableName = strutil.CamelCase(tableName)
	tableName = strings.ToUpper(tableName[:1]) + tableName[1:]
	fileName := tableName + ".ts"

	classCode := getClassCode(tableName, fields)
	showFieldCode := getShowField(tableName, fields)
	rulesCode := getRules(fields)
	// 生成当前的操作，如用户，权限等
	code = strings.ReplaceAll(code, "//class", classCode)
	code = strings.ReplaceAll(code, "//showField", showFieldCode)
	code = strings.ReplaceAll(code, "//rules", rulesCode)
	code = strings.ReplaceAll(code, "Permission", tableName)
	code = strings.ReplaceAll(code, "/*", "")
	code = strings.ReplaceAll(code, "*/", "")

	fileName = OutDir + "/" + fileName
	WriteFile(fileName, code)
}

func getClassCode(className string, fields []Field) string {
	str := "class " + className + " {"
	for _, field := range fields {
		str += "    // @ts-ignore \n"
		str += "    " + FormatStr(field.ColumnName) + ": " + getType(field.DataType) + ";\n"
	}
	str += "}"
	return str
}

// {
// title: 'HttpPath',
// key: 'HttpPath',
// },
func getShowField(className string, fields []Field) string {
	str := ""
	for _, field := range fields {
		fieldName := FormatStr(field.ColumnName)
		if fieldName == "DeletedAt" {
			continue
		} else if fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
			str += "        {\n"
			str += "            title:'" + GetTableComment(field.ColumnComment) + "',\n"
			str += "            key:'" + fieldName + "',\n"
			str += `            render(row: ` + className + `) {
                return h(
                    NTime,
                    {
                        format:"yyyy-MM-dd HH:mm:ss",
                    },
                )
            },`
			str += "        },\n"
		} else {
			str += "        {\n"
			str += "            title:'" + GetTableComment(field.ColumnComment) + "',\n"
			str += "            key:'" + fieldName + "',\n"
			str += "        },\n"
		}
	}
	return str
}

// Name: [
// {
// required: true,
// message: '权限名称不能为空',
// trigger: ['blur', 'change'],
// },
// ],
func getRules(fields []Field) string {
	str := ""
	for _, field := range fields {
		str += "    " + FormatStr(field.ColumnName) + ": [\n"
		str += "        {\n"
		str += "            type: '" + getType(field.DataType) + "',\n"
		str += "            required: true,\n"
		str += "            message: '" + GetTableComment(field.ColumnComment) + "',\n"
		str += "            trigger: ['blur', 'change'],\n"
		str += "        },\n"
		str += "    ],\n"
	}
	return str
}
