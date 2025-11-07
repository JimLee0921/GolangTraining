// server/main.go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net"
	"net/rpc"
	"os"
)

type Cal struct{}

type Result struct {
	Num int
	Ans int
}

func (c *Cal) Square(n int, out *Result) error {
	out.Num = n
	out.Ans = n * n
	return nil
}

func main() {
	// 1) 载入服务端证书
	srvCert, err := tls.LoadX509KeyPair("08_advanced/02_rpc-and-ssl/04_mutual-tls/server/server.crt", "08_advanced/02_rpc-and-ssl/04_mutual-tls/server/server.key")
	if err != nil {
		log.Fatal("load server cert/key:", err)
	}

	// 2) 构造 ClientCAs：信任客户端的证书（此处直接信任对方的“叶子证书”）
	clientCAPool := x509.NewCertPool()
	clientBytes, err := os.ReadFile("08_advanced/02_rpc-and-ssl/04_mutual-tls/client/client.crt")
	if err != nil {
		log.Fatal("read client.crt:", err)
	}
	if !clientCAPool.AppendCertsFromPEM(clientBytes) {
		log.Fatal("append client.crt failed")
	}

	// 3) TLS 配置：要求并验证客户端证书
	cfg := &tls.Config{
		Certificates: []tls.Certificate{srvCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCAPool,
		MinVersion:   tls.VersionTLS12,
	}

	// 4) 监听 TLS
	ln, err := tls.Listen("tcp", ":1234", cfg)
	if err != nil {
		log.Fatal("tls listen:", err)
	}
	log.Println("RPC TLS server on :1234")

	// 5) 注册 RPC
	if err := rpc.Register(&Cal{}); err != nil {
		log.Fatal("rpc register:", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Println("temporary accept error:", err)
				continue
			}
			if err == io.EOF {
				continue
			}
			log.Fatal("accept:", err)
		}
		go rpc.ServeConn(conn)
	}
}
