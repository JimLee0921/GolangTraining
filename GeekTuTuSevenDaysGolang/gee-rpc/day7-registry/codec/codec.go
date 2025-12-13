package codec

import "io"

// Header 代表 一次 RPC 调用的头部信息，请求和响应都会带一个 Header
type Header struct {
	ServiceMethod string // 服务名和方法名，相当于路由，映射到某个 Go 结构体的某个方法上
	Seq           uint64 // 客户端请求序号，区分不同请求
	Error         string // 错误信息，客户端置空，发生错误由服务端补充
}

// Codec 统一解码编码行为，抽象一套解码编码方案 RPC 消息在一个连接上的“收发器”：
// ReadHeader / ReadBody：负责收
// Write：负责发
// Close：负责关
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

// NewCodecFunc 构造函数类型，表示构造一个 Codec 的函数
// io.ReadWriteCloser 表示一个可以读写关的连接，比如 net.Conn
// 返回 Codec 是某种具体实现，比如 *GobCodec
type NewCodecFunc func(io.ReadWriteCloser) Codec

// Type 表示 Codec 类型，别名方便表达语义
type Type string

// JsonType 未实现但原理类似
const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // 未实现
)

// NewCodecFuncMap 注册表，工厂，全局变量
// key：对应 Type(比如 GobType, JsonType)
// value：对应的构造函数
// 可以理解为一个 Codec 工厂注册表，想要支持一个新的编码解码方式(比如Json)，需要以下操作：
// 1. 实现一个 JsonCodec，满足 Codec 接口
// 2. 写一个构造函数 NewJsonCodec(conn io.ReadWriteCloser)
// 3. 在 init() 中进行注册 NewCodecFuncMap[JsonType] = NewJsonCodec
var NewCodecFuncMap map[Type]NewCodecFunc

// 自动执行
func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
