/**
* Created by Goland
* @file request.go
* @version: 1.0.0
* @author 李锦 <lijin@shihuituan.com>
* @date 2021/11/10 2:52 下午
* @desc 网络请求
 */

package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Request struct {
	client           *http.Client
	debug            bool
	transport        *http.Transport
	headers          map[string]string
	cookies          map[string]string
	method           string
	url              string
	data             any
	disableKeepAlive bool
	timeout          time.Duration
}

func NewRequest() *Request {
	return new(Request)
}

func (r *Request) SetTimeOut(time time.Duration) *Request {
	r.timeout = time
	return r
}
func (r *Request) Debug(debug ...bool) *Request {
	if len(debug) > 0 {
		r.debug = debug[0]
	} else {
		r.debug = true
	}
	return r
}

// SetTimeout 设置超时时间
func (r *Request) SetTimeout(d time.Duration) *Request {
	r.timeout = d
	return r
}

// 线上请求信息-调试模式
func (r *Request) debugTrace() {
	if r.debug {
		log.SetPrefix("[*] ")
		log.Println("============================= HttpUtils =============================")
		log.Printf("Request: %s ", r.url)
		log.Printf("Method : %s ", r.method)
		log.Printf("Headers: %v", r.headers)
		log.Printf("Cookies: %v", r.cookies)
		log.Printf("Timeout: %d", r.timeout)
		log.Printf("ReqBody: %+v", r.data)
		log.Println("======================================================================")
	}
}

// Json Json模式请求
func (r *Request) Json() *Request {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers["Content-Type"] = "application/json"

	return r
}

func (r *Request) SetTransport(v *http.Transport) *Request {
	r.transport = v
	return r
}

// 获取 Transport
func (r *Request) getTransport() http.RoundTripper {
	if r.transport == nil {
		return http.DefaultTransport
	}
	r.transport.DisableKeepAlives = r.disableKeepAlive
	return http.RoundTripper(r.transport)
}

// DisableKeepAlive 禁用长链接
func (r *Request) DisableKeepAlive(keep bool) *Request {
	r.disableKeepAlive = keep
	return r
}

// 构建客户端
func (r *Request) buildClient() *http.Client {
	if r.client == nil {
		return &http.Client{
			Transport: r.getTransport(),
			// Jar:           r.jar,
			// CheckRedirect: r.checkRedirect,
			Timeout: time.Second * r.timeout,
		}
	}
	return r.client
}

// 构建请求的url
func (r *Request) buildUrl(url string, params ...any) (string, error) {
	query, err := r.parseUrl(url)
	if err != nil {
		return url, err
	}
	if len(params) > 0 && params[0] != nil { // 确定有参数
		paramsValue := ""
		paramsType := reflect.TypeOf(params[0]).String()
		switch paramsType {
		case "map[string]interface {}":
			for key, value := range params[0].(map[string]any) {
				if reflect.TypeOf(value).String() == "string" {
					paramsValue = value.(string)
				} else {
					b, err := json.Marshal(value)
					if err != nil {
						return url, err
					}
					paramsValue = string(b)
				}
				query = append(query, fmt.Sprintf("%s=%s", key, paramsValue))
			}
		case "string":
			param := params[0].(string)
			if param != "" {
				query = append(query, param)
			}
		default:
			return url, errors.New("does not support  data type.")
		}
	}

	list := strings.Split(url, "?")
	if len(list) > 1 {
		return fmt.Sprintf("%s?%s", list[0], strings.Join(query, "&")), nil
	}
	return fmt.Sprintf("%s?%s", url, strings.Join(query, "&")), nil
}

// 解析url
func (r *Request) parseUrl(url string) ([]string, error) {
	urlList := strings.Split(url, "?")
	if len(urlList) < 2 { // 不带参数的纯种url
		return make([]string, 0), nil
	}
	query := make([]string, 0)
	for _, value := range strings.Split(urlList[1], "&") {
		v := strings.Split(value, "=")
		if len(v) < 2 {
			return make([]string, 0), errors.New("query parameter error")
		}
		query = append(query, fmt.Sprintf("%s=%s", v[0], v[1]))
	}
	return query, nil
}
func (r *Request) runTime(n int64, resp *Response) {
	end := time.Now().UnixNano() / 1e6
	resp.time = end - n
}

func (r *Request) request(url string, method string, params ...any) (*Response, error) {
	if url == "" || method == "" {
		return nil, errors.New("parameter method and url is required")
	}
	response := &Response{}
	// Start time
	start := time.Now().UnixNano() / 1e6
	defer r.runTime(start, response)

	if len(params) > 0 {
		r.data = params[0]
	} else {
		r.data = ""
	}
	r.client = r.buildClient()
	r.method = strings.ToUpper(method)
	if r.method == http.MethodGet && len(params) > 0 { // GET方式拼凑url
		requestUrl, err := r.buildUrl(url, params...)
		if err != nil {
			return nil, err
		}
		r.url = requestUrl
	} else if r.method == http.MethodDelete && len(params) > 0 { // DELETE方式拼凑url
		requestUrl, err := r.buildUrl(url, params...)
		if err != nil {
			return nil, err
		}
		r.url = requestUrl
	} else {
		r.url = url
	}
	r.debugTrace()
	var body io.Reader
	var err error
	if len(params) > 0 {
		body, err = r.buildBody(params...)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(r.method, r.url, body)
	if err != nil {
		return nil, err
	}
	r.initCookies(req)
	r.initHeaders(req)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Build Response
	response.response = resp
	response.url = url
	return response, nil
}

// SetHeaders 设置请求的header
func (r *Request) SetHeaders(headers map[string]string) *Request {
	if headers != nil && len(headers) > 0 {
		r.headers = headers
	}
	return r
}

// 添加请求的header
func (r *Request) initHeaders(request *http.Request) *Request {
	if r.headers != nil && len(r.headers) > 0 {
		for key, value := range r.headers {
			request.Header.Set(key, value)
		}
	}
	return r
}

// 添加请求的cookie
func (r *Request) initCookies(request *http.Request) {
	for key, value := range r.cookies {
		request.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
}

// SetCookies 设置cookie
func (r *Request) SetCookies(cookies map[string]string) *Request {
	if cookies != nil && len(cookies) > 0 {
		r.cookies = cookies
	}
	return r
}

// 构建请求消息体
func (r *Request) buildBody(params ...any) (io.Reader, error) {
	if r.method == http.MethodGet || r.method == http.MethodDelete {
		return nil, nil
	}
	data := make([]string, 0)
	paramsType := reflect.TypeOf(params[0]).String()
	if paramsType == "string" {
		return strings.NewReader(params[0].(string)), nil
	} else if paramsType == "map[string]interface {}" {
		if r.headers["Content-Type"] == "application/json" {
			payloadBytes, _ := json.Marshal(params[0])
			return strings.NewReader(string(payloadBytes)), nil
		} else {
			for key, value := range params[0].(map[string]any) {
				if param, ok := value.(string); ok {
					data = append(data, fmt.Sprintf("%s=%s", key, param))
					continue
				}
				b, err := json.Marshal(value)
				if err != nil {
					return nil, err
				}
				data = append(data, fmt.Sprintf("%s=%s", key, string(b)))
			}
		}
	} else {
		payloadBytes, _ := json.Marshal(params[0])
		return bytes.NewReader(payloadBytes), nil
	}
	return strings.NewReader(strings.Join(data, "&")), nil
}

func (r *Request) Get(url string, params ...any) (*Response, error) {
	return r.request(url, http.MethodGet, params...)
}
func (r *Request) Delete(url string, params ...any) (*Response, error) {
	return r.request(url, http.MethodDelete, params...)
}

func (r *Request) Post(url string, params ...any) (*Response, error) {
	return r.request(url, http.MethodPost, params...)
}
