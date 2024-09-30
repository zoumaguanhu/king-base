package rdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

type RdsCnf struct {
	Addr     string
	DB       int
	Password string
	PoolSize int
}
type RedisClient struct {
	opt    *redis.Options
	client *redis.Client
	cnf    *RdsCnf
	ctx    context.Context
}

func New(cnf *RdsCnf) *RedisClient {
	opt := &redis.Options{
		Addr:     cnf.Addr,
		DB:       cnf.DB,
		Password: cnf.Password,
		PoolSize: cnf.PoolSize,
	}
	rds := &RedisClient{cnf: cnf, opt: opt, ctx: context.Background(), client: redis.NewClient(opt)}
	_, err := rds.client.Ping(rds.ctx).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return nil
	}
	logc.Infof(rds.ctx, "redis comnect success")
	return rds
}

/*------------------------------------ 字符 操作 ------------------------------------*/

// Set 设置 key的值
func (rds *RedisClient) Set(key, value string) bool {
	result, err := rds.client.Set(rds.ctx, key, value, 0).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return result == "OK"
}

// SetEX 设置 key的值并指定过期时间
func (rds *RedisClient) SetEX(key, value string, ex time.Duration) bool {
	result, err := rds.client.Set(rds.ctx, key, value, ex).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return result == "OK"
}

// Get 获取 key的值
func (rds *RedisClient) Get(key string) (bool, string) {
	result, err := rds.client.Get(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false, ""
	}
	return true, result
}

// GetSet 设置新值获取旧值
func (rds *RedisClient) GetSet(key, value string) (bool, string) {
	oldValue, err := rds.client.GetSet(rds.ctx, key, value).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false, ""
	}
	return true, oldValue
}

// Incr key值每次加一 并返回新值
func (rds *RedisClient) Incr(key string) int64 {
	val, err := rds.client.Incr(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// IncrBy key值每次加指定数值 并返回新值
func (rds *RedisClient) IncrBy(key string, incr int64) int64 {
	val, err := rds.client.IncrBy(rds.ctx, key, incr).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func (rds *RedisClient) IncrByFloat(key string, incrFloat float64) float64 {
	val, err := rds.client.IncrByFloat(rds.ctx, key, incrFloat).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// Decr key值每次递减 1 并返回新值
func (rds *RedisClient) Decr(key string) int64 {
	val, err := rds.client.Decr(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// DecrBy key值每次递减指定数值 并返回新值
func (rds *RedisClient) DecrBy(key string, incr int64) int64 {
	val, err := rds.client.DecrBy(rds.ctx, key, incr).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// Del 删除 key
func (rds *RedisClient) Del(key string) bool {
	result, err := rds.client.Del(rds.ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// Expire 设置 key的过期时间
func (rds *RedisClient) Expire(key string, ex time.Duration) bool {
	result, err := rds.client.Expire(rds.ctx, key, ex).Result()
	if err != nil {
		return false
	}
	return result
}

/*------------------------------------ list 操作 ------------------------------------*/

// LPush 从列表左边插入数据，并返回列表长度
func (rds *RedisClient) LPush(key string, date ...interface{}) int64 {
	result, err := rds.client.LPush(rds.ctx, key, date).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return result
}

// RPush 从列表右边插入数据，并返回列表长度
func (rds *RedisClient) RPush(key string, date ...interface{}) int64 {
	result, err := rds.client.RPush(rds.ctx, key, date).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return result
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (rds *RedisClient) LPop(key string) (bool, string) {
	val, err := rds.client.LPop(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false, ""
	}
	return true, val
}

// RPop 从列表右边删除第一个数据，并返回删除的数据
func (rds *RedisClient) RPop(key string) (bool, string) {
	val, err := rds.client.RPop(rds.ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, val
}

// LIndex 根据索引坐标，查询列表中的数据
func (rds *RedisClient) LIndex(key string, index int64) (bool, string) {
	val, err := rds.client.LIndex(rds.ctx, key, index).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false, ""
	}
	return true, val
}

// LLen 返回列表长度
func (rds *RedisClient) LLen(key string) int64 {
	val, err := rds.client.LLen(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func (rds *RedisClient) LRange(key string, start, stop int64) []string {
	vales, err := rds.client.LRange(rds.ctx, key, start, stop).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return vales
}

// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func (rds *RedisClient) LRem(key string, count int64, data interface{}) bool {
	_, err := rds.client.LRem(rds.ctx, key, count, data).Result()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// LInsert 在列表中 pivot 元素的后面插入 data
func (rds *RedisClient) LInsert(key string, pivot int64, data interface{}) bool {
	err := rds.client.LInsert(rds.ctx, key, "after", pivot, data).Err()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return true
}

/*------------------------------------ set 操作 ------------------------------------*/

// SAdd 添加元素到集合中
func (rds *RedisClient) SAdd(key string, data ...interface{}) bool {
	err := rds.client.SAdd(rds.ctx, key, data).Err()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return true
}

// SCard 获取集合元素个数
func (rds *RedisClient) SCard(key string) int64 {
	size, err := rds.client.SCard(rds.ctx, "key").Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return size
}

// SIsMember 判断元素是否在集合中
func (rds *RedisClient) SIsMember(key string, data interface{}) bool {
	ok, err := rds.client.SIsMember(rds.ctx, key, data).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return ok
}

// SMembers 获取集合所有元素
func (rds *RedisClient) SMembers(key string) []string {
	es, err := rds.client.SMembers(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return es
}

// SRem 删除 key集合中的 data元素
func (rds *RedisClient) SRem(key string, data ...interface{}) bool {
	_, err := rds.client.SRem(rds.ctx, key, data).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return true
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func (rds *RedisClient) SPopN(key string, count int64) []string {
	vales, err := rds.client.SPopN(rds.ctx, key, count).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return vales
}

/*------------------------------------ hash 操作 ------------------------------------*/

// HSet 根据 key和 field字段设置，field字段的值
func (rds *RedisClient) HSet(key, field, value string) bool {
	err := rds.client.HSet(rds.ctx, key, field, value).Err()
	if err != nil {
		return false
	}
	return true
}

// HGet 根据 key和 field字段，查询field字段的值
func (rds *RedisClient) HGet(key, field string) string {
	val, err := rds.client.HGet(rds.ctx, key, field).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return val
}

// HMGet 根据key和多个字段名，批量查询多个 hash字段值
func (rds *RedisClient) HMGet(key string, fields ...string) []interface{} {
	vales, err := rds.client.HMGet(rds.ctx, key, fields...).Result()
	if err != nil {
		panic(err)
	}
	return vales
}

// HGetAll 根据 key查询所有字段和值
func (rds *RedisClient) HGetAll(key string) map[string]string {
	data, err := rds.client.HGetAll(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return data
}

// HKeys 根据 key返回所有字段名
func (rds *RedisClient) HKeys(key string) []string {
	fields, err := rds.client.HKeys(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return fields
}

// HLen 根据 key，查询hash的字段数量
func (rds *RedisClient) HLen(key string) int64 {
	size, err := rds.client.HLen(rds.ctx, key).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
	}
	return size
}

// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func (rds *RedisClient) HMSet(key string, data map[string]interface{}) bool {
	result, err := rds.client.HMSet(rds.ctx, key, data).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return result
}

// HSetNX 如果 field字段不存在，则设置 hash字段值
func (rds *RedisClient) HSetNX(key, field string, value interface{}) bool {
	result, err := rds.client.HSetNX(rds.ctx, key, field, value).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return result
}

// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func (rds *RedisClient) HDel(key string, fields ...string) bool {
	_, err := rds.client.HDel(rds.ctx, key, fields...).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return true
}

// HExists 检测 hash字段名是否存在
func (rds *RedisClient) HExists(key, field string) bool {
	result, err := rds.client.HExists(rds.ctx, key, field).Result()
	if err != nil {
		logc.Errorf(rds.ctx, "redis err:%v", err)
		return false
	}
	return result
}
