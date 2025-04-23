package rdb

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/convert"
	"king.com/king/base/common/strs"
	"time"
)

type RdbBaseImpl struct {
	R *RedisClient
}

func (r *RdbBaseImpl) BuildSiteKey(virtId int64, bz string) *string {
	s := fmt.Sprintf("site:%v:bz:%v", virtId, bz)
	return &s
}
func (r *RdbBaseImpl) BuildUserKey(virtId int64, userId int64, bz string) *string {
	s := fmt.Sprintf("site:%v:userId:%v:bz:%v", virtId, userId, bz)
	return &s
}
func (r *RdbBaseImpl) BuildCache(k string, v interface{}) {
	m, err := json.Marshal(v)
	if err != nil {
		return
	}
	if !r.R.Set(k, string(m)) {
		logx.Errorf("setCache err k:%v,v:%v", k, v)
	}
}
func (r *RdbBaseImpl) ParseString(v1 string, v2 interface{}) {
	err := json.Unmarshal([]byte(v1), v2)
	if err != nil {
		logx.Errorf("ParseString err%v", err)
		return
	}
}

type RedisManger struct {
	RdbBaseImpl
	s     interface{}    //源数据
	tp    interface{}    //mode类型
	k     *string        //key
	v     *[]byte        // value
	t     *time.Duration //超时
	build bool           //构建完成后放可执行
}

func NewRM(R *RedisClient) *RedisManger {
	r := &RedisManger{}
	r.R = R
	return r
}

func (m *RedisManger) Source(s interface{}) *RedisManger {
	m.s = s
	return m
}
func (m *RedisManger) Mode(tp interface{}) *RedisManger {
	m.tp = tp
	return m
}
func (m *RedisManger) WithSiteKey(virtId int64, bz string) *RedisManger {
	m.k = m.BuildSiteKey(virtId, bz)
	return m
}
func (m *RedisManger) WithUserKey(virtId int64, userId int64, bz string) *RedisManger {
	m.k = m.BuildUserKey(virtId, userId, bz)
	return m
}
func (m *RedisManger) WithExp(t *time.Duration) *RedisManger {
	m.t = t
	return m
}

func (m *RedisManger) MustBuild() *RedisManger {
	if m.validMode() {
		return m
	}
	if err := copier.CopyWithOption(m.tp, m.s, convert.Time2DefaultFormatStr()); err != nil {
		logx.Errorf("MustBuild copy err:%v,tp:%+v", err, m.tp)
		return m
	}
	ms, err := json.Marshal(m.tp)
	if err != nil {
		logx.Errorf("MustBuild marshal err:%v,tp:%+v", err, m.tp)
		return m
	}
	m.v = &ms
	m.build = true
	return m
}

// 执行设置
func (m *RedisManger) SResult() bool {
	if !m.valid() {
		return false
	}
	if strs.IsEmpty(m.k) {
		return false
	}
	if m.t != nil {
		return m.setEx()
	}
	r := m.dfSet()
	m.build = false
	return r
}

// 执行查询
func (m *RedisManger) QResult() interface{} {
	if !m.valid() {
		return false
	}
	b, s := m.R.Get(*m.k)
	if !b {
		return nil
	}
	v := []byte(s)
	m.v = &v
	return m
}

// 执行删除
func (m *RedisManger) DResult() bool {
	if strs.IsEmpty(m.k) {
		return false
	}
	return m.R.Del(*m.k)
}
func (m *RedisManger) dfSet() bool {
	return m.R.Set(*m.k, string(*m.v))
}
func (m *RedisManger) setEx() bool {
	return m.R.SetEX(*m.k, string(*m.v), *m.t)
}
func (m *RedisManger) validMode() bool {
	if m.tp == nil {
		logx.Errorf("not invoke Mode fun")
		return false
	}
	return true
}
func (m *RedisManger) valid() bool {
	if m.tp == nil {
		logx.Errorf("not invoke Mode fun")
		return false
	}
	if !m.build {
		logx.Errorf("not invoke MustBuild fun")
		return false
	}
	return true
}
