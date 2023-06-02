/**
 * Created by Goland.
 * @file   validator.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/3/2 16:31
 * @desc   validator.go
 */

package panel

import (
	"errors"
	"github.com/x-module/utils/handler"
	"github.com/x-module/utils/utils/xlog"
	"strings"
)

// ValidatorItem Validator 数据验证体定义
type ValidatorItem struct {
	ValidatorType ValidatorType
	ValidatorMsg  string
}

const (
	DataAdd ActionType = iota
	DataEdit
)

// ValidatorList 待验证数据集合
// 表:操作类型:字段:[]规则
var ValidatorList = map[string]map[ActionType]map[string][]ValidatorItem{}

type Validator struct {
	baseModel *handler.CommonModel
	model     handler.ModelAction
}

func NewValidator(model handler.ModelAction) *Validator {
	return &Validator{
		model: model,
	}
}

// SetBaseModel 设置基础model
func (v *Validator) SetBaseModel(model *handler.CommonModel) *Validator {
	v.baseModel = model
	return v
}

// DataValidator 添加验证方法 --内部
func (v *Validator) DataValidator(actionType ActionType, params map[string]string) error {
	var errorList []string
	if actionList, ok := ValidatorList[v.model.TableName()]; ok {
		if fieldList, ok := actionList[actionType]; ok { // 没有当前操作类型的验证
			for field, rules := range fieldList {
				value, exist := params[field]
				if !exist || value == "" { // 当前字段未设置或设置为空
					for _, rule := range rules {
						if rule.ValidatorType == ValidatorRequire { // 必填规则
							errorList = append(errorList, rule.ValidatorMsg)
						}
					}
				} else { // 字段不为空，其他规则验证
					for _, rule := range rules {
						if rule.ValidatorType == ValidatorUnique { // 必填规则
							v.checkDataExist(field, value, params)
							// errorList = append(errorList, errors.New(rule.ValidatorMsg))
						}
					}
				}
			}
		}
	}
	if len(errorList) > 0 {
		xlog.Logger.Warning(strings.Join(errorList, ","))
		return errors.New(strings.Join(errorList, ","))
	} else {
		return nil
	}
}

func (v *Validator) checkDataExist(field string, value string, params map[string]string) {
	// not := map[string]interface{}{}
	// if _, exist := params["id"]; exist {
	// 	not["id"] = 30
	// }
	// v.baseModel.First(map[string]interface{}{
	// 	field: value,
	// }, not)
	// v.baseModel.GetModel()
}

// AddValidator 添加验证规则--内部方法
func (v *Validator) AddValidator(actionType ActionType, formField FormField, validatorItem []ValidatorItem) {
	tableName := v.model.TableName()
	if _, ok := ValidatorList[tableName]; !ok {
		ValidatorList[tableName] = map[ActionType]map[string][]ValidatorItem{
			actionType: map[string][]ValidatorItem{
				formField.Field: validatorItem,
			},
		}
	} else if _, ok := ValidatorList[tableName][actionType]; !ok { // 没有当前操作类型的验证
		ValidatorList[tableName][actionType] = map[string][]ValidatorItem{
			formField.Field: validatorItem,
		}
	} else if _, ok := ValidatorList[tableName][actionType][formField.Field]; !ok { // 没有当前操作的字段
		ValidatorList[tableName][actionType][formField.Field] = validatorItem
	} else {
		ValidatorList[tableName][actionType][formField.Field] = append(ValidatorList[tableName][actionType][formField.Field], validatorItem...)
	}
}
