/**
 * Created by GoLand
 * @file   data.go
* @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/5/7 13:32
 * @desc   data.go
*/

package utils

import (
	"errors"
	"fmt"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/golang-module/carbon"
	"math"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

// DateFormat
const DateFormat = "2006-01-02"

// MonthFormat 月份模板
const MonthFormat = "2006-01"

const (
	// DateTimeTemplate 日期时间模板
	DateTimeTemplate  = "2006-01-02 15:04:05"
	ZeroTimeTemplate  = "2006-01-02 00:00:00"
	DateTemplate      = "2006-01-02"
	OnlyTimeTemplate  = "15:04:05"
	ParseTimeTemplate = "2006-01-02T15:04:05+08:00"
	YearMonthTemplate = "2006-01"
)

// TransDateArr 获取日期区间
func TransDateArr(start string, end string) []string {
	var dateList []string
	st, _ := time.ParseInLocation(DateTemplate, start, time.Local)
	et, _ := time.ParseInLocation(DateTemplate, end, time.Local)
	startUnix := st.Unix()
	endUnix := et.Unix()
	for i := startUnix; i <= endUnix; i = i + 86400 {
		dateList = append(dateList, time.Unix(i, 0).Format(DateTemplate))
	}
	return dateList
}

// GetTimeDuration 获取时间区间
func GetTimeDuration(startTime int64, endTime int64) int64 {
	// st, _ := time.ParseInLocation(Template, startTime, time.Local)
	// en, _ := time.ParseInLocation(Template, endTime, time.Local)
	// return en.Unix() - st.Unix()
	return endTime - startTime
}

// GetDayStringSub 获取两个日期的天数之差
func GetDayStringSub(startDate string, endDate string) int {
	st, _ := time.Parse(DateTimeTemplate, startDate)
	et, _ := time.Parse(DateTimeTemplate, endDate)
	std := GetYearDay(st)
	end := GetYearDay(et)
	return end - std
}

// GetDaySub 获取日期差天数
func GetDaySub(startDate time.Time, endDate time.Time) int {
	return int(carbon.Time2Carbon(startDate).DiffAbsInDays(carbon.Time2Carbon(endDate)))
}

// GetYearDay 获取一年中的第几天
func GetYearDay(date time.Time) int {
	var days int = 0
	var y, m, d int = date.Year(), int(date.Month()), date.Day()
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	return days
}

// GetDayRangeList 获取日期区间数据
func GetDayRangeList(startDate string, endDate string) ([]string, error) {
	st, _ := time.ParseInLocation(DateTemplate, startDate, time.Local)
	if st.Unix() < 0 {
		xlog.Logger.Warning("error date format,expect[2006-01-02] , input[" + startDate + "]")
		return nil, errors.New("error date format,expect[2006-01-02] , input[" + startDate + "]")
	}
	et, _ := time.ParseInLocation(DateTemplate, endDate, time.Local)
	if et.Unix() < 0 {
		xlog.Logger.Warning("error date format,expect[2006-01-02] , input[" + endDate + "]")
		return nil, errors.New("error date format,expect[2006-01-02] , input[" + endDate + "]")
	}
	if et.Unix() < st.Unix() {
		xlog.Logger.Warning("end time must after start time[" + startDate + "]")
		return nil, errors.New("end time must after start time[" + startDate + "]")
	}
	var dateMap []string
	for i := 0; ; i++ {
		currentTimes := st.Unix() + int64(86400*i)
		if currentTimes <= et.Unix() {
			dateMap = append(dateMap, time.Unix(currentTimes, 0).Format(DateTemplate))
		} else {
			break
		}
	}
	return dateMap, nil
}

// GetWeekRangeList 将开始时间和结束时间分割为周为单位
func GetWeekRangeList(startDate string, endDate string) []string {
	l, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation(DateTemplate, startDate, l)
	endTime, _ := time.ParseInLocation(DateTemplate, endDate, l)
	weekDate := make([]WeekDate, 0)
	diffDuration := endTime.Sub(startTime)
	days := int(math.Ceil(float64(diffDuration/(time.Hour*24)))) + 1
	currentWeekDate := WeekDate{}
	currentWeekDate.WeekTh = weekByDate(endTime)
	currentWeekDate.EndTime = endTime
	currentWeekDay := int(endTime.Weekday())
	if currentWeekDay == 0 {
		currentWeekDay = 7
	}
	currentWeekDate.StartTime = endTime.AddDate(0, 0, -currentWeekDay+1)
	nextWeekEndTime := currentWeekDate.StartTime
	weekDate = append(weekDate, currentWeekDate)

	for i := 0; i < (days-currentWeekDay)/7; i++ {
		weekData := WeekDate{}
		weekData.EndTime = nextWeekEndTime
		weekData.StartTime = nextWeekEndTime.AddDate(0, 0, -7)
		weekData.WeekTh = weekByDate(weekData.StartTime)
		nextWeekEndTime = weekData.StartTime
		weekDate = append(weekDate, weekData)
	}

	if lastDays := (days - currentWeekDay) % 7; lastDays > 0 {
		lastData := WeekDate{}
		lastData.EndTime = nextWeekEndTime
		lastData.StartTime = nextWeekEndTime.AddDate(0, 0, -lastDays)
		lastData.WeekTh = weekByDate(lastData.StartTime)
		weekDate = append(weekDate, lastData)
	}

	var weekList []string
	for _, week := range weekDate {
		weekList = append(weekList, week.WeekTh)
	}
	return reverse(weekList)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// WeekByDate 判断时间是当年的第几周
func weekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	// 今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("%d%d", t.Year(), week)
}

type WeekDate struct {
	WeekTh    string
	StartTime time.Time
	EndTime   time.Time
}

// GetMonthRangeList 月份区间
func GetMonthRangeList(startDate string, endDate string) []string {
	l, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation(DateTemplate, startDate, l)
	endTime, _ := time.ParseInLocation(DateTemplate, endDate, l)
	monthList := []string{
		startTime.Format(MonthFormat),
	}

	if startTime.Format(MonthFormat) == endTime.Format(MonthFormat) {
		return monthList
	}

	for i := 1; ; i++ {
		month := startTime.AddDate(0, i, 0).Format(MonthFormat)
		monthList = append(monthList, month)
		if month == endTime.Format(MonthFormat) {
			break
		}
	}
	return monthList
}

// TransDateFormat 时间日期转换
func TransDateFormat(dateTime string, template string) (string, error) {
	if strings.TrimSpace(dateTime) == "" {
		return "", nil
	}
	ts, err := now.Parse(dateTime)
	if err != nil {
		return "", err
	}
	return ts.Format(template), nil
}

// ConvertUTC UTC 时间转换
func ConvertUTC(t time.Time) string {
	// t1, _ := time.Parse("2006-01-02T15:04:05Z", t)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") // 上海
	return t.In(cstSh).Format("2006-01-02 15:04:05")
}

// ConvertStringUTC UTC 时间转换
func ConvertStringUTC(t string) string {
	t1, _ := time.Parse("2006-01-02T15:04:05Z", t)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") // 上海
	return t1.In(cstSh).Format("2006-01-02 15:04:05")
}
