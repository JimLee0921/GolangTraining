# 安装 Go

> 官网地址: https://golang.google.cn/

## 下载

> 下载地址：https://golang.google.cn/dl/

### Windows 下载

直接选择 `Microsoft Windows` 安装包，下载后直接运行安装即可

### Linux 下载

包管理器下载（下载版本可能不会太新）

```bash
sudo apt install golang-go  # Ubuntu / Debian
sudo dnf install golang   # Fedora
sudo yum install golang   # CentOS
```

官方二进制安装（推荐，版本最新）

1. 下载 Go 安装包（1.25.1换为需要的版本）
    ```bash
    curl -o /tmp/go1.25.1.tar.gz https://go.dev/dl/go1.25.1.linux-amd64.tar.gz  # curl 拉取
    wegt -O /tmp/go1.25.1.tar.gz https://go.dev/dl/go1.25.1.linux-amd64.tar.gz  # wget 拉取
    ```
2. 解压到 /usr/local
    ```bash
    tar -C /usr/local -xzf /tmp/go1.25.1.linux-amd64.tar.gz
    ```
3. 配置环境变量
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```
4. 刷新环境变量
    ```bash
    source ~/.bashrc
    ```
5. 验证安装
    ```bash
    go version
    ```

> 可能出现网络原因，换成下面地址
> * https://mirrors.aliyun.com/golang/go1.25.1.linux-amd64.tar.gz
> * https://mirrors.tuna.tsinghua.edu.cn/go/go1.25.1.linux-amd64.tar.gz