package rdb

import (
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/strs"
	"time"
)

func (m *RedisManger) dfSet() bool {
	return m.R.Set(m.k, *m.v)
}
func (m *RedisManger) dfHSet() bool {
	k := *m.v
	return m.R.HSet(m.k, m.f, k)
}

func (m *RedisManger) setEx() bool {
	return m.R.SetEX(m.k, *m.v, *m.t)
}
func (m *RedisManger) hSetEx() bool {
	return m.R.SetEX(m.k, *m.v, *m.t)
}
func (m *RedisManger) hMSet(data *map[string]interface{}) bool {
	r := m.R.HMSet(m.k, *data)
	if m.t != nil {
		m.R.Expire(m.k, *m.t)
	}
	return r
}
func (m *RedisManger) validMode() bool {
	if m.tp == nil {
		logx.Errorf("not invoke Mode fun")
		return false
	}
	return true
}
func (m *RedisManger) validVal() bool {
	if m.v == nil {
		logx.Errorf("not invoke validVal fun")
		return false
	}
	return true
}
func (m *RedisManger) valid() bool {
	if !m.validKey() {
		return false
	}
	if !m.validVal() {
		return false
	}
	if !m.validBuild() {
		return false
	}
	return true
}
func (m *RedisManger) qValid() bool {
	if !m.validKey() {
		return false
	}

	return true
}
func (m *RedisManger) validBuild() bool {
	if !m.build {
		logx.Errorf("not invoke MustBuild fun")
	}
	return true
}
func (m *RedisManger) validField() bool {
	if strs.IsDefault(m.f) {
		logx.Errorf("not invoke WithField fun")
		return false
	}
	return true
}
func (m *RedisManger) validKey() bool {
	if strs.IsDefault(m.k) {
		logx.Errorf("not invoke WithKey fun")
		return false
	}
	return true
}
func (m *RedisManger) hValid() bool {
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}
	if !m.build {
		logx.Errorf("not invoke MustBuild fun")
		return false
	}
	return true
}
func (m *RedisManger) hmValid() bool {
	if !m.validMode() {
		return false
	}
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}
	if !m.build {
		logx.Errorf("not invoke MustBuild fun")
		return false
	}
	return true
}
func (m *RedisManger) scriptValid() bool {
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}

	return true
}
func (m *RedisManger) incrHScript() *string {
	script := `
		local key = KEYS[1]
		local field = ARGV[1]
		local increment = tonumber(ARGV[2])
		-- local ttl = tonumber(ARGV[3])
		
		-- 如果字段不存在，先初始化为0
		if redis.call("HEXISTS", key, field) == 0 then
			redis.call("HSET", key, field, 0)
		--	redis.call("EXPIRE", key, ttl)
		end
		
		-- 执行自增并返回新值
		return redis.call("HINCRBY", key, field, increment)
	`
	return &script
}
func (m *RedisManger) incrScript() *string {
	script := `
		local current = redis.call('INCRBY', KEYS[1], ARGV[2] or 1)
		if current == 1 then
			redis.call('EXPIRE', KEYS[1], ARGV[1])
		end
		return current
	`
	return &script
}

func (m *RedisManger) formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		return 1
	}
	return int64(dur / time.Second)
}
func (m *RedisManger) ProductPageScript() *string {
	script := `-- 获取总数量
		local k = KEYS[1]
		local k1= k .. '_set'
		local k2= k .. '_hash'
		local total = redis.call('ZCARD',k1)
		-- 获取分页的产品ID列表
		local productIDs = redis.call('ZREVRANGE', k1, ARGV[1], ARGV[2])
	
		-- 批量获取产品详情
		local products = redis.call('HMGET',k2, unpack(productIDs))
		
		return {total, products} `
	return &script
}
func (m *RedisManger) addProductScript() *string {
	script := `
	local k = KEYS[1]
	local id = KEYS[2]
	
	local k1 = k .. '_hash' 
	local k2 = k .. '_set' 
	local k3 = KEYS[3]
	
	local sortScore = ARGV[1]
	local productData = ARGV[2]
	local detailData = ARGV[3]
	
	-- 更新Hash
	redis.call('HSET', k1, id, productData)
	
	-- 更新ZSet索引
	redis.call('ZADD', k2, sortScore, id)

	-- 更新详情
	redis.call('SET', k3, detailData)
	
	return 1
	`
	return &script
}
func (m *RedisManger) delProductScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash'
	local k2 = k .. '_set'
	local zsetResult = redis.call('ZREM', k2, ARGV[1])
   	local hashResult = redis.call('HDEL', k1, ARGV[1])
   	local detaulResult = redis.call('DEL', KEYS[2])
    
 
    return {zsetResult, hashResult,detaulResult}`
	return &script
}
func (m *RedisManger) bannerListScript() *string {
	script := `
	local k = KEYS[1]
		local k1= k .. '_set'
		local k2= k .. '_hash'

		-- 获取分页的产品ID列表
		local bannerIDs = redis.call('ZREVRANGE', k1, 0, -1)
	
		-- 批量获取产品详情
		local products = redis.call('HMGET',k2, unpack(bannerIDs))
		
		return products`
	return &script
}
func (m *RedisManger) addBannerScript() *string {
	script := `
	local k = KEYS[1]
	local id = KEYS[2]
	local k1 = k .. '_hash' 
	local k2 = k .. '_set' 
	
	local sortScore = ARGV[1]
	local bannerData = ARGV[2]
	
	-- 更新Hash
	redis.call('HSET', k1, id, bannerData)
	
	-- 更新ZSet索引
	redis.call('ZADD', k2, sortScore, id)
	
	return 1`
	return &script
}
func (m *RedisManger) delBannerScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash'
	local k2 = k .. '_set'
	local zsetResult = redis.call('ZREM', k2, ARGV[1])
   	local hashResult = redis.call('HDEL', k1, ARGV[1])
    
 
    return 1`
	return &script
}

func (m *RedisManger) CartPageScript() *string {
	script := `-- 获取总数量
		local k = KEYS[1]
		local k1= k .. '_set'
		local k2= k .. '_hash'
		local total = redis.call('ZCARD',k1)
		-- 获取分页的cart combId 列表
		local productIDs = redis.call('ZREVRANGE', k1, ARGV[1], ARGV[2])
	
		-- 批量获取cart详情
		local carts = redis.call('HMGET',k2, unpack(productIDs))
		
		return {total, carts} `
	return &script
}
func (m *RedisManger) addCartScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash' 
	local k2 = k .. '_set'
	local count = redis.call('ZCARD', k2)
	if count >= 50 then
		return 2
	end

	local combId = KEYS[2]
	local expireTime = tonumber(ARGV[3])  -- 超时时间（秒）
	local sortScore = ARGV[1]
	local cartData = ARGV[2]
	
	-- 更新Hash
	redis.call('HSET', k1, combId, cartData)

	-- 为Hash设置超时时间
	redis.call('EXPIRE', k1, expireTime)
	
	-- 更新ZSet索引
	redis.call('ZADD', k2, sortScore, combId)

	-- 为Sorted Set设置超时时间
	redis.call('EXPIRE', k2, expireTime)
	
	return 1
	`
	return &script
}
func (m *RedisManger) delCartScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash'
	local k2 = k .. '_set'
	local zsetResult = redis.call('ZREM', k2, ARGV[1])
   	local hashResult = redis.call('HDEL', k1, ARGV[1])
 
    return {zsetResult, hashResult}`
	return &script
}
func (m *RedisManger) delUserCartScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash'
	local k2 = k .. '_set'
	local zsetResult = redis.call('DEL', k2)
   	local hashResult = redis.call('DEL', k1)
 
    return {zsetResult, hashResult}`
	return &script
}
func (m *RedisManger) hMSetExpScript() *string {
	script := `
    -- 删除整个 key
    redis.call('DEL', KEYS[1])
    
    -- 设置新数据
    for i = 1, #ARGV-1, 2 do
        redis.call('HSET', KEYS[1], ARGV[i], ARGV[i+1])
    end
    
    -- 设置过期时间（秒）
    if tonumber(ARGV[#ARGV]) > 0 then
        redis.call('EXPIRE', KEYS[1], ARGV[#ARGV])
    end
    
    return 1
`
	return &script
}

func (m *RedisManger) blogPageScript() *string {
	script := `-- 获取总数量
		local k = KEYS[1]
		local k1= k .. '_set'
		local k2= k .. '_hash'
		local total = redis.call('ZCARD',k1)
		-- 获取分页的blog id 列表
		local productIDs = redis.call('ZREVRANGE', k1, ARGV[1], ARGV[2])
	
		-- 批量获取blog内容
		local carts = redis.call('HMGET',k2, unpack(productIDs))
		
		return {total, carts} `
	return &script
}
func (m *RedisManger) addBlogScript() *string {
	script := `
	local k = KEYS[1]
	local id = KEYS[2]
	local k1 = k .. '_hash' 
	local k2 = k .. '_set' 
	local k3 = k .. '_detail_hash' 
	
	local sortScore = ARGV[1]
	local data = ARGV[2]
	local dData = ARGV[3]
	
	-- 更新Hash
	redis.call('HSET', k1, id, data)

	redis.call('HSET', k3, id, dData)
	
	-- 更新ZSet索引
	redis.call('ZADD', k2, sortScore, id)
	
	return 1`
	return &script
}
func (m *RedisManger) delBlogScript() *string {
	script := `
	local k = KEYS[1]
	local k1 = k .. '_hash'
	local k2 = k .. '_set'
	local k3 = k .. '_detail_hash'
	local zsetResult = redis.call('ZREM', k2, ARGV[1])
	local detailResult = redis.call('ZREM', k3, ARGV[1])
   	local hashResult = redis.call('HDEL', k1, ARGV[1])
    
 
    return 1`
	return &script
}
