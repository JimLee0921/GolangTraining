package geecache

// ByteView 表示缓存值
type ByteView struct {
	b []byte // 存储真实的缓存值，只读属性，所以绑定的方法都不是指针方法
}

// Len 方法返回 ByteView 所占内存大小
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice b 是只读的，所以返回一个拷贝，防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes 返回 b 的值拷贝
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
