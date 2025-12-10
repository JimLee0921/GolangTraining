package geerpc

import (
	"go/ast"
	"log"
	"reflect"
	"sync/atomic"
)

// 通过反射实现 service
type methodType struct {
	method    reflect.Method // 反射得到的 reflect.Method 本体(可以调用里面的 Func 字段)
	ArgType   reflect.Type   // rpc 调用时第一个参数
	ReplyType reflect.Type   // rpc 调用时第二个参数必须为指针类型，为了写入数据
	numCalls  uint64         // 统计被调用了多少次
}

// NumCalls 外部调用返回调用次数
func (m *methodType) NumCalls() uint64 {
	// 用 atomic.LoadUint64 来保证后面多 goroutine 统计时的原子性
	return atomic.LoadUint64(&m.numCalls)
}

// newArgv 根据方法入参类型 ArgType，创建一个合适的、可写的参数值 argv，后面好用反射/codec 去填充它
func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	if m.ArgType.Kind() == reflect.Ptr {
		// 方法参数是指针类型，返回对应数据的指针类型 Value
		argv = reflect.New(m.ArgType.Elem())
	} else {
		// 方法参数是值类型，返回对应数据的值类型 Value
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

// newReplyV 创建 reply 对象（必须为一个指针类型）并初始化 map/slice
func (m *methodType) newReplyV() reflect.Value {
	// reflect.New(m.ReplyType.Elem()) 只会分配一个 零值指针 nil，不会初始化 map 或 slice 的底层内存
	replyv := reflect.New(m.ReplyType.Elem())
	// 需要再通过类型判断进行初始化
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return replyv
}

// service 描述一整个服务对象
type service struct {
	name   string                 // 服务名，即结构体名字
	typ    reflect.Type           // 结构体反射类型 	reflect.TypeOf(recv)
	recv   reflect.Value          // 结构体实例的反射值	reflect.ValueOf(recv)
	method map[string]*methodType // 方法名到 *methodType 的映射
}

// newService 传入任意实例对象变为 service 并注册
func newService(recv any) *service {
	s := new(service)
	s.recv = reflect.ValueOf(recv)
	s.name = reflect.Indirect(s.recv).Type().Name()
	s.typ = reflect.TypeOf(recv)
	// 检查结构体 name 是否为首字母大写，如果不是导出类型不能通过 rpc 调用
	if !ast.IsExported(s.name) {
		log.Fatalf("rpc server: %s not a valid service name", s.name)
	}
	s.registerMethods()
	return s
}

// registerMethods 筛选结构体中合法的 RPC 方法
func (s *service) registerMethods() {
	s.method = make(map[string]*methodType)
	for i := 0; i < s.typ.NumMethod(); i++ {
		method := s.typ.Method(i)
		mType := method.Type
		// 检查是否为三个入参(本身相当于this/self+参数+返回值指针)和一个出参
		if mType.NumIn() != 3 || mType.NumOut() != 1 {
			continue
		}
		// 检查返回值类型是否为 error
		if mType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			continue
		}
		argType, replyType := mType.In(1), mType.In(2)
		if !isExportedOrBuiltinType(argType) || !isExportedOrBuiltinType(replyType) {
			continue
		}
		s.method[method.Name] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
		log.Printf("rpc server: register %s.%s", s.name, method.Name)
	}
}

// 辅助函数，判断一个类型是否为首字母大写到处类型和PkgPath为空字符串（识别 int、string 之类）
func isExportedOrBuiltinType(t reflect.Type) bool {
	return ast.IsExported(t.Name()) || t.PkgPath() == ""
}

// call 真正调用方法
func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	// 调用次数加1
	atomic.AddUint64(&m.numCalls, 1)
	// 获取 reflect.Value 形式的函数
	f := m.method.Func // 获取到的是 unbound function (没有 receiver)
	// m.method.Func 是一个反射函数，相当于一个普通函数，所以需要手动传入第一个隐含的参数 receiver（结构体实例或指针）
	returnValues := f.Call([]reflect.Value{s.recv, argv, replyv})
	// 返回值只关心第一个，转为 Go 值后如果非 nil 就断言为 error 返回
	if errInter := returnValues[0].Interface(); errInter != nil {
		return errInter.(error)
	}
	return nil
}
