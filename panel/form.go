/**
* Created by GoLand
* @file add.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/8 9:44 下午
* @desc 表单模具，数据编辑/展示
 */

package panel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-utils-module/utils/handler"
	"html/template"
	"strings"
)

// EditSourceDataFun 自定义编辑元数据修改方法
type EditSourceDataFun func(map[string]any) map[string]any

// Form Form表单
type Form struct {
	curFieldListIndex int
	model             handler.ModelAction
	fieldList         FormFieldList
	fieldSort         []string
	context           *gin.Context
	isEdit            bool
	editSourceDataFun EditSourceDataFun
	actionType        ActionType
}

func (f *Form) SetModel(model handler.ModelAction) *Form {
	f.model = model
	return f
}
func (f *Form) SetAction(actionType ActionType) *Form {
	f.actionType = actionType
	return f
}

// NewForm 新表单定义
func NewForm(context *gin.Context, actionType ActionType) *Form {
	return &Form{
		curFieldListIndex: -1,
		context:           context,
		fieldSort:         make([]string, 0),
		isEdit:            true,
		fieldList:         []FormField{},
		actionType:        actionType,
	}
}

// EditForm 编辑的表单
func (f *Form) EditForm() *Form {
	f.isEdit = true
	return f
}

// AddField 添加字段
func (f *Form) AddField(fieldName string, field string, formType Type) *Form {
	f.fieldList = append(f.fieldList, FormField{
		Field:       field,
		FieldClass:  field,
		FormType:    formType,
		FieldName:   fieldName,
		Editable:    true,
		NoIcon:      false,
		Placeholder: "请输入 " + fieldName,
		Validator:   []ValidatorItem{},
	})
	f.fieldSort = append(f.fieldSort, field)
	f.curFieldListIndex++
	return f
}

// Placeholder 添加字段input输入提示
func (f *Form) Placeholder(label string) *Form {
	f.fieldList[f.curFieldListIndex].Placeholder = label
	return f
}

// SetEditSourceDataFun 增加当前字段编辑器修改方法
func (f *Form) SetEditSourceDataFun(fun EditSourceDataFun) *Form {
	f.editSourceDataFun = fun
	return f
}

// SetSelectedLabel 设置select选择项
func (f *Form) SetSelectedLabel(labels template.HTML) *Form {
	f.fieldList[f.curFieldListIndex].SelectedLabel = labels
	return f
}

// Options 设置select标记的options数据
func (f *Form) Options(options []SelectOption) *Form {
	newOpt := make([]SelectOption, len(options))
	copy(newOpt, options)
	f.fieldList[f.curFieldListIndex].Options = newOpt
	return f
}

// OptionGroup 设置select标记的options数据
func (f *Form) OptionGroup(optionGroups []OptionGroup) *Form {
	f.fieldList[f.curFieldListIndex].OptionGroups = optionGroups
	return f
}

// Require 当前字段必填  -ok
func (f *Form) Require() *Form {
	f.fieldList[f.curFieldListIndex].Must = true
	f.fieldList[f.curFieldListIndex].Validator = append(f.fieldList[f.curFieldListIndex].Validator, ValidatorItem{
		ValidatorType: ValidatorRequire,
		ValidatorMsg:  f.fieldList[f.curFieldListIndex].FieldName + " 不能为空",
	})
	return f
}

// Require 当前字段唯一
func (f *Form) Unique() *Form {
	f.fieldList[f.curFieldListIndex].Must = true
	f.fieldList[f.curFieldListIndex].Validator = append(f.fieldList[f.curFieldListIndex].Validator, ValidatorItem{
		ValidatorType: ValidatorUnique,
		ValidatorMsg:  f.fieldList[f.curFieldListIndex].FieldName + " 已经存在",
	})
	return f
}

// HelpMsg 提示信息
func (f *Form) HelpMsg(msg string) *Form {
	f.fieldList[f.curFieldListIndex].HelpMsg = msg
	return f
}

// Editable 可编辑
func (f *Form) Editable() *Form {
	f.fieldList[f.curFieldListIndex].Editable = true
	return f
}

// Readonly 可编辑
func (f *Form) Readonly() *Form {
	f.fieldList[f.curFieldListIndex].Readonly = true
	return f
}

// Disable 不可编辑
func (f *Form) Disable() *Form {
	f.fieldList[f.curFieldListIndex].Editable = false
	return f
}

// Default 默认值
func (f *Form) Default(defaultValue any) *Form {
	f.fieldList[f.curFieldListIndex].Value = defaultValue
	return f
}

// DefaultValueTow 默认值
func (f *Form) DefaultValueTow(defaultValue any) *Form {
	f.fieldList[f.curFieldListIndex].Value2 = defaultValue
	return f
}

// NotAllowEdit 不允许编辑
func (f *Form) NotAllowEdit(notAllowEdit bool) *Form {
	f.fieldList[f.curFieldListIndex].NotAllowEdit = notAllowEdit
	return f
}

// ParseForm 解析当前form定义
func (f *Form) ParseForm() *Form {
	Validator := NewValidator(f.model)
	for k, _ := range f.fieldList {
		if len(f.fieldList[k].Validator) != 0 {
			Validator.AddValidator(f.actionType, f.fieldList[k], f.fieldList[k].Validator)
		}
	}
	return f
}

// HideBackButton 隐藏返回按钮
func (f *Form) HideBackButton(hideBackButton bool) *Form {
	f.fieldList[f.curFieldListIndex].HideBackButton = hideBackButton
	return f
}

// GetField 获取form表单元素--template使用
func (f *Form) GetField() FormFieldList {
	return f.fieldList
}

// GetForm 获取当前表单
func (f *Form) GetForm(data any) *Form {
	if !f.isEdit {
		return f
	}
	values := ConvertModeToMap(data)
	if f.editSourceDataFun != nil {
		values = f.editSourceDataFun(values)
	}
	for k, v := range f.fieldList {
		if value, exist := values[v.Field]; exist {
			f.fieldList[k].Value = value
		}
	}
	return f.extend()
}

// 扩展选项
func (f *Form) extend() *Form {
	for k, v := range f.fieldList {
		if len(f.fieldList[k].Options) != 0 {
			for _k, _v := range f.fieldList[k].Options {
				if strings.Contains(fmt.Sprint(v.Value), _v.Value) {
					f.fieldList[k].Options[_k].SelectedLabel = "selected"
				}
			}
		}
	}
	return f
}
