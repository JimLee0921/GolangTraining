package parse

import (
	"encoding/binary"
	"errors"
)

// ParseUint32 解析函数（协议，header，自定义格式等）用于补 panic 错误是可控的
func ParseUint32(b []byte) (uint32, error) {
	if len(b) < 4 {
		return 0, errors.New("too short")
	}
	return binary.BigEndian.Uint32(b[:4]), nil
}
