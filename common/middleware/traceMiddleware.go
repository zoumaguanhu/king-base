package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type TraceIDMiddleware struct {
}

func NewTraceIDMiddleware() *TraceIDMiddleware {
	return &TraceIDMiddleware{}
}
func (m *TraceIDMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 生成一个新的 TraceID
		//traceID := uuid.New().String()

		// 将 TraceID 存储在请求上下文中
		//ctx := context.WithValue(r.Context(), "traceID", traceID)
		logx.WithContext(r.Context()).Infof("1111111111111111")
		//withContext.Infof("11111111111111111111")

		// 在日志中包含 TraceID
		//logx.AddGlobalFields(logx.LogField{Key: "TraceID", Value: traceID})

		// 继续处理请求
		next.ServeHTTP(w, r.WithContext(r.Context()))
	}
}
