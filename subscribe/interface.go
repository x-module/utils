/**
 * Created by PhpStorm.
 * @file   interface.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/9 20:42
 * @desc   interface.go
 */

package subscribe

const (
	PublishErr = "发布消息异常"
)

// SubscribeCallback 订阅回调
type SubscribeCallback func(message string)

var SubscribeHandler SubPub

// SubPub 消息发布定义
type SubPub interface {
	// Subscribe 订阅消息
	Subscribe(channel string, callback SubscribeCallback)
	// Publish 发布消息
	Publish(channel string, message any) error
}
