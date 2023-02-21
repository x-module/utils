/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/2 15:31
 * @desc   utils.go
 */

package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils/xlog"
	"net"
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
