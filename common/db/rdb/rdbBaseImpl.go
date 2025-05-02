package rdb

import (
	"context"
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

func (r *RdbBaseImpl) BuildKey(key, bz string) *string {
	s := fmt.Sprintf("site:%v:%v", key, bz)
	return &s
}
func (r *RdbBaseImpl) BuildHostKey(host, bz string) *string {
	s := fmt.Sprintf("site:vhost:%v:%v", host, bz)
	return &s
}
func (r *RdbBaseImpl) BuildSiteKey(virtId int64, bz string) *string {
	s := fmt.Sprintf("site:vsite:%v:%v", virtId, bz)
	return &s
}
func (r *RdbBaseImpl) BuildUserKey(virtId int64, userId int64, bz string) *string {
	s := fmt.Sprintf("site:virtId:%v:userId:%v:%v", virtId, userId, bz)
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
	f     *string        //field
	v     *string        // value
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
func (m *RedisManger) WithKey(key, bz string) *RedisManger {
	m.k = m.BuildKey(key, bz)
	return m
}
func (m *RedisManger) WithSiteKey(virtId int64, bz string) *RedisManger {
	m.k = m.BuildSiteKey(virtId, bz)
	return m
}
func (m *RedisManger) WithHostKey(host, bz string) *RedisManger {
	m.k = m.BuildHostKey(host, bz)
	return m
}
func (m *RedisManger) WithUserKey(virtId int64, userId int64, bz string) *RedisManger {
	m.k = m.BuildUserKey(virtId, userId, bz)
	return m
}
func (m *RedisManger) WithField(f *string) *RedisManger {
	m.f = f
	return m
}
func (m *RedisManger) WithExp(t *time.Duration) *RedisManger {
	m.t = t
	return m
}

func (m *RedisManger) MustBuild() *RedisManger {
	if !m.validMode() {
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
	i := string(ms)
	m.v = &i
	m.build = true
	return m
}

func (m *RedisManger) QMustBuild() *RedisManger {
	if !m.validMode() {
		return m
	}
	m.build = true
	return m
}

// 执行set
func (m *RedisManger) SetResult() bool {
	if !m.valid() {
		return false
	}
	if strs.IsEmpty(m.k) {
		return false
	}
	m.build = false
	if m.t != nil {
		return m.setEx()
	}
	return m.dfSet()
}
func (m *RedisManger) IncrResult() bool {
	if !m.validKey() {
		return false
	}
	return m.R.Incr(*m.k) > 0
}

// 执行hset
func (m *RedisManger) HSetResult() bool {
	if !m.hValid() {
		return false
	}

	r := m.dfHSet()
	m.build = false
	return r
}
func (m *RedisManger) HMSetResult() bool {
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
func (m *RedisManger) QueryResult() interface{} {
	if !m.valid() {
		return false
	}
	b, s := m.R.Get(*m.k)
	if !b {
		return nil
	}
	m.v = &s
	json.Unmarshal([]byte(s), m.tp)
	return m.tp
}
func (m *RedisManger) HQueryResult() interface{} {
	if !m.valid() {
		return false
	}
	s := m.R.HGet(*m.k, *m.f)
	if strs.IsDefault(s) {
		return m.tp
	}
	m.v = &s
	if err := json.Unmarshal([]byte(*m.v), m.tp); err != nil {
		return nil
	}
	return m.tp
}
func (m *RedisManger) HQueryResultVal() interface{} {
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}
	s := m.R.HGet(*m.k, *m.f)
	if strs.IsDefault(s) {
		return m.tp
	}
	m.v = &s

	return m.v
}

// 执行删除
func (m *RedisManger) DelResult() bool {
	if strs.IsEmpty(m.k) {
		return false
	}
	return m.R.Del(*m.k)
}

func (m *RedisManger) RunScriptResult(script string) interface{} {
	r := m.R.client.Eval(context.Background(), script, []string{*m.k}, *m.f, *m.v).Val()
	return r
}
func (m *RedisManger) StatScriptExpResult() bool {
	if !m.scriptValid() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.incrScript(), []string{*m.k}, *m.f, 1, m.formatSec(*m.t)).Int()
	if err != nil {
		logx.Errorf("StatIncr key:%v,field:%v, err:%v", *m.k, *m.f, err)
		return false
	}
	logx.Infof("StatScriptExpResult info:%v", v)
	return true
}
