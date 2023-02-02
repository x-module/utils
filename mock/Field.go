package mock

import (
	"fmt"
	"strconv"
)

const (
	Lower  = "lower"  // 小写
	Upper  = "upper"  // 大写
	Number = "number" // 数字
	Symbol = "number" // 符号
)

func Boolean(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s|1", field)] = "@boolean"
	return rule
}

func Integer(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@integer(%d,%d)", min, max)
	return rule
}

func Float(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@float")
	return rule
}
func Character(field string, charType string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@character(%s)", charType)
	return rule
}
func String(field string, charType string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@string(%s,%d,%d)", charType, min, max)
	return rule
}
func Range(field string, start int, stop int, step int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@range(%d,%d,%d)", start, stop, step)
	return rule
}
func Date(field string, format string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@date(%s)", format)
	return rule
}
func Time(field string, format string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@time(%s)", format)
	return rule
}
func Datetime(field string, format string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@datetime(%s)", format)
	return rule
}
func Now(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@Now")
	return rule
}
func Image(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@image")
	return rule
}
func Sentence(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@sentence")
	return rule
}

func Paragraph(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@paragraph(%d,%d)", min, max)
	return rule
}
func Title(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@title(%d,%d)", min, max)
	return rule
}

// CParagraph 中文段落
func CParagraph(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@cparagraph(%d,%d)", min, max)
	return rule
}

// CSentence 中文句子
func CSentence(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@csentence(%d,%d)", min, max)
	return rule
}

// CTitle 中文标题
func CTitle(field string, min int, max int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@ctitle(%d,%d)", min, max)
	return rule
}

// func DataImage(field string, rule map[string]interface{}) map[string]interface{} {
//	//rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@dataImage")
//	return rule
// }

// First 英文姓
func First(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@first")
	return rule
}

// Last 英文名
func Last(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@last")
	return rule
}

// Name 英文姓名
func Name(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@name")
	return rule
}

// CFirst 中文姓
func CFirst(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@cfirst")
	return rule
}

// CLast 中文名
func CLast(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@clast")
	return rule
}

// CName 中文姓名
func CName(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@cname")
	return rule
}

func Url(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@url")
	return rule
}

func Domain(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@domain")
	return rule
}
func Email(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@email")
	return rule
}
func Region(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@region")
	return rule
}

// Province 省
func Province(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@province")
	return rule
}
func Protocol(field string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@protocol")
	return rule
}
func ArrayRand(field string, count int, array []string, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s|%d", field, count)] = array
	return rule
}

func Object(field string, count int, object any, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s|%d", field, count)] = object
	return rule
}

func LengthRand(field string, from int, to int, count int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s|%d-%d", field, from, to)] = count
	return rule
}

func Increment(field string, step int, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@increment(%d)", step)
	return rule
}

func City(field string, prefix bool, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@city(%s)", strconv.FormatBool(prefix))
	return rule
}

// County 区县
func County(field string, prefix bool, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = fmt.Sprintf("@county(%s)", strconv.FormatBool(prefix))
	return rule
}
func Regexp(field string, regexp any, rule map[string]any) map[string]any {
	rule[fmt.Sprintf("%s", field)] = regexp
	return rule
}
