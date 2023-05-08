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

// SubscribeCallback 订阅回调
type SubscribeCallback func(message string)

type RedisClient struct {
	pool    *redis.Pool
	con     redis.Conn
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

func (r *RedisClient) GetCon() redis.Conn {
	r.con = r.pool.Get()
	return r.con
}

// Set 字符串设置
func (r *RedisClient) Set(key string, value any, expiration ...time.Duration) error {
	client := r.pool.Get()
	defer r.Close(client)

	// 执行命令
	_, err := client.Do("Set", key, value)
	if err != nil {
		return err
	}
	if len(expiration) > 0 {
		_, err = client.Do("expire", key, int64(expiration[0]))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisClient) Close(conn redis.Conn) {
	_ = conn.Close()
	r.con = nil
}

func (r *RedisClient) HSetNX(key, field string, value any) (bool, error) {
	client := r.pool.Get()
	defer r.Close(client)
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
func (r *RedisClient) HSet(key string, values map[string]string, con ...redis.Conn) (int64, error) {
	var client redis.Conn
	if len(con) > 0 {
		client = con[0]
	} else {
		client = r.pool.Get()
		defer r.Close(client)
	}
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
	defer r.Close(client)
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
	defer r.Close(client)
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
	defer r.Close(client)
	result, err := client.Do("HGet", key, field)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", nil
	}
	return string(result.([]uint8)), nil
}

func (r *RedisClient) HDel(key, field string, con ...redis.Conn) (int64, error) {
	var client redis.Conn
	if len(con) > 0 {
		client = con[0]
	} else {
		client = r.pool.Get()
		defer r.Close(client)
	}
	result, err := client.Do("HDEL", key, field)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, nil
	}
	return result.(int64), nil
}

func (r *RedisClient) Get(key string) (bool, string, error) {
	client := r.pool.Get()
	defer r.Close(client)
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
	var client redis.Conn
	if r.con != nil {
		client = r.con
	} else {
		client = r.pool.Get()
		defer r.Close(client)
	}
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
	defer r.Close(client)
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
	defer r.Close(client)
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
	defer r.Close(client)
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
	defer r.Close(client)
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
	defer r.Close(client)
	result, err := client.Do("Publish", channel, message)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, nil
	}
	return result.(int64), nil
}

// Subscribe 订阅
func (r *RedisClient) Subscribe(channel string, callback SubscribeCallback) error {
	client := r.pool.Get()
	defer r.Close(client)
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
