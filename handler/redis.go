/**
 * Created by goland.
 * @file   redis.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/16 15:30
 * @desc   redis.go
 */

package handler

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/x-module/utils/dirver"
	"time"
)

// RedisHandler Redis操作句柄
var RedisHandler *RedisClient

// RedisClient Redis客户端
type RedisClient struct {
	*redis.Ring
	ctx context.Context
}

// NewRedis 实例化
func NewRedis(config dirver.RedisConfig) *RedisClient {
	RedisHandler = &RedisClient{
		Ring: dirver.InitializeRedis(config),
		ctx:  context.Background(),
	}
	return RedisHandler
}

// SetContext 设置context
func (r *RedisClient) SetContext(ctx context.Context) {
	r.ctx = ctx
}

// Set 字符串设置
func (r *RedisClient) Set(key string, value any, expiration ...time.Duration) *redis.StatusCmd {
	exp := time.Duration(0)
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	return r.Ring.Set(r.ctx, key, value, exp)
}
func (r *RedisClient) Get(key string) *redis.StringCmd {
	return r.Ring.Get(r.ctx, key)
}
func (r *RedisClient) Delete(key string) *redis.IntCmd {
	return r.Ring.Del(r.ctx, key)
}

// SetNx 字符串设置
func (r *RedisClient) SetNx(key string, value any, expiration ...time.Duration) *redis.BoolCmd {
	exp := time.Duration(0)
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	return r.Ring.SetNX(r.ctx, key, value, exp)
}

// HSet hash设置
func (r *RedisClient) HSet(key string, params map[string]any) *redis.IntCmd {
	var paramsList []any
	for k, v := range params {
		paramsList = append(paramsList, k)
		paramsList = append(paramsList, v)
	}
	return r.Ring.HSet(r.ctx, key, paramsList...)
}

// HSetNX hash设置
func (r *RedisClient) HSetNX(key string, field string, value any) *redis.BoolCmd {
	return r.Ring.HSetNX(r.ctx, key, field, value)
}

// HGetAll 获取全部hash
func (r *RedisClient) HGetAll(key string) *redis.MapStringStringCmd {
	return r.Ring.HGetAll(r.ctx, key)
}

func (r *RedisClient) HGet(key string, field string) *redis.StringCmd {
	return r.Ring.HGet(r.ctx, key, field)
}

func (r *RedisClient) HDel(key string, field string) *redis.IntCmd {
	return r.Ring.HDel(r.ctx, key, field)
}

func (r *RedisClient) LPush(key string, values ...any) *redis.IntCmd {
	return r.Ring.LPush(r.ctx, key, values...)
}

func (r *RedisClient) LPop(key string) *redis.StringCmd {
	return r.Ring.LPop(r.ctx, key)
}

func (r *RedisClient) Keys(pattern string) *redis.StringSliceCmd {
	return r.Ring.Keys(r.ctx, pattern)
}

// Publish 发布
func (r *RedisClient) Publish(channel string, message any) error {
	return r.Ring.Publish(r.ctx, channel, message).Err()
}

// SubscribeCallback 订阅回调
type SubscribeCallback func(message string)

// Subscribe 订阅
func (r *RedisClient) Subscribe(channel string, callback SubscribeCallback) error {
	pubSub := r.Ring.Subscribe(r.ctx, channel)
	defer pubSub.Close()
	for {
		msg, err := pubSub.ReceiveMessage(r.ctx)
		if err != nil {
			panic(err)
		}
		callback(msg.Payload)
	}
}
