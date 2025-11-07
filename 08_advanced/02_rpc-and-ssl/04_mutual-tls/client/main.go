// client/main.go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/rpc"
	"os"
)

type Result struct {
	Num int
	Ans int
}

func main() {
	// 1) 载入客户端证书（用于被服务器校验）
	cltCert, err := tls.LoadX509KeyPair("08_advanced/02_rpc-and-ssl/04_mutual-tls/client/client.crt", "08_advanced/02_rpc-and-ssl/04_mutual-tls/client/client.key")
	if err != nil {
		log.Fatal("load client cert/key:", err)
	}

	// 2) RootCAs：信任服务端（这里直接信任服务端的“叶子证书”）
	root := x509.NewCertPool()
	srvBytes, err := os.ReadFile("08_advanced/02_rpc-and-ssl/04_mutual-tls/server/server.crt")
	if err != nil {
		log.Fatal("read server.crt:", err)
	}
	if !root.AppendCertsFromPEM(srvBytes) {
		log.Fatal("append server.crt failed")
	}

	// 3) TLS 配置
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cltCert},
		RootCAs:      root,
		ServerName:   "localhost", // 必须与 server.crt 的 SAN 匹配
		MinVersion:   tls.VersionTLS12,
	}

	// 4) 拨号 + RPC
	conn, err := tls.Dial("tcp", "localhost:1234", cfg)
	if err != nil {
		log.Fatal("dial tls:", err)
	}
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	var out Result
	if err := client.Call("Cal.Square", 12, &out); err != nil {
		log.Fatal("rpc call:", err)
	}
	log.Printf("%d^2 = %d", out.Num, out.Ans)
}
