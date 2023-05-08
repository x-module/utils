/**
 * Created by PhpStorm.
 * @file   generate.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/4/23 06:48
 * @desc   generate.go
 */

package model

import (
	"github.com/go-xmodule/utils/handler"
)

func GetTableInfo(database string, table string) (string, []handler.Field, error) {
	tableInfo, err := handler.DBHandler.GetTableComment(database, table)
	if err != nil {
		return "", nil, err
	}

	fields, err := handler.DBHandler.GetTableFieldComment(database, table)
	if err != nil {
		return "", nil, err
	}
	return tableInfo, fields, nil
}
