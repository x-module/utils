/**
* Created by Goland
* @file response.go
* @version: 1.0.0
* @author 李锦 <lijin@shihuituan.com>
* @date 2021/11/10 6:06 下午
* @desc 请求响应
 */

package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xmodule/utils/global"
	"io/ioutil"
	"net/http"
)

type Response struct {
	response *http.Response
	url      string
	body     []byte
	time     int64
}

// StatusCode 获取响应状态码
func (r *Response) StatusCode() int {
	if r != nil {
		return r.response.StatusCode
	}
	return 0
}

// Time 请求耗时
func (r *Response) Time() string {
	if r != nil {
		return fmt.Sprintf("%dms", r.time)
	}
	return "0ms"
}

// Url 请求链接
func (r *Response) Url() string {
	if r != nil {
		return r.url
	}
	return ""
}

// Headers 响应header
func (r *Response) Headers() http.Header {
	if r != nil {
		return r.response.Header
	}
	return nil
}

// Cookies 响应cookie
func (r *Response) Cookies() []*http.Cookie {
	if r != nil {
		return r.response.Cookies()
	}
	return []*http.Cookie{}
}

// Response 获取请求响应体
func (r *Response) Response() *http.Response {
	if r != nil {
		return r.response
	}
	return nil
}

// Body 响应体
func (r *Response) Body() ([]byte, error) {
	if r == nil {
		return []byte{}, errors.New("response is null")
	}
	if r.response == nil || r.response.Body == nil {
		return nil, errors.New("response or body is nil")
	}
	defer r.response.Body.Close()
	b, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return []byte{}, errors.New("response body is null")
	}
	r.body = b
	return b, nil
}

// Content 响应内容
func (r *Response) Content() (string, error) {
	b, err := r.Body()
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (r *Response) Json(T any) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, T)
	if err != nil {
		return err
	}
	return nil
}

type Result struct {
	Code global.ErrCode `json:"code"`
	Msg  string         `json:"msg"`
	Data string         `json:"data"`
}

func (r *Response) Result() (Result, error) {
	body, err := r.Body()
	if err != nil {
		return Result{}, err
	}
	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Result{}, err
	}
	return result, nil
}

func (r *Response) JsonReturn(T any) (string, error) {
	body, err := r.Body()
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &T)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

// Close 关闭连接
func (r *Response) Close() error {
	if r != nil {
		return r.response.Body.Close()
	}
	return nil
}
