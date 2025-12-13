package client

import (
	"context"
	"time"
)

type Client struct {
}

// Do 模拟一次耗时操作
func (client *Client) Do(ctx context.Context) error {
	// 模拟一个固定耗时的 RPC 调用操作
	select {
	case <-time.After(200 * time.Microsecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close 模拟关闭操作，这里直接延时空操作
func (client *Client) Close() {
	time.Sleep(time.Microsecond * 10)
}
