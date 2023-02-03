/**
 * Created by PhpStorm.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:41
 * @desc   redis.go
 */

package subscribe

import (
	"github.com/go-utils-module/utils/handler"
	"github.com/go-utils-module/utils/utils"
	"github.com/go-utils-module/utils/utils/convertor"
	"github.com/go-utils-module/utils/utils/xlog"
)

type RedisSubscribe struct {
}

func NewRedisSubscribe() {
	SubscribeHandler = new(RedisSubscribe)
}

// Subscribe 订阅消息
func (s *RedisSubscribe) Subscribe(channel string, callback SubscribeCallback) {
	xlog.Logger.Debug("start subscribe data, channel:", channel)
	messageList := handler.RedisHandler.Subscribe(channel)
	for message := range messageList {
		// 处理消息
		xlog.Logger.Debug("consumer data:", utils.JsonString(message))
		var data handler.SubscribeData
		_ = convertor.TransInterfaceToStruct(message, &data)
		callback(data.Payload)
	}
}

// Publish 发布数据
func (s *RedisSubscribe) Publish(channel string, message any) error {
	xlog.Logger.Debug("start publish data, message:", utils.JsonString(message))
	_, err := handler.RedisHandler.Publish(channel, message)
	if err != nil {
		xlog.Logger.Error(PublishErr, err.Error())
		return err
	}
	xlog.Logger.Debug("publish data success")
	return nil
}
