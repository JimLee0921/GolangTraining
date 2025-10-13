package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"
)

/* ---------- 装饰器类型与链式组合 ---------- */

type Middleware func(http.Handler) http.Handler

// Chain Chain(m1, m2, m3)(h) == m1(m2(m3(h)))
func Chain(mw ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(mw) - 1; i >= 0; i-- {
			final = mw[i](final)
		}
		return final
	}
}

/* ---------- 常用中间件实现 ---------- */

type ctxKey string

const reqIDKey ctxKey = "reqID"

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := genID()
			ctx := context.WithValue(r.Context(), reqIDKey, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func LoggingWithTiming() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			id := getReqID(r.Context())
			log.Printf("[REQ %s] %s %s", id, r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
			log.Printf("[REQ %s] done in %v", id, time.Since(start))
		})
	}
}

func Recoverer() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					id := getReqID(r.Context())
					log.Printf("[REQ %s] panic: %v", id, rec)
					http.Error(w, "internal error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Timeout(d time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()

			done := make(chan struct{})
			go func() {
				next.ServeHTTP(w, r.WithContext(ctx))
				close(done)
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				id := getReqID(r.Context())
				log.Printf("[REQ %s] timeout after %v", id, d)
				http.Error(w, "request timeout", http.StatusGatewayTimeout)
			}
		})
	}
}

/* ---------- 业务 Handler ---------- */

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 模拟业务逻辑
	if r.URL.Query().Get("panic") == "1" {
		panic("boom")
	}
	if r.URL.Query().Get("sleep") != "" {
		time.Sleep(2 * time.Second)
	}
	id := getReqID(r.Context())
	fmt.Fprintf(w, "hello, reqID=%s\n", id)
}

/* ---------- 工具函数 ---------- */

func genID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func getReqID(ctx context.Context) string {
	if v := ctx.Value(reqIDKey); v != nil {
		return v.(string)
	}
	return "-"
}

/* ---------- 启动服务 ---------- */

func main() {
	// 组合装饰器：请求ID -> 日志+耗时 -> panic恢复 -> 1s超时
	decorators := Chain(
		RequestID(),
		LoggingWithTiming(),
		Recoverer(),
		Timeout(1*time.Second),
	)

	http.Handle("/hello", decorators(http.HandlerFunc(helloHandler)))

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
