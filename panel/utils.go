/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/2 15:41
 * @desc   utils.go
 */

package panel

import (
	"encoding/json"
	"github.com/go-utils-module/utils/utils/datetime"
	"time"
)

// ConvertModeToMap 转换model 为map格式
func ConvertModeToMap(data any) map[string]any {
	s, _ := json.Marshal(data)
	var mapData map[string]any
	_ = json.Unmarshal(s, &mapData)
	for field, value := range mapData {
		if field == "created_at" || field == "updated_at" || field == "deleted_at" {
			if value != nil {
				t, _ := time.ParseInLocation(datetime.ParseTimeTemplate, value.(string), time.Local)
				mapData[field] = t.Format(datetime.DateTimeTemplate)
			}
		}
	}
	return mapData
}
