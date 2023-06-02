/**
 * Created by PhpStorm.
 * @file   view.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/4/23 20:12
 * @desc   view.go
 */

package tools

import (
	"fmt"
	"github.com/x-module/utils/utils/fileutil"
	"github.com/x-module/utils/utils/strutil"
	"strings"
)

const ViewTemplate = "application/modules/generate/source/template.vue"
const Out = "frontend/src/views"

func GenerateView(table Table, fields []Field) {
	code, err := fileutil.ReadFileToString(ViewTemplate)
	if err != nil {
		panic(err)
	}
	temp := strings.Split(table.Name, "_")
	sourceTableName := strings.Join(temp[1:], "_")
	sourceTableName = strutil.CamelCase(sourceTableName)
	tableName := strings.ToUpper(sourceTableName[:1]) + sourceTableName[1:]
	fileName := tableName + "View.vue"

	// 生成当前的操作，如用户，权限等
	code = strings.ReplaceAll(code, "@target@", table.Comment)
	code = strings.ReplaceAll(code, "Permission", tableName)
	code = strings.ReplaceAll(code, "permission", sourceTableName)

	addCode := getAddFieldCode(sourceTableName, fields)
	code = strings.ReplaceAll(code, "<!--addField-->", addCode)
	showCode := getViewCode(sourceTableName, fields)
	code = strings.ReplaceAll(code, "<!--showField-->", showCode)
	searchCode := getSearchCode(sourceTableName, fields)
	code = strings.ReplaceAll(code, "<!--search-->", searchCode)

	fileName = Out + "/" + fileName
	WriteFile(fileName, code)
}

// <n-form-item path="Name" label="请输入装备分发表名称">
// <n-input v-model:value="distributeForm.Name" class="input-width-style"
//
// </n-form-item>
func getAddFieldCode(tableName string, fields []Field) string {
	str := ""
	for _, field := range fields {
		fieldName := FormatStr(field.ColumnName)
		if fieldName == "CreatedAt" || fieldName == "Id" || fieldName == "UpdatedAt" || fieldName == "DeletedAt" {
			continue
		}
		fmt.Println("------", field.ColumnName, ":", field.DataType)
		str += `<n-form-item path="` + fieldName + `" label="` + GetTableComment(field.ColumnComment) + `"> ` + "\n"
		if getType(field.DataType) == "string" {
			str += `<n-input v-model:value="` + tableName + `Form.` + fieldName + `" class="input-width-style"  placeholder="` + GetTableComment(field.ColumnComment) + `"/>` + "\n"
		} else {
			str += `<n-input-number v-model:value="` + tableName + `Form.` + fieldName + `" class="input-width-style"  placeholder="` + GetTableComment(field.ColumnComment) + `"/>` + "\n"
		}
		str += "</n-form-item>\n"
	}
	return str
}

func getViewCode(tableName string, fields []Field) string {
	str := ""
	for _, field := range fields {
		fieldName := FormatStr(field.ColumnName)
		if fieldName == "DeletedAt" {
			continue
		}
		str += `<n-descriptions-item label="` + GetTableComment(field.ColumnComment) + `">` + "\n"
		str += `{{ ` + tableName + `Form.` + fieldName + ` }}` + "\n"
		str += "</n-descriptions-item>\n"
	}
	return str
}
func getSearchCode(tableName string, fields []Field) string {
	str := ""
	for _, field := range fields {
		fieldName := FormatStr(field.ColumnName)
		if fieldName == "CreatedAt" || fieldName == "Id" || fieldName == "UpdatedAt" || fieldName == "DeletedAt" {
			continue
		}
		fmt.Println("------", field.ColumnName, ":", field.DataType)
		str += `<n-form-item path="` + fieldName + `" label="` + GetTableComment(field.ColumnComment) + `"> ` + "\n"
		if getType(field.DataType) == "string" {
			str += `<n-input v-model:value="searchForm.` + fieldName + `" class="search-input" clearable  placeholder="请输入搜索` + GetTableComment(field.ColumnComment) + `"/>` + "\n"
		} else {
			str += `<n-input-number v-model:value="searchForm.` + fieldName + `" class="search-input" clearable  placeholder="请输入搜索` + GetTableComment(field.ColumnComment) + `" :show-button="false"/>` + "\n"
		}
		str += "</n-form-item>\n"
	}
	return str
}
