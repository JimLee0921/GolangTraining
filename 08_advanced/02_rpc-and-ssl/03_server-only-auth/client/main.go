package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	X, Y int
}
type Result struct {
	Product int
}

func main() {
	// 如果需要对服务器端鉴权，那么需要将服务端的证书添加到信任证书池中
	certPool := x509.NewCertPool()
	certBytes, err := os.ReadFile("08_advanced/02_rpc-and-ssl/03_server-only-auth/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.cert", err)
	}
	certPool.AppendCertsFromPEM(certBytes)

	// TLS 配置
	config := &tls.Config{RootCAs: certPool}

	// 建立 TLS 连接
	conn, err := tls.Dial("tcp", "localhost:1234", config) // result:  42

	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	defer conn.Close()

	// 创建 RPC 客户端
	client := rpc.NewClient(conn)
	var reply Result
	args := Args{
		X: 6,
		Y: 7,
	}
	err = client.Call("Cal.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}
	// 4. 查询结果
	fmt.Println("result: ", reply.Product) // result:  42

}
