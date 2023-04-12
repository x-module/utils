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
	"fmt"
	"github.com/go-xmodule/utils/dirver"
	"github.com/gomodule/redigo/redis"
	"time"
)

// RedisHandler Redis操作句柄
var RedisHandler *RedisClient

type RedisClient struct {
	pool    *redis.Pool
	context context.Context
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
func InitializeRedisPool(config RedisConfig) *RedisClient {
	pool := dirver.InitializeRedis(config.Host, config.Port, config.Password, config.Db)
	RedisHandler = NewRedis(pool)
	return RedisHandler
}

// NewRedis 实例化
func NewRedis(pool *redis.Pool) *RedisClient {
	redisClient := new(RedisClient)
	redisClient.context = context.Background()
	redisClient.pool = pool
	return redisClient
}
func (r *RedisClient) GetPool() *redis.Pool {
	return r.pool
}

// Set 字符串设置
func (r *RedisClient) Set(key string, value any, expiration time.Duration) error {
	client := r.pool.Get()
	defer r.closet(client)

	// 执行命令
	_, err := client.Do("Set", key, value)
	if err != nil {
		return err
	}
	_, err = client.Do("expire", key, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) closet(conn redis.Conn) {
	_ = conn.Close()
}

func (r *RedisClient) HSetNX(key, field string, value any) (bool, error) {
	client := r.pool.Get()
	defer r.closet(client)
	// 执行命令
	result, err := client.Do("HSetNX", key, field, value)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return result.(int64) == 1, nil
}

// HSet hash-set
func (r *RedisClient) HSet(key string, values map[string]string) (int64, error) {
	client := r.pool.Get()
	defer r.closet(client)
	paramsList := []interface{}{key}
	for k, v := range values {
		paramsList = append(paramsList, k, v)
	}
	result, err := client.Do("HSet", paramsList...)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, nil
	}
	return result.(int64), nil
}

func (r *RedisClient) SetNX(key string, value any, expiration time.Duration) (bool, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("SetNX", key, value)
	if err != nil {
		return false, err
	}
	_, err = client.Do("expire", key, expiration)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return result.(int64) == 1, nil
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	res := map[string]string{}
	result, err := client.Do("HGetAll", key)
	if err != nil {
		return res, err
	}
	if result == nil {
		return res, nil
	}
	tKey := ""
	for _, v := range result.([]interface{}) {
		temp := string(v.([]uint8))
		if tKey == "" {
			tKey = temp
		} else {
			res[tKey] = temp
			tKey = ""
		}
	}
	return res, nil
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("HGet", key, field)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", nil
	}
	return string(result.([]uint8)), nil
}

func (r *RedisClient) Get(key string) (bool, string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("Get", key)
	if err != nil {
		return false, "", err
	}
	if result == nil {
		return false, "", nil
	}
	if result == nil {
		return false, "", nil
	}
	return true, string(result.([]uint8)), nil
}

func (r *RedisClient) Delete(keys ...string) error {
	client := r.pool.Get()
	defer r.closet(client)
	for _, k := range keys {
		_, err := client.Do("Del", k)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisClient) LPop(key string) (string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("LPOP", key)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", nil
	}
	return string(result.([]uint8)), nil
}

func (r *RedisClient) LPush(key string, value any) (int64, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("LPUSH", key, value)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, nil
	}
	return result.(int64), nil
}

func (r *RedisClient) BLPop(key string, timeout time.Duration) ([]string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("BLPop", key, timeout)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, v := range result.([]interface{}) {
		res = append(res, string(v.([]uint8)))
	}
	return res, nil
}

func (r *RedisClient) Keys(pattern string) ([]string, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("Keys", pattern)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, v := range result.([]interface{}) {
		res = append(res, string(v.([]uint8)))
	}
	return res, nil
}

// Publish 发布
func (r *RedisClient) Publish(channel string, message any) (int64, error) {
	client := r.pool.Get()
	defer r.closet(client)
	result, err := client.Do("Publish", channel, message)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, nil
	}
	return result.(int64), nil
}

// SubscribeCallback 订阅回调
type SubscribeCallback func(message string)

// Subscribe 订阅
func (r *RedisClient) Subscribe(channel string, callback SubscribeCallback) error {
	client := r.pool.Get()
	defer r.closet(client)
	psc := redis.PubSubConn{Conn: client}
	err := psc.Subscribe(channel)
	if err != nil {
		return err
	}
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			callback(string(v.Data))
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			panic(v)
		}
	}
}

// type PipAction func(pipeLiner redis.Pipeliner)
// func (r *RedisClient) SelectDbAction(index int, action PipAction) (map[int]any, error) {
// 	pipeLine := r.client.Pipeline()
// 	pipeLine.Do(context.Background(), "select", index)
// 	action(pipeLine)
// 	result, err := pipeLine.Exec(context.Background())
// 	if err != nil {
// 		return map[int]any{}, err
// 	}
// 	return r.getCmdResult(result), nil
// }
//
// func (r *RedisClient) getCmdResult(cmdRes []redis.Cmder) map[int]any {
// 	strMap := make(map[int]any, len(cmdRes))
// 	for idx, cmder := range cmdRes {
// 		// *ClusterSlotsCmd 未实现
// 		switch reflect.TypeOf(cmder).String() {
// 		case "*redis.Cmd":
// 			cmd := cmder.(*redis.Cmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StringCmd":
// 			cmd := cmder.(*redis.StringCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.SliceCmd":
// 			cmd := cmder.(*redis.SliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StringSliceCmd":
// 			cmd := cmder.(*redis.StringSliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StringStringMapCmd":
// 			cmd := cmder.(*redis.StringStringMapCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StringIntMapCmd":
// 			cmd := cmder.(*redis.StringIntMapCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.BoolCmd":
// 			cmd := cmder.(*redis.BoolCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.BoolSliceCmd":
// 			cmd := cmder.(*redis.BoolSliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.IntCmd":
// 			cmd := cmder.(*redis.IntCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.FloatCmd":
// 			cmd := cmder.(*redis.FloatCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StatusCmd":
// 			cmd := cmder.(*redis.StatusCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.TimeCmd":
// 			cmd := cmder.(*redis.TimeCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.DurationCmd":
// 			cmd := cmder.(*redis.DurationCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.StringStructMapCmd":
// 			cmd := cmder.(*redis.StringStructMapCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.XMessageSliceCmd":
// 			cmd := cmder.(*redis.XMessageSliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.XStreamSliceCmd":
// 			cmd := cmder.(*redis.XStreamSliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.XPendingCmd":
// 			cmd := cmder.(*redis.XPendingCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.XPendingExtCmd":
// 			cmd := cmder.(*redis.XPendingExtCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.ZSliceCmd":
// 			cmd := cmder.(*redis.ZSliceCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.ZWithKeyCmd":
// 			cmd := cmder.(*redis.ZWithKeyCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.CommandsInfoCmd":
// 			cmd := cmder.(*redis.CommandsInfoCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.GeoLocationCmd":
// 			cmd := cmder.(*redis.GeoLocationCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		case "*redis.GeoPosCmd":
// 			cmd := cmder.(*redis.GeoPosCmd)
// 			strMap[idx], _ = cmd.Result()
// 			break
// 		}
// 	}
// 	return strMap
// }
