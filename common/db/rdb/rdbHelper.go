package rdb

import (
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/strs"
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
