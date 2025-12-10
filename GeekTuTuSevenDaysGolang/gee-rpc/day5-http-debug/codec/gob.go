package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser // 底层连接 一般是 net.Conn，只关心是能读能写还能关的东西
	buf  *bufio.Writer      // 包一层输出缓冲，调用 enc.Encode 其实是写到 buf 最后 Flush 才真正写入 conn 减少 write 调用次数，避免频繁小包写网络
	dec  *gob.Decoder       // 从 conn 上读取字节，按 gob 协议解码 再变回 Go 对象
	enc  *gob.Encoder       // 反之，从 Go 对象编码为 gob 字节，再写入 buf
}

// ReadHeader 调用 c.dec.Decode 读取一段 gob 数据，解码到传入的对象里
func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

// ReadBody 调用 c.dec.Decode 读取一段 gob 数据，解码到传入的对象里，body 类型由上层决定
func (c *GobCodec) ReadBody(body any) error {
	return c.dec.Decode(body)
}

// Write 写入数据
func (c *GobCodec) Write(h *Header, body any) (err error) {
	// 命名返回值可以在 defer 中直接操作
	defer func() {
		// 无论是否失败都进行 Flush 不造成数据残留污染
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	// 先写 header 再写 body
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header: ", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body: ", err)
		return err
	}
	return nil
}
func (c *GobCodec) Close() error {
	return c.conn.Close()
}

// 接口实现断言，编译器断言，编译期会进行检查（*GobCodec 一定是一个合法的 Codec 实现）
var _ Codec = (*GobCodec)(nil)

// NewGobCodec 构造函数，参数 conn 保持抽象，以后可以使用其它实现
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn) // 把 conn 包装成带缓冲 Writer 后面1 Encoder 写到 buf 上
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}
