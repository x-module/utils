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
	"time"
)

type DateTime struct {
	time.Time
}

func (u *DateTime) UnmarshalJSON(data []byte) error {
	*u = DateTime{
		//Time: carbon.Parse(string(data)).ToStdTime(),
		Time: time.Now(),
	}
	return nil
}
func (u *DateTime) MarshalJSON() ([]byte, error) {
	dateTime := carbon.FromStdTime(u.Time).ToDateTimeString()
	return json.Marshal(dateTime)
}

func (u *DateTime) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	*u = DateTime{
		Time: dbValue.(time.Time),
	}
	//err = field.Set(ctx, dst, DateTime{
	//	Time: carbon.FromStdTime(time.Now()).ToStdTime(),
	//})
	//carbon.FromStdTime(dbValue.(time.Time)).ToDateTimeString())
	//
	////
	//fieldValue := reflect.New(field.FieldType)
	//fmt.Println(dbValue)
	//if dbValue != nil {
	//	var bytes []byte
	//	switch v := dbValue.(type) {
	//	case []byte:
	//		bytes = v
	//	case string:
	//		bytes = []byte(v)
	//		fmt.Println("-----------string---------------")
	//	case time.Time:
	//		bytes = []byte(carbon.FromStdTime(dbValue.(time.Time)).ToDateTimeString())
	//	default:
	//		fmt.Println("===========::", dbValue.(DateTime).Time.String())
	//		return fmt.Errorf("-----failed to unmarshal JSONB value: %#v", dbValue)
	//	}
	//	err = json.Unmarshal(bytes, fieldValue.Interface())
	//}
	//
	//field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (u *DateTime) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return carbon.FromStdTime(fieldValue.(DateTime).Time).ToDateTimeString(), nil
}
