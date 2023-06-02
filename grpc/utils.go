/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/2 15:16
 * @desc   utils.go
 */

package grpc

import (
	"fmt"
	"github.com/x-module/utils/utils/cryptor"
	"time"
)

func RpcSign(secret string) (string, int64) {
	ts := time.Now().Unix()
	signStr := fmt.Sprintf("%s@%d@%s", secret, ts, secret)
	return cryptor.Sha1(signStr), ts
}
func RequestSign(ts string, secret string) string {
	signStr := fmt.Sprintf("%s@%s@%s", secret, ts, secret)
	return cryptor.Sha1(signStr)
}
