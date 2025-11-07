## Go RPC & TLS 鉴权

具体 go 中使用 RPC 见 stdlib 的 net 下的 rpc 模块

> 学习文档：https://geektutu.com/post/quick-go-rpc.html


虽然可以直接使用 `HandleHTTP` 开始 HTTP 服务，但是 HTTP 协议默认是不加密的，可以使用证书来保证通信过程的安全。
生成私钥和自签名的证书，并将 server.key 权限设置为只读，保证私钥的安全。

```
# 生成私钥
openssl genrsa -out server.key 2048
# 生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650
# 只读权限
chmod 400 server.key
```

执行完，当前文件夹下多出了 server.crt 和 server.key 2 个文件。

在使用案例时因为 较新的 Go（>=1.18）默认不再信任只靠 CN 的证书，必须在证书里配置 SAN（Subject Alternative
Name）。所以上面命令传创建的证书和密钥可能不能直接用了
使用下面命令进行创建：

```
openssl req -x509 -newkey rsa:2048 -sha256 -days 365 -nodes \
-keyout server.key -out server.crt \
-subj "/CN=localhost" \
-addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
```

> windows 上使用 bash 就在 /C 前面再加个 / 转义

### 双向鉴权

TLS 双向鉴权，也就是 mTLS（Mutual TLS）双向 TLS
服务器端对客户端的鉴权是类似的，核心在于 tls.Config 的配置
客户端增加

```
Certificates: []tls.Certificate{clientCert},  // 客户端证明自己是谁
```

服务器增加

```
ClientAuth: tls.RequireAndVerifyClientCert, // 服务器要求验证客户端证书
ClientCAs:  clientCertPool,                 // 服务器信任客户端
```

见 mutual-tls，其中生成服务端的证书和客户端证书命令如下：

```
server

openssl req -x509 -newkey rsa:2048 -sha256 -days 365 -nodes \
  -keyout server.key -out server.crt \
  -subj "/CN=localhost" \
  -addext "subjectAltName = DNS:localhost,IP:127.0.0.1"


client

openssl req -x509 -newkey rsa:2048 -sha256 -days 365 -nodes \
  -keyout client.key -out client.crt \
  -subj "/CN=dev-client" \
  -addext "subjectAltName = DNS:dev-client" \
  -addext "extendedKeyUsage = clientAuth"
```