/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/3 19:37
 * @desc   utils.go
 */

package request

import (
	"fmt"
	"github.com/go-utils-module/utils/utils/cryptor"
)

func RequestSign(ts string, secret string) string {
	signStr := fmt.Sprintf("%s@%s@%s", secret, ts, secret)
	return cryptor.Sha1(signStr)
}
