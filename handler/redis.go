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
	"github.com/go-redis/redis/v8"
	"github.com/go-utils-module/utils/dirver"
	"github.com/go-utils-module/utils/global"
	"github.com/go-utils-module/utils/utils/xerror"
	"reflect"
	"time"
)

// RedisHandler Redis操作句柄
var RedisHandler *RedisClient

type RedisClient struct {
	client       *redis.Client
	backupClient *redis.Client
	context      context.Context
}
type SubscribeData struct {
	Channel      string `json:"Channel"`
	Pattern      string `json:"Pattern"`
	Payload      string `json:"Payload"`
	PayloadSlice any    `json:"PayloadSlice"`
}

// RedisConfig 配置
type RedisConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Db         int    `yaml:"db"`
	Password   string `yaml:"password"`
	MaxRetries int    `yaml:"maxRetries"`
}

// InitializeRedisPool 初始化redis连接池
func InitializeRedisPool(config RedisConfig) *redis.Client {
	c, err := dirver.InitializeRedis(config.Host, config.Port, config.Password, config.Db)
	xerror.PanicErr(err, global.InitRedisErr.String())
	RedisHandler = NewRedis(c)
	return c
}

// NewRedis 实例化
func NewRedis(client *redis.Client) *RedisClient {
	redisClient := new(RedisClient)
	redisClient.context = context.Background()
	redisClient.client = client
	return redisClient
}

// Set 字符串设置
func (r *RedisClient) Set(key string, value any, expiration time.Duration) error {
	// 执行命令
	err := r.client.Set(r.context, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) HSetNX(key, field string, value any) (bool, error) {
	// 执行命令
	result := r.client.HSetNX(r.context, key, field, value)
	if err := result.Err(); err != nil {
		return false, err
	}
	return result.Result()
}

// HSet hash-set
func (r *RedisClient) HSet(key string, values map[string]string) (int64, error) {
	// 执行命令
	result := r.client.HSet(r.context, key, values)
	if err := result.Err(); err != nil {
		return 0, err
	}
	return result.Result()
}

func (r *RedisClient) SetNX(key string, value any, expiration time.Duration) (bool, error) {
	// 执行命令
	// r.client.Conn(r.context).Select(r.context,1)
	result := r.client.SetNX(r.context, key, value, expiration)
	if err := result.Err(); err != nil {
		return false, err
	}
	return result.Result()
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	// r.checkConnect()
	// 执行命令
	result := r.client.HGetAll(r.context, key)
	if err := result.Err(); err != nil {
		return nil, err
	}
	return result.Result()
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	// r.checkConnect()
	// 执行命令
	result := r.client.HGet(r.context, key, field)
	if err := result.Err(); err != nil {
		return "", err
	}
	return result.Result()
}

func (r *RedisClient) Get(key string) (bool, string, error) {
	// r.checkConnect()
	result, err := r.client.Get(r.context, key).Result()
	if err == redis.Nil {
		return false, "", nil
	} else if err != nil {
		return false, "", err
	} else {
		return true, result, nil
	}
}

func (r *RedisClient) Delete(key ...string) error {
	// r.checkConnect()
	err := r.client.Del(r.context, key...).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r *RedisClient) LPop(key string) (string, error) {
	result := r.client.LPop(r.context, key)
	if err := result.Err(); err != nil {
		return "", err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) LPush(key string, value any) (int64, error) {
	result := r.client.LPush(r.context, key, value)
	if err := result.Err(); err != nil {
		return 0, err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) BLPop(key string, timeout time.Duration) ([]string, error) {
	result := r.client.BLPop(r.context, timeout, key)
	if err := result.Err(); err != nil {
		return []string{}, err
	} else {
		return result.Result()
	}
}

func (r *RedisClient) Keys(pattern string) ([]string, error) {
	result := r.client.Keys(r.context, pattern)
	if err := result.Err(); err != nil {
		return nil, err
	} else {
		return result.Result()
	}
}

// Publish 发布
func (r *RedisClient) Publish(channel string, message any) (int64, error) {
	result := r.client.Publish(r.context, channel, message)
	if err := result.Err(); err != nil {
		return 0, err
	} else {
		return result.Result()
	}
}

// Subscribe 订阅
func (r *RedisClient) Subscribe(channel string) <-chan *redis.Message {
	result := r.client.Subscribe(r.context, channel)
	return result.Channel()
}

type PipAction func(pipeLiner redis.Pipeliner)

func (r *RedisClient) SelectDbAction(index int, action PipAction) (map[int]any, error) {
	pipeLine := r.client.Pipeline()
	pipeLine.Do(context.Background(), "select", index)
	action(pipeLine)
	result, err := pipeLine.Exec(context.Background())
	if err != nil {
		return map[int]any{}, err
	}
	return r.getCmdResult(result), nil
}

func (r *RedisClient) getCmdResult(cmdRes []redis.Cmder) map[int]any {
	strMap := make(map[int]any, len(cmdRes))
	for idx, cmder := range cmdRes {
		// *ClusterSlotsCmd 未实现
		switch reflect.TypeOf(cmder).String() {
		case "*redis.Cmd":
			cmd := cmder.(*redis.Cmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringCmd":
			cmd := cmder.(*redis.StringCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.SliceCmd":
			cmd := cmder.(*redis.SliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringSliceCmd":
			cmd := cmder.(*redis.StringSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStringMapCmd":
			cmd := cmder.(*redis.StringStringMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringIntMapCmd":
			cmd := cmder.(*redis.StringIntMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolCmd":
			cmd := cmder.(*redis.BoolCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.BoolSliceCmd":
			cmd := cmder.(*redis.BoolSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.IntCmd":
			cmd := cmder.(*redis.IntCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.FloatCmd":
			cmd := cmder.(*redis.FloatCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StatusCmd":
			cmd := cmder.(*redis.StatusCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.TimeCmd":
			cmd := cmder.(*redis.TimeCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.DurationCmd":
			cmd := cmder.(*redis.DurationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.StringStructMapCmd":
			cmd := cmder.(*redis.StringStructMapCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XMessageSliceCmd":
			cmd := cmder.(*redis.XMessageSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XStreamSliceCmd":
			cmd := cmder.(*redis.XStreamSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingCmd":
			cmd := cmder.(*redis.XPendingCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.XPendingExtCmd":
			cmd := cmder.(*redis.XPendingExtCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZSliceCmd":
			cmd := cmder.(*redis.ZSliceCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.ZWithKeyCmd":
			cmd := cmder.(*redis.ZWithKeyCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.CommandsInfoCmd":
			cmd := cmder.(*redis.CommandsInfoCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoLocationCmd":
			cmd := cmder.(*redis.GeoLocationCmd)
			strMap[idx], _ = cmd.Result()
			break
		case "*redis.GeoPosCmd":
			cmd := cmder.(*redis.GeoPosCmd)
			strMap[idx], _ = cmd.Result()
			break
		}
	}
	return strMap
}
