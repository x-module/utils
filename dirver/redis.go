/**
 * Created by Goland.
 * @file   redis.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2023/4/9 21:12
 * @desc   redis.go
 */

package dirver

import (
	"fmt"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils/xerror"
	"github.com/go-xmodule/utils/utils/xlog"
	"github.com/gomodule/redigo/redis"
	"time"
)

// InitializeRedis 初始化redis 链接
func InitializeRedis(host string, port int, password string, db int) *redis.Pool {
	address := fmt.Sprintf("%s:%d", host, port)
	xlog.Logger.Debug("start collect  redis.  address:", address)
	return &redis.Pool{ // 实例化一个连接池
		MaxIdle:     50,  // 最初的连接数量
		MaxActive:   0,   // 连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, // 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial("tcp", address,
				// redis.DialUseTLS(true),
				// redis.DialReadTimeout(20*time.Second),
				redis.DialWriteTimeout(20*time.Second),
				redis.DialDatabase(db),
				redis.DialPassword(password),
				// redis.DialTLSSkipVerify(true),
			)
			xerror.PanicErr(err, global.InitRedisErr.String())
			return client, err
		},
	}
}
