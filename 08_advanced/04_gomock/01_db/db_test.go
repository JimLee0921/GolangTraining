package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	// 1 创建 gomock 控制器：管理本次测试里的所有 mock
	ctrl := gomock.NewController(t)
	// 2 测试结束时收尾：核对所有“期望的调用”是否都发生过（次数、参数都要匹配）
	defer ctrl.Finish() // 使用 ctrl.Finish() 断言 DB.Get() 被是否被调用，如果没有被调用，后续的 mock 就失去了意义
	// 3 基于接口 DB 生成的 mock 类型，传入控制器得到一个 mock 实例
	m := NewMockDB(ctrl)
	/*
		4 录制“期望”：当调用 m.Get("Tom") 时，返回 (100, errors.New("not exist"))
			- EXPECT() 进入录制模式，返回 recorder
			- recorder.Get(...)：由 mockgen 自动生成，用来声明“会被怎样调用”
			- gomock.Eq("Tom")：参数匹配器，要求参数严格等于 Tom
			- Return(...)：声明回放时给出的返回值（顺序对应接口签名）
	*/
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	/*
		5 调用被测函数：它接收 DB 接口，因此可传入定义的 mock
			GetFromDB 内部逻辑是：
			value, err := db.Get(key)
			如果 err == nil，就返回 value；否则返回 -1
	*/
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
