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
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils/cryptor"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/golang-module/carbon"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func Success(status int) bool {
	return global.ErrCode(status) == global.Success
}
func JsonString(params any) string {
	b, _ := json.Marshal(params)
	return string(b)
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

func ApiSign(url string, secret string) string {
	ts := fmt.Sprint(carbon.Now().Timestamp())
	signStr := fmt.Sprintf("%s@%s@%s", secret, ts, secret)
	// 对字符串进行sha1哈希
	sign := cryptor.Sha1(signStr)
	if strings.Contains(url, "?") {
		url += fmt.Sprintf("&ts=%s&sign=%s", ts, sign)
	} else {
		url += fmt.Sprintf("?ts=%s&sign=%s", ts, sign)
	}
	return url
}

// TransStrToImage base64 字符串转图片
func TransStrToImage(sourceString string, imageName string) error {
	dist, err := base64.StdEncoding.DecodeString(sourceString)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(imageName, dist, os.ModePerm)
	return err
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
