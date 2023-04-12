/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:41
 * @desc   redis.go
 */

package subscribe

import (
	"github.com/go-xmodule/module/global"
	global2 "github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/handler"
	"github.com/go-xmodule/utils/utils"
	"github.com/go-xmodule/utils/utils/xlog"
)

type RedisSubscribe struct {
}

func NewRedisSubscribe() {
	SubscribeHandler = new(RedisSubscribe)
}

// Subscribe 订阅消息
func (s *RedisSubscribe) Subscribe(channel string, callback handler.SubscribeCallback) {
	xlog.Logger.Debug("start subscribe data, channel:", channel)
	err := handler.RedisHandler.Subscribe(channel, func(message string) {
		// 处理消息
		xlog.Logger.Debug("consumer data:", utils.JsonString(message))
		callback(message)
	})
	if err != nil {
		xlog.Logger.WithField(global.ErrField, err).Error(global2.SubscribeDataErr.String())
	}
}

// Publish 发布数据
func (s *RedisSubscribe) Publish(channel string, message any) error {
	xlog.Logger.Debug("start publish data,channel:", channel, " message:", utils.JsonString(message))
	_, err := handler.RedisHandler.Publish(channel, message)
	if err != nil {
		xlog.Logger.Error(PublishErr, err.Error())
		return err
	}
	xlog.Logger.Debug("publish data success")
	return nil
}
