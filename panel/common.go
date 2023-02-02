/**
* Created by GoLand
* @file common.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/7 4:58 下午
* @desc  基础组件
 */

package panel

import "html/template"

// Type 组件类型
type Type uint8

// 组件类型定义
const (
	Default Type = iota
	Text
	Hidden
	SelectSingle
	SelectSingleGroup
	Select
	IconPicker
	SelectBox
	File
	Password
	RichText
	Datetime
	DatetimeRange
	Radio
	Checkbox
	CheckboxStacked
	CheckboxSingle
	Email
	Date
	DateRange
	Url
	Ip
	Color
	Array
	Currency
	Rate
	Number
	NumberRange
	TextArea
	Custom
	Switch
	Code
	Slider
	Image
)

// 获取类型对应元素
func (t Type) String() string {
	switch t {
	case Default:
		return "default"
	case Text:
		return "text"
	case Hidden:
		return "hidden"
	case SelectSingle:
		return "select_single"
	case SelectSingleGroup:
		return "select_single_group"
	case Select:
		return "select"
	case IconPicker:
		return "icon_picker"
	case SelectBox:
		return "selectbox"
	case File:
		return "file"
	case Password:
		return "password"
	case RichText:
		return "richtext"
	case Rate:
		return "rate"
	case Checkbox:
		return "checkbox"
	case CheckboxStacked:
		return "checkbox_stacked"
	case CheckboxSingle:
		return "checkbox_single"
	case Date:
		return "datetime"
	case DateRange:
		return "datetime_range"
	case Datetime:
		return "datetime"
	case DatetimeRange:
		return "datetime_range"
	case Radio:
		return "radio"
	case Slider:
		return "slider"
	case Array:
		return "array"
	case Email:
		return "email"
	case Url:
		return "url"
	case Ip:
		return "ip"
	case Color:
		return "color"
	case Currency:
		return "currency"
	case Number:
		return "number"
	case NumberRange:
		return "number_range"
	case TextArea:
		return "textarea"
	case Custom:
		return "custom"
	case Switch:
		return "switch"
	case Code:
		return "code"
	case Image:
		return "image"
	default:
		panic("wrong form type")
	}
}

// FieldModel display 元素
type FieldModel struct {
	ID    any
	Value any
	Row   map[string]any
}

// FieldDisplayFun table展示自定义方法
type FieldDisplayFun func(model FieldModel) any

// TableField table 元数据
type TableField struct {
	Field    string
	ShowName string
	Sortable bool
	Display  FieldDisplayFun
	Export   bool
}
type FieldList []TableField

// ========================== form ================================

type InitSearch struct {
	OptionGroups []OptionGroup
	DefaultValue string
}
type OptionGroup struct {
	Group   string         `json:"group"`
	Options []SelectOption `json:"options"`
}

// SelectOption 下来框元数据
type SelectOption struct {
	Text          string            `json:"text"`
	Value         string            `json:"value"`
	TextHTML      template.HTML     `json:"-"`
	Selected      bool              `json:"-"`
	SelectedLabel template.HTML     `json:"-"`
	Extra         map[string]string `json:"-"`
}

// ValidatorType 验证类型
type ValidatorType int

// 验证类型定义
const (
	ValidatorRequire ValidatorType = iota
	ValidatorUnique
)

// ActionType 操作类型
type ActionType int

// FormField form表单原数据定义
type FormField struct {
	Field          string
	Value          any
	Value2         any
	FormType       Type
	FieldName      string
	FieldClass     string
	Placeholder    string
	Must           bool
	HelpMsg        string
	NotAllowEdit   bool
	HideBackButton bool
	Editable       bool
	NoIcon         bool
	Readonly       bool

	Validator []ValidatorItem

	OptionGroups  []OptionGroup
	Options       []SelectOption
	SelectedLabel template.HTML
	OptionExt     template.JS
	OptionExt2    template.JS
	ValueArr      []string
	Label         string
	Display       FieldDisplayFun

	CustomContent any
	CustomJs      template.JS
	CustomCss     template.CSS
}

// FormFieldList form 表单数据列表
type FormFieldList []FormField
