/**
* Created by GoLand
* @file add.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/8 9:44 下午
* @desc add.go
 */

package panel

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// Info 显示结构体
type Info struct {
	curFieldListIndex int
	fieldList         FormFieldList
	fieldSort         []string
	context           *gin.Context
}

func NewInfo(context *gin.Context) *Info {
	return &Info{
		curFieldListIndex: -1,
		context:           context,
		fieldSort:         make([]string, 0),
	}
}

// Display 格式化输出数据
func (s *Info) Display(fun FieldDisplayFun) *Info {
	s.fieldList[s.curFieldListIndex].Display = fun
	return s
}

// AddField 添加字段
func (s *Info) AddField(fieldName string, field string, formType ...Type) *Info {
	fType := Default
	if len(formType) > 0 {
		fType = formType[0]
	}
	s.fieldList = append(s.fieldList, FormField{
		Field:       field,
		FieldClass:  field,
		FormType:    fType,
		FieldName:   fieldName,
		Editable:    false,
		NoIcon:      false,
		Placeholder: "请输入 " + fieldName,
	})
	s.fieldSort = append(s.fieldSort, field)
	s.curFieldListIndex++
	return s
}

// SetSelectedLabel 设置selected label
func (s *Info) SetSelectedLabel(labels template.HTML) *Info {
	s.fieldList[s.curFieldListIndex].SelectedLabel = labels
	return s
}

// Options 设置select 选项
func (s *Info) Options(options []SelectOption) *Info {
	s.fieldList[s.curFieldListIndex].Options = options
	return s
}

// HelpMsg 帮助信息
func (s *Info) HelpMsg(msg string) *Info {
	s.fieldList[s.curFieldListIndex].HelpMsg = msg
	return s
}

// HideBackButton 隐藏返回按钮
func (s *Info) HideBackButton(hideBackButton bool) *Info {
	s.fieldList[s.curFieldListIndex].HideBackButton = hideBackButton
	return s
}

// GetField 获取显示的字段，模板调用
func (s *Info) GetField() FormFieldList {
	return s.fieldList
}

// GetInfo 获取显示模板对象
func (s *Info) GetInfo(data any) *Info {
	values := ConvertModeToMap(data)
	for k, v := range s.fieldList {
		s.fieldList[k].Value = values[v.Field]
		if v.Display != nil {
			s.fieldList[k].Value = s.fieldList[k].Display(FieldModel{
				ID:    s.fieldList[k].Field,
				Value: s.fieldList[k].Value,
				Row:   s.getRowValue(),
			})
		}
	}
	return s
}

// 获取当前数据
func (s *Info) getRowValue() map[string]any {
	row := map[string]any{}
	for _, v := range s.fieldList {
		row[v.Field] = v.Value
	}
	return row
}
