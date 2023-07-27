/**
 * Created by Goland.
 * @file   base.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/19 15:02
 * @desc   base.go
 */

package validation

import (
	"encoding/json"
	"fmt"
	"github.com/x-module/module/utils"
	"github.com/x-module/utils/utils/xlog"
)

// BaseValidation 基类
type BaseValidation[T any] struct {
	validate T
}

func NewBaseValidation[T any](validate T) *BaseValidation[T] {
	return &BaseValidation[T]{
		validate: validate,
	}
}

// BaseValidationParams 基础参数
type BaseValidationParams struct {
}

func (i *BaseValidation[T]) Validation() error {
	params, _ := json.Marshal(i.validate)
	if err := utils.Validation(params, i.validate); err != nil {
		fmt.Println("=====================================================")
		fmt.Println("request:", string(params))
		fmt.Println("=====================================================")
		xlog.Logger.Warn(err)
		return err
	}
	return nil
}
