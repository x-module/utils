/**
 * Created by Goland.
 * @file   display.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/11 16:17
 * @desc   display.go
 */

package utils

import (
	"fmt"
	"github.com/x-module/utils/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 响应结构体
type ResponseData struct {
	Code fmt.Stringer `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

// ApiResponse  异常通知
func ApiResponse(context *gin.Context, errorCode fmt.Stringer, data ...any) {
	response := ResponseData{
		Code: errorCode,
		Msg:  errorCode.String(),
	}
	if len(data) > 0 {
		response.Data = data[0]
	} else if global.Success == errorCode {
		response.Data = "success"
	}
	context.JSON(http.StatusOK, response)
}

// WebResponse 异常通知
func WebResponse(context *gin.Context, errorCode any, data ...any) {
	msg := ""
	var errCode global.ErrCode = 201
	if code, ok := errorCode.(global.ErrCode); ok {
		msg = code.String()
		errCode = code
	} else if err, ok := errorCode.(error); ok {
		msg = err.Error()
	} else {
		msg = errorCode.(string)
	}
	response := ResponseData{
		Code: errCode,
		Msg:  msg,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	context.HTML(http.StatusOK, "notice.html", gin.H{"notice": response})
}
