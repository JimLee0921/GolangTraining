package geerpc

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type Foo int

type Args struct {
	Num1, Num2 int
}

// 方法需要符合 RPC 调用规则
func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// 非导出方法不可以被 PRC 调用
func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func _assert(condition bool, msg string, v ...any) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

func TestNewServer(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	_assert(len(s.method) == 1, "wrong service Method. expect 1, but got %d", len(s.method))
	mType := s.method["Sum"]
	_assert(mType != nil, "wrong Method. Sum shouldn't nil")
}

func TestMethodType_Call(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	mType := s.method["Sum"]

	argv := mType.newArgv()
	log.Println(argv.Type())
	replyv := mType.newReplyV()
	argv.Set(reflect.ValueOf(Args{
		Num1: 1,
		Num2: 2,
	}))
	err := s.call(mType, argv, replyv)
	_assert(err == nil && *replyv.Interface().(*int) == 3 && mType.numCalls == 1, "failed to call Foo.Sum")
}
