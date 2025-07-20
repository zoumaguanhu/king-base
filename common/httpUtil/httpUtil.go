package httpUtil

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"king.com/king/base/common/constants"
	"king.com/king/rpc-oms/pkg/xerror"
	"net/http"
	"time"
)

func Send(ctx context.Context, r *http.Request, v any) error {
	client := &http.Client{
		Timeout: time.Duration(20) * time.Second,
	}
	//发起请求
	logc.Infof(ctx, "request method:%v,url:%v,req:%v", r.Method, r.URL.Path, r)
	now := time.Now()
	result, err := client.Do(r)

	//处理响应结果
	if result == nil || !(ok(result)) {
		logc.Errorf(ctx, "request method:%v,url:%v,holdTime:%v,result:%v", r.Method, r.URL, time.Since(now).String(), result)
		return xerror.New(xerror.THIRD_ERROR)
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		logc.Errorf(ctx, "resp body parse byte err:%v", err)
		return err
	}
	if err1 := json.Unmarshal(body, v); err1 != nil {
		return err1
	}
	return nil
}
func ok(resp *http.Response) bool {
	return resp != nil && (resp.StatusCode == constants.HTTP_OK || resp.StatusCode == constants.HTTP_201)
}
