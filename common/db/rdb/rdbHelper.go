package rdb

import (
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/strs"
	"time"
)

func (m *RedisManger) dfSet() bool {
	return m.R.Set(*m.k, *m.v)
}
func (m *RedisManger) dfHSet() bool {
	k := *m.v
	return m.R.HSet(*m.k, *m.f, k)
}

func (m *RedisManger) setEx() bool {
	return m.R.SetEX(*m.k, *m.v, *m.t)
}
func (m *RedisManger) hSetEx() bool {
	return m.R.SetEX(*m.k, *m.v, *m.t)
}
func (m *RedisManger) validMode() bool {
	if m.tp == nil {
		logx.Errorf("not invoke Mode fun")
		return false
	}
	return true
}
func (m *RedisManger) valid() bool {
	if !m.validMode() {
		return false
	}
	if !m.validKey() {
		return false
	}
	if !m.validBuild() {
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
	if strs.IsEmpty(m.f) {
		logx.Errorf("not invoke WithField fun")
		return false
	}
	return true
}
func (m *RedisManger) validKey() bool {
	if strs.IsEmpty(m.k) {
		logx.Errorf("not invoke WithKey fun")
		return false
	}
	return true
}
func (m *RedisManger) hValid() bool {
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
func (m *RedisManger) incrScript() *string {
	script := `
		local key = KEYS[1]
		local field = ARGV[1]
		local increment = tonumber(ARGV[2])
		local ttl = tonumber(ARGV[3])
		
		-- 如果字段不存在，先初始化为0
		if redis.call("HEXISTS", key, field) == 0 then
			redis.call("HSET", key, field, 0)
			redis.call("EXPIRE", key, ttl)
		end
		
		-- 执行自增并返回新值
		return redis.call("HINCRBY", key, field, increment)
	`
	return &script
}

func (m *RedisManger) formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		return 1
	}
	return int64(dur / time.Second)
}
