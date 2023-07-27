/**
 * Created by Goland.
 * @file   datetime.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/17 12:47
 * @desc   datetime.go
 */

package types

import (
	"context"
	"encoding/json"
	"github.com/golang-module/carbon"
	"gorm.io/gorm/schema"
	"reflect"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (u *DateTime) UnmarshalJSON(data []byte) error {
	timeStr := strings.ReplaceAll(string(data), "\"", "")
	*u = DateTime{
		Time: carbon.Parse(timeStr).ToStdTime(),
	}
	return nil
}
func (u *DateTime) MarshalJSON() ([]byte, error) {
	dateTime := carbon.FromStdTime(u.Time).ToDateTimeString()
	return json.Marshal(dateTime)
	//return []byte(dateTime), nil
}

func (u *DateTime) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	*u = DateTime{
		Time: dbValue.(time.Time),
	}
	return
}

func (u *DateTime) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return carbon.FromStdTime(fieldValue.(DateTime).Time).ToDateTimeString(), nil
}
