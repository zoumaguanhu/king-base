package rdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/constants"
	"king.com/king/base/common/convert"
	"king.com/king/base/common/strs"
	"time"
)

type RdbBaseImpl struct {
	R   *RedisClient
	ctx context.Context
}

func (r *RdbBaseImpl) BuildKey(key, bz string) string {
	s := fmt.Sprintf("site:%v:%v", key, bz)
	return s
}
func (r *RdbBaseImpl) BuildHostKey(host, bz string) string {
	s := fmt.Sprintf("site:vhost:%v:%v", host, bz)
	return s
}
func (r *RdbBaseImpl) BuildAdminKey(admin, bz string) string {
	s := fmt.Sprintf("site:admin:%v:%v", admin, bz)
	return s
}
func (r *RdbBaseImpl) BuildClientKey(host, client, bz string) string {
	s := fmt.Sprintf("site:vhost:%v:client:%v:%v", host, client, bz)
	return s
}
func (r *RdbBaseImpl) BuildHostKeyWithDate(host, date, bz string) string {
	s := fmt.Sprintf("site:vhost:%v:%v:%v", host, date, bz)
	return s
}
func (r *RdbBaseImpl) BuildSiteKey(virtId int64, bz string) string {
	s := fmt.Sprintf("site:vsite:%v:%v", virtId, bz)
	return s
}
func (r *RdbBaseImpl) BuildBannerKey(virtId int64, bz string) string {
	s := fmt.Sprintf("site:vsite:%v:product:banner:%v", virtId, bz)
	return s
}
func (r *RdbBaseImpl) BuildSiteProductKey(virtId int64, bz string) string {
	s := fmt.Sprintf("site:vsite:%v:product:page:%v", virtId, bz)
	return s
}
func (r *RdbBaseImpl) BuildUserKey(virtId int64, userId int64, bz string) string {
	s := fmt.Sprintf("site:vsite:%v:member:%v:%v", virtId, userId, bz)
	return s
}

func (r *RdbBaseImpl) BuildEmailKey(virtId int64, email string, bz string) string {
	s := fmt.Sprintf("site:vsite:%v:email:%v:%v", virtId, email, bz)
	return s
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
	k     string         //key
	f     string         //field
	v     *string        // value
	t     *time.Duration //超时
	step  int64
	build bool //构建完成后放可执行
}

func NewRM(R *RedisClient) *RedisManger {
	r := &RedisManger{}
	r.R = R
	return r
}
func (m *RedisManger) WithCtx(ctx context.Context) *RedisManger {
	m.ctx = ctx
	return m
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
func (m *RedisManger) WithCustomKey(key string) *RedisManger {
	m.k = key
	return m
}
func (m *RedisManger) WithSiteKey(virtId int64, bz string) *RedisManger {
	m.k = m.BuildSiteKey(virtId, bz)
	return m
}
func (m *RedisManger) WithSiteProductKey(virtId int64, bz string) *RedisManger {
	m.k = m.BuildSiteProductKey(virtId, bz)
	return m
}
func (m *RedisManger) WithHostKey(host, bz string) *RedisManger {
	m.k = m.BuildHostKey(host, bz)
	return m
}
func (m *RedisManger) WithAdminKey(admin, bz string) *RedisManger {
	m.k = m.BuildAdminKey(admin, bz)
	return m
}

func (m *RedisManger) WithClientKey(host, client, bz string) *RedisManger {
	m.k = m.BuildClientKey(host, client, bz)
	return m
}
func (m *RedisManger) WithHostKeyWithDate(host, date, bz string) *RedisManger {
	m.k = m.BuildHostKeyWithDate(host, date, bz)
	return m
}
func (m *RedisManger) WithUserKey(virtId int64, userId int64, bz string) *RedisManger {
	m.k = m.BuildUserKey(virtId, userId, bz)
	return m
}
func (m *RedisManger) WithBannerKey(virtId int64, bz string) *RedisManger {
	m.k = m.BuildBannerKey(virtId, bz)
	return m
}
func (m *RedisManger) WithEmailKey(virtId int64, email string, bz string) *RedisManger {
	m.k = m.BuildEmailKey(virtId, email, bz)
	return m
}
func (m *RedisManger) WithField(f string) *RedisManger {
	m.f = f
	return m
}
func (m *RedisManger) WithExp(t *time.Duration) *RedisManger {
	m.t = t
	return m
}
func (m *RedisManger) WithVal(v string) *RedisManger {
	m.v = &v
	return m
}
func (m *RedisManger) WithStep(step int64) *RedisManger {
	m.step = step
	return m
}

func (m *RedisManger) MustBuild() *RedisManger {
	if m.tp != nil {
		if err := copier.CopyWithOption(m.tp, m.s, convert.Time2DefaultFormatStr()); err != nil {
			logc.Errorf(m.ctx, "MustBuild copy err:%v,tp:%+v", err, m.tp)
			return m
		}
		ms, err := json.Marshal(m.tp)
		if err != nil {
			logc.Errorf(m.ctx, "MustBuild marshal err:%v,tp:%+v", err, m.tp)
			return m
		}
		i := string(ms)
		m.v = &i
	}

	m.build = true
	return m
}

func (m *RedisManger) QMustBuild() *RedisManger {

	m.build = true
	return m
}

func (m *RedisManger) SetResult() bool {
	if !m.valid() {
		return false
	}
	if strs.IsDefault(m.k) {
		return false
	}
	m.build = false
	if m.t != nil {
		return m.setEx()
	}
	return m.dfSet()
}

func (m *RedisManger) IncrResult() (bool, int64) {
	if !m.validKey() {
		return false, constants.ZERO_INT64
	}
	v := m.R.Incr(m.k)
	return v > 0, v
}
func (m *RedisManger) QIncrResult() int64 {
	if !m.validKey() {
		return 0
	}
	b, s := m.R.Get(m.k)
	if !b {
		return 0
	}
	return strs.StrsToInt64(s)
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

func (m *RedisManger) HMSetResult(data *map[string]interface{}) bool {
	if !m.validKey() {
		return false
	}
	if strs.IsDefault(m.k) {
		return false
	}
	r := m.hMSet(data)
	m.build = false
	return r
}

// 执行查询
func (m *RedisManger) QueryResult() interface{} {
	if !m.qValid() {
		return false
	}
	b, s := m.R.Get(m.k)
	if !b {
		return nil
	}
	m.v = &s
	json.Unmarshal([]byte(s), m.tp)
	return m.tp
}
func (m *RedisManger) QueryVal() interface{} {
	if !m.qValid() {
		return false
	}
	b, s := m.R.Get(m.k)
	if !b {
		return nil
	}
	m.v = &s
	return m.v
}
func (m *RedisManger) HQueryResult() interface{} {
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}
	s := m.R.HGet(m.k, m.f)
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
	s := m.R.HGet(m.k, m.f)
	if strs.IsDefault(s) {
		return nil
	}
	m.v = &s

	return m.v
}

// 执行删除
func (m *RedisManger) DelResult() bool {
	if !m.validKey() {
		return false
	}
	return m.R.Del(m.k)
}

// 执行删除
func (m *RedisManger) HDelResult() bool {
	if !m.validKey() {
		return false
	}
	if !m.validField() {
		return false
	}
	return m.R.HDel(m.k, m.f)
}
func (m *RedisManger) ZCountResult() (int, bool) {
	if !m.validKey() {
		return 0, false
	}
	v, err := m.R.client.Eval(context.Background(), *m.zSetCountScript(), []string{m.k}).Int()
	if err != nil {
		logc.Errorf(m.ctx, "StatIncr key:%v,field:%v, err:%v", m.k, m.f, err)
		return 0, false
	}
	logc.Infof(m.ctx, "StatScriptExpResult info:%v", v)
	return v, true
}
func (m *RedisManger) KeyReName(newKey string) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Rename(m.ctx, m.k, newKey).Result()
	if err != nil {
		logc.Errorf(m.ctx, "KeyReName key:%v,field:%v, err:%v", m.k, m.f, err)
		return false
	}
	logc.Infof(m.ctx, "KeyReName info:%v", v)
	return v == newKey
}
func (m *RedisManger) KeyReNameNx(newKey string) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.RenameNX(m.ctx, m.k, newKey).Result()
	if err != nil {
		logc.Errorf(m.ctx, "KeyReName key:%v,field:%v, err:%v", m.k, m.f, err)
		return false
	}
	logc.Infof(m.ctx, "KeyReName info:%v", v)
	return v
}
func (m *RedisManger) RunScriptResult(script string) interface{} {
	r := m.R.client.Eval(context.Background(), script, []string{m.k}, m.f, *m.v).Val()
	return r
}
func (m *RedisManger) StatHScriptResult() bool {
	if !m.scriptValid() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.incrHScript(), []string{m.k}, m.f, 1, m.formatSec(*m.t)).Int()
	if err != nil {
		logc.Errorf(m.ctx, "StatIncr key:%v,field:%v, err:%v", m.k, m.f, err)
		return false
	}
	logc.Infof(m.ctx, "StatScriptExpResult info:%v", v)
	return true
}
func (m *RedisManger) StatScriptResult() (bool, int) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.incrScript(), []string{m.k}, m.formatSec(*m.t), m.step).Int()
	if err != nil {
		logc.Errorf(m.ctx, "StatIncr key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "StatScriptExpResult info:%v", v)
	return true, v
}
func (m *RedisManger) ProductPageScriptResult(start int64, end int64) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.ProductPageScript(), []string{m.k}, start, end).Result()
	if err != nil {
		logc.Errorf(m.ctx, "ProductPageScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "ProductPageScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) ProductListScriptResult(ids *[]int64) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.ProductListScript(), []string{m.k}, ids).Result()
	if err != nil {
		logc.Errorf(m.ctx, "ProductListScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "ProductListScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) StockProductScriptResult(keys *[]string, stocks *[]string) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.reStoreScript(), *keys, *stocks).Result()
	if err != nil {
		logc.Errorf(m.ctx, "StockProductScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "StockProductScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) CheckStockProductScriptResult(keys *[]string, stocks *[]string) (bool, interface{}) {
	v, err := m.R.client.Eval(context.Background(), *m.checkReStoreScript(), *keys, *stocks).Result()
	if err != nil {
		logc.Errorf(m.ctx, "StockProductScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "StockProductScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) AddProductScriptResult(data any, sort int64, id string, dKey string, dData string) (bool, int) {
	if !m.validKey() {
		return false, 0
	}
	d := strs.ObjToStr(data)
	v, err := m.R.client.Eval(context.Background(), *m.addProductScript(), []string{m.k, id, dKey}, sort, d, dData).Int()
	if err != nil {
		logc.Errorf(m.ctx, "AddProductScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	return true, v
}
func (m *RedisManger) DelProductScriptResult(id, dKey string) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.delProductScript(), []string{m.k, dKey}, id).Result()
	if err != nil {
		logx.Errorf("DelProductScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Errorf(m.ctx, "DelProductScriptResult info:%v", v)
	return true, v
}

func (m *RedisManger) AddBannerScriptResult(id string, sort int64, data string) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.addBannerScript(), []string{m.k, id}, sort, data).Int()
	if err != nil {
		logx.Errorf("AddBannerScriptResult key:%v,sort:%v,value:%v, err:%v", m.k, sort, data, err)
		return false
	}
	logc.Errorf(m.ctx, "AddBannerScriptResult info:%v", v)
	return v > 0
}
func (m *RedisManger) BannerListScriptResult() (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.bannerListScript(), []string{m.k}).Result()
	if err != nil {
		logx.Errorf("BannerListScriptResult key:%v, err:%v", m.k, err)
		return false, 0
	}
	logc.Errorf(m.ctx, "BannerListScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) DelBannerScriptResult(id int64) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.delBannerScript(), []string{m.k}, id).Int()
	if err != nil {
		logx.Errorf("DelBannerScriptResult key:%v,sort:%v err:%v", m.k, id, err)
		return false
	}
	logc.Errorf(m.ctx, "DelBannerScriptResult info:%v", v)
	return v > 0
}
func (m *RedisManger) HAllQResult() *map[string]string {
	if !m.validKey() {
		return nil
	}
	all := m.R.HGetAll(m.k)
	return &all
}
func (m *RedisManger) GetInt(ms *map[string]string, f string) int64 {
	if c, ok := (*ms)[f]; ok {
		return strs.StrsToInt64(c)
	}
	return 0
}

func (m *RedisManger) CartPageScriptResult(start int64, end int64) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.CartPageScript(), []string{m.k}, start, end).Result()
	if err != nil {
		logc.Errorf(m.ctx, "CartPageScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "CartPageScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) ReNameCartPageScriptResult(start int64, end int64) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.ReNameCartPageScript(), []string{m.k}, start, end).Result()
	if err != nil {
		logc.Errorf(m.ctx, "CartPageScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "CartPageScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) AddCartScriptResult(data any, sort int64, combId string) (bool, int) {
	if !m.validKey() {
		return false, 0
	}
	d := strs.ObjToStr(data)
	v, err := m.R.client.Eval(context.Background(), *m.addCartScript(), []string{m.k, combId}, sort, d, m.formatSec(*m.t)).Int()
	if err != nil {
		logc.Errorf(m.ctx, "AddCartScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Errorf(m.ctx, "AddCartScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) DelCartScriptResult(combId string) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.delCartScript(), []string{m.k}, combId).Result()
	if err != nil {
		logx.Errorf("DelProductScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Errorf(m.ctx, "DelProductScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) DelUserCartScriptResult() (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.delUserCartScript(), []string{m.k}).Result()
	if err != nil {
		logx.Errorf("DelUserCartScriptResult key:%v, err:%v", m.k, err)
		return false, 0
	}
	logc.Errorf(m.ctx, "DelUserCartScriptResult info:%v", v)
	return true, v
}

func (m *RedisManger) AddSiteAndOptionsResult(data *map[string]interface{}) bool {
	if !m.validKey() {
		return false
	}
	// 构建参数列表
	args := make([]interface{}, 0, len(*data)*2+1)

	// 展开map为键值对列表
	for k, v := range *data {
		args = append(args, k, v)
	}

	// 添加过期时间参数
	args = append(args, m.formatSec(*m.t))
	cmd := m.R.client.Eval(context.Background(), *m.hMSetExpScript(), []string{m.k}, args...)
	if cmd.Err() != nil {
		logc.Errorf(m.ctx, "AddSiteAndOptionsResult key:%v,field:%v, err:%v", m.k, m.f, cmd.Err())
		return false
	}
	logc.Errorf(m.ctx, "AddSiteAndOptionsResult info:%+v", data)
	return true
}
func (m *RedisManger) AddBlogScriptResult(id string, detailKey string, sort int64, data string, dData string) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.addBlogScript(), []string{m.k, id, detailKey}, sort, data, dData).Int()
	if err != nil {
		logx.Errorf("AddBlogScriptResult key:%v,sort:%v,value:%v, err:%v", m.k, sort, data, err)
		return false
	}
	logc.Errorf(m.ctx, "AddBlogScriptResult info:%v", v)
	return v > 0
}

func (m *RedisManger) BlogPageScriptResult(start int64, end int64) (bool, interface{}) {
	if !m.validKey() {
		return false, 0
	}
	v, err := m.R.client.Eval(context.Background(), *m.blogPageScript(), []string{m.k}, start, end).Result()
	if err != nil {
		logc.Errorf(m.ctx, "BlogPageScriptResult key:%v,field:%v, err:%v", m.k, m.f, err)
		return false, 0
	}
	logc.Infof(m.ctx, "BlogPageScriptResult info:%v", v)
	return true, v
}
func (m *RedisManger) DelBlogScriptResult(id int64, detailKey string) bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.delBlogScript(), []string{m.k, detailKey}, id).Int()
	if err != nil {
		logx.Errorf("DelBlogScriptResult key:%v,sort:%v err:%v", m.k, id, err)
		return false
	}
	logc.Errorf(m.ctx, "DelBlogScriptResult info:%v", v)
	return v > 0
}
func (m *RedisManger) ReCartKeyNameToOrderName() bool {
	if !m.validKey() {
		return false
	}
	v, err := m.R.client.Eval(context.Background(), *m.reCartNameScript(), []string{m.k}).Int()
	if err != nil {
		logx.Errorf("DelBlogScriptResult key:%v, err:%v", m.k, err)
		return false
	}
	logc.Errorf(m.ctx, "DelBlogScriptResult info:%v", v)
	return v > 0
}
