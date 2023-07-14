package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
)

/*************************************** 下划线json ***************************************/
/*
func main() {
	type Person struct {
		HelloWold       string
		LightWeightBaby string
	}
	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
	res, _ := json.Marshal(jsonconv.JsonSnakeCase{a})
	fmt.Printf("%s", res)
}
*/

func init() {

}

type JsonSnakeCase struct {
	Value interface{}
}

func (c JsonSnakeCase) UnmarshalJSON(b []byte) error {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	converted := keyMatchRegex.ReplaceAllFunc(
		b,
		func(match []byte) []byte {
			return []byte(Case2Camel(string(match)))
		},
	)
	return json.Unmarshal(converted, c.Value)
}
func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

/*************************************** 驼峰json ***************************************/
/*
func main() {
	type Person struct {
		HelloWold       string `json:"hello_wold"`
		LightWeightBaby string `json:"light_weight_baby"`
	}
	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
	res, _ := json.Marshal(jsonconv.JsonCamelCase{a})
	fmt.Printf("%s", res)
}
*/

type JsonCamelCase struct {
	Value interface{}
}

func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := UcFirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// Camel2Case /*************************************** 其他方法 ***************************************/

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
func TransJsonString(source []byte) []byte {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	return keyMatchRegex.ReplaceAllFunc(
		source,
		func(match []byte) []byte {
			return []byte(Case2Camel(string(match)))
		},
	)
}
