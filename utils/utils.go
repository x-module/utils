/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/2 15:31
 * @desc   utils.go
 */

package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/x-module/utils/global"
	"github.com/x-module/utils/utils/xlog"
	"gorm.io/gorm"
	"net"
	"os"
	"strconv"
	"strings"
)

func Success(status int) bool {
	return global.ErrCode(status) == global.Success
}
func JsonString(params any) string {
	b, _ := json.Marshal(params)
	return string(b)
}
func CatchErr(err error, errCode fmt.Stringer) bool {
	return HasErr(err, errCode)
}
func DebugResponse(err error, errCode fmt.Stringer, debugMsg string) {
	if !HasErr(err, errCode) {
		xlog.Logger.Debug(debugMsg)
	}
}

func HasErr(err error, errCode fmt.Stringer) bool {
	if err != nil {
		xlog.Logger.WithField("err", err).Error(errCode.String())
		return true
	}
	return false
}

func TransToByte(params any) []byte {
	bites, _ := json.Marshal(params)
	return bites
}

func HasWar(err error, errCode fmt.Stringer) bool {
	if err != nil {
		xlog.Logger.WithField("err", err).Warn(errCode.String())
		return true
	}
	return false
}

// HasQueryErr 数据库查询异常
func HasQueryErr(err error, errCode fmt.Stringer) bool {
	if err == nil {
		return false
	}
	if err == gorm.ErrRecordNotFound {
		msg := fmt.Sprintf("%s desc:%s", errCode.String(), global.NoRecordErr.String())
		xlog.WithField(global.ErrField, err).Warn(msg)
	} else {
		xlog.WithField(global.ErrField, err).Error(errCode.String())
	}
	return true
}

// OpenFreeUDPPort opens free UDP port
// This example does not actually use UDP ports,
// but to avoid port collisions with the HTTP server,
// it binds the same number of UDP port in advance.
func OpenFreeUDPPort(portBase int, num int) (net.PacketConn, int, error) {
	for i := 0; i < num; i++ {
		port := portBase + i
		conn, err := net.ListenPacket("udp", fmt.Sprint(":", port))
		if err != nil {
			continue
		}
		return conn, port, nil
	}
	return nil, 0, errors.New("failed to open free port")
}

func JsonDisplay(obj any) {
	b, _ := json.Marshal(obj)
	fmt.Println("---------------------------------json obj-------------------------------------")
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "\t")
	_, _ = out.WriteTo(os.Stdout)
	fmt.Printf("\n")
	fmt.Println("---------------------------------json obj-------------------------------------")
}

// IsStaticRequest 判断是否是静态文件请求
func IsStaticRequest(context *gin.Context) bool {
	if strings.Contains(context.Request.URL.Path, "/image/upload/") ||
		strings.Contains(context.Request.URL.Path, "/admin/") ||
		strings.Contains(context.Request.URL.Path, "/favicon.ico") ||
		strings.Contains(context.Request.URL.Path, ".js") {
		return true
	} else {
		return false
	}
}
func CheckErr(err error) bool {
	return err != nil
}

// Decimal 保留2位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// TransInterfaceToMap 转换struct 等结构到map
func TransInterfaceToMap(params any) map[string]any {
	var paramsMap map[string]any
	jsonData, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonData, &paramsMap)
	return paramsMap
}

// TransStruct 转换struct
func TransStruct[T any](source any, target T) (T, error) {
	jsonData, err := json.Marshal(source)
	if err != nil {
		return target, err
	}
	err = json.Unmarshal(jsonData, target)
	return target, err
}

func GetApiServer(url string) string {
	if strings.Contains(url, "@") {
		temps := strings.Split(url, "@")
		return temps[0]
	}
	return ""
}
func MaxNum(arr []int) (max int, maxIndex int) {
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	return max, maxIndex
}

func MinNum(arr []int) (min int, minIndex int) {
	min = arr[0]
	for index, val := range arr {
		if min > val {
			min = val
			minIndex = index
		}
	}
	return min, minIndex
}
