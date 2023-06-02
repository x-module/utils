/**
* Created by GoLand
* @file table.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/7 3:57 下午
* @desc table.go
 */

package panel

import (
	"fmt"
	"github.com/x-module/utils/handler"
	"github.com/x-module/utils/utils/datetime"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// SearchField 搜索字段定义
type SearchField struct {
	Field    string
	ShowName string
	Notice   string
	Type     string
}
type WhereClosure func(*gorm.DB) *gorm.DB

// Grid 当前显示数据表格
type Grid struct {
	curFieldListIndex int
	fieldList         FieldList
	context           *gin.Context
	PageData          handler.PageData
	fieldFunMap       map[string]FieldDisplayFun
	fieldLabelMap     map[string]bool
	fieldImgMap       map[string]bool
	fieldIconMap      map[string]bool
	fieldSort         []string
	SearchFields      []SearchField
	// 页面提示
	Tips       string
	DeleteAble bool
	EditAble   bool
	DetailAble bool
	where      WhereClosure
	// ===============================
	Export     bool
	DataAdd    bool
	HasSearch  bool
	InitSearch InitSearch
}

// Field 显示字段信息
type Field struct {
	Field   string
	Row     map[string]any
	Value   any
	Display FieldDisplayFun
	Label   bool
	Icon    bool
	Image   bool
	Array   []string
}

// NewGrid 新对象
func NewGrid(context *gin.Context) *Grid {
	return &Grid{
		curFieldListIndex: -1,
		context:           context,
		fieldFunMap:       make(map[string]FieldDisplayFun),
		fieldLabelMap:     make(map[string]bool),
		fieldImgMap:       make(map[string]bool),
		fieldIconMap:      make(map[string]bool),
		fieldSort:         make([]string, 0),
		DataAdd:           true,
		Export:            false,
	}
}

// SetSelectWhere 设置当前数据查询条件
func (t *Grid) SetSelectWhere(where WhereClosure) *Grid {
	t.where = where
	return t
}

// GetSelectWhere 获取当前数据查询条件
func (t *Grid) GetSelectWhere() WhereClosure {
	return t.where
}

// EnableDelete 允许删除
func (t *Grid) EnableDelete() *Grid {
	t.DeleteAble = true
	return t
}

// EnableEdit 允许编辑
func (t *Grid) EnableEdit() *Grid {
	t.EditAble = true
	return t
}

// EnableDetail 允许详情
func (t *Grid) EnableDetail() *Grid {
	t.DetailAble = true
	return t
}

// DisableDelete 禁用删除
func (t *Grid) DisableDelete() *Grid {
	t.DeleteAble = false
	return t
}

// DisableEdit 禁用编辑
func (t *Grid) DisableEdit() *Grid {
	t.EditAble = false
	return t
}

// DisableDetail 禁用详情
func (t *Grid) DisableDetail() *Grid {
	t.DetailAble = false
	return t
}

// SetTips 设置提示
func (t *Grid) SetTips(tips string) *Grid {
	t.Tips = tips
	return t
}

// SetInitSearch 设置提示
func (t *Grid) SetInitSearch(optionGroups []OptionGroup, defaultValue string) *Grid {
	t.InitSearch = InitSearch{
		OptionGroups: optionGroups,
		DefaultValue: defaultValue,
	}
	return t
}

// Search 设置当前字段可搜素
func (t *Grid) Search(searchFields ...SearchField) *Grid {
	if len(searchFields) > 0 {
		t.SearchFields = append(t.SearchFields, searchFields[0])
	} else {
		t.SearchFields = append(t.SearchFields, SearchField{
			Field:    t.fieldList[t.curFieldListIndex].Field,
			ShowName: t.fieldList[t.curFieldListIndex].ShowName,
			Notice:   fmt.Sprintf("%s %s", "请输入", t.fieldList[t.curFieldListIndex].ShowName),
		})
	}
	return t
}

// AddField 添加字段显示
func (t *Grid) AddField(head string, field string) *Grid {
	t.fieldList = append(t.fieldList, TableField{
		Field:    field,
		ShowName: head,
	})
	t.fieldFunMap[strings.ToUpper(strings.Replace(field, "_", "", 10))] = nil
	t.fieldSort = append(t.fieldSort, field)
	t.curFieldListIndex++
	return t
}

// FieldSortable 当前字段可排序
func (t *Grid) FieldSortable() *Grid {
	t.fieldList[t.curFieldListIndex].Sortable = true
	return t
}

// DisableAddData 禁用添加按钮
func (t *Grid) DisableAddData() *Grid {
	t.DataAdd = false
	return t
}

// Display 自定义当中数据展示
func (t *Grid) Display(displayFun FieldDisplayFun) *Grid {
	t.fieldFunMap[strings.ToUpper(t.fieldList[t.curFieldListIndex].Field)] = displayFun
	return t
}

// Label 以标签形式展示数据
func (t *Grid) Label() *Grid {
	t.fieldLabelMap[strings.ToUpper(t.fieldList[t.curFieldListIndex].Field)] = true
	return t
}

// Image 以图片展示数据
func (t *Grid) Image() *Grid {
	t.fieldImgMap[strings.ToUpper(t.fieldList[t.curFieldListIndex].Field)] = true
	return t
}

// Icon 以icon展示数据
func (t *Grid) Icon() *Grid {
	t.fieldIconMap[strings.ToUpper(t.fieldList[t.curFieldListIndex].Field)] = true
	return t
}

// BaseUrl 获取当前请求的基本uril，搜索用
func (t *Grid) BaseUrl() string {
	return strings.Split(t.context.Request.RequestURI, "?")[0]
}

// CurrentUrl 获取当前url 页面分页用
func (t *Grid) CurrentUrl(params ...any) string {
	if len(params) > 0 {
		par := map[string]string{
			"_ts": "s",
		}
		for k, v := range t.context.Request.URL.Query() {
			if k != "_pjax" {
				par[k] = v[0]
			}
		}
		switch params[1].(type) {
		case int:
			par[params[0].(string)] = strconv.Itoa(params[1].(int))
		case string:
			par[params[0].(string)] = params[1].(string)
		}
		var url []string
		for k, v := range par {
			url = append(url, k+"="+v)
		}
		return t.BaseUrl() + "?" + strings.Join(url, "&")
	} else {
		return t.context.Request.RequestURI
	}
}

// SearchValue 当前搜索的字段信息，页面回显用
func (t *Grid) SearchValue(params any) any {
	return t.context.DefaultQuery(fmt.Sprintf("_search_params_%s", params), "")
}

// Show 数据显现
func (f *Field) Show() any {
	res := f.Display(FieldModel{
		ID:    f.Row["id"],
		Value: f.Value,
		Row:   f.Row,
	})
	return res
}

// GetTable 获取显现的table对象
func (t *Grid) GetTable(pageData handler.PageData) *Grid {
	t.PageData = pageData
	var response [][]Field
	if t.PageData.DataList != nil {
		res := t.PageData.DataList.([]map[string]any)
		for _, v := range res {
			var tmp []Field
			var value any
			for _, key := range t.fieldSort {
				if key == "created_at" || key == "updated_at" || key == "deleted_at" {
					if v[key] != nil {
						ts, _ := time.ParseInLocation(datetime.ParseTimeTemplate, v[key].(string), time.Local)
						value = ts.Format(datetime.DateTemplate)
					}
				} else {
					value = v[key]
				}
				tmp = append(tmp, Field{
					Field:   key,
					Row:     v,
					Value:   value,
					Display: t.fieldFunMap[strings.ToUpper(key)],
					Label:   t.fieldLabelMap[strings.ToUpper(key)],
					Icon:    t.fieldIconMap[strings.ToUpper(key)],
					Image:   t.fieldImgMap[strings.ToUpper(key)],
				})
			}
			response = append(response, tmp)
		}
		t.PageData.DataList = response
		if len(t.SearchFields) > 0 {
			t.HasSearch = true
		}
	}
	return t
}

// GetField 获取字段
func (t *Grid) GetField() FieldList {
	return t.fieldList
}
