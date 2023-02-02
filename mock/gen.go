package mock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

type MockResult struct {
	Data any `json:"data"`
}

func MockData(fields map[string]any, count int, target any, path string) error {
	extra.RegisterFuzzyDecoders()
	key := fmt.Sprintf("%s|%d", "data", count)
	paramsBite := map[string][]map[string]any{
		key: {
			fields,
		},
	}
	params, _ := json.Marshal(paramsBite)
	reg := string(params)
	reg = strings.Replace(reg, "\"/", "/", -1)
	reg = strings.Replace(reg, "/\"", "/", -1)
	reg = strings.Replace(reg, "\\\\", "\\", -1)
	code := `var Mock = require("mockjs")
let data = Mock.mock(#code#)
console.log(JSON.stringify(data))`
	code = strings.Replace(code, "#code#", reg, 1)
	jsFile := path + "/temp_mock.js"
	// jsFile = getCurrentPath()
	_ = ioutil.WriteFile(jsFile, []byte(code), 0666)
	command := "/opt/homebrew/bin/node " + jsFile
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()

	_ = os.Remove(jsFile)
	if err != nil {
		return err
	}
	var mockResult MockResult
	_ = json.Unmarshal(data, &mockResult)
	result, _ := json.Marshal(mockResult.Data)
	_ = jsoniter.Unmarshal(result, target)
	return nil
}

// 获取当前执行程序所在的绝对路径
func getCurrentAbPathByExecutable() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func MockRequestData(fields map[string]any, dataKey string, count int) (string, error) {
	path := getCurrentAbPathByExecutable()
	key := fmt.Sprintf("%s|%d", dataKey, count)
	// rule, _ := json.Marshal(fields)
	paramsBite := map[string][]map[string]any{
		key: {
			fields,
		},
	}
	params, _ := json.Marshal(paramsBite)
	reg := string(params)
	reg = strings.Replace(reg, "\"/", "/", -1)
	reg = strings.Replace(reg, "/\"", "/", -1)
	reg = strings.Replace(reg, "\\\\", "\\", -1)
	code := `var Mock = require("mockjs")
let data = Mock.mock(#code#)
console.log(JSON.stringify(data))`
	code = strings.Replace(code, "#code#", reg, 1)
	jsFile := path + "/script/temp_mock.js"
	_ = ioutil.WriteFile(jsFile, []byte(code), 0666)
	command := "/opt/homebrew/bin/node " + jsFile
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()
	_ = os.Remove(jsFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
