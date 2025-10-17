package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	Pass   string `json:"-"`
	secret string
}

// main 流式编码
func main() {
	/*
		Marshal/Unmarshal：对象 <-> []byte，一次性在内存中完成；简单直观，但对大数据不友好
		Encoder/Decoder：对象 <-> io.Writer/io.Reader，边读边写，省内存、可持续读写多条记录，是生产场景更常用的形态

		Encoder：流式编码，把 Go 数据结构编码成 JSON 并写入到一个 流 (io.Writer) 里
		enc := json.NewEncoder(w) // w 是 io.Writer，比如 os.Stdout、文件、http.ResponseWriter
		err := enc.Encode(v)      // 把 v 编码成 JSON，写到 w
		json.NewEncoder(w)：创建一个绑定到 io.Writer 的编码器。
			Encode(v)：把 v 转成 JSON，写入 w
			会在 JSON 结尾自动加上一个换行符（方便连续写多条）
		注意事项：
			只有 导出字段（首字母大写）会被编码
			也可以用 struct tag 控制输出字段名、忽略字段
			比 Marshal 更适合文件、网络场景，支持逐条写 JSON，不用一次性生成整个 []byte
	*/
	user := User{
		ID:     101,
		Name:   "JimLee",
		Pass:   "58f4e8q2",
		secret: "faw2-fr42-pc2e-4i54",
	}
	encoder := json.NewEncoder(os.Stdout)

	err := encoder.Encode(user)
	if err != nil {
		log.Fatalln(err)
	}

}
