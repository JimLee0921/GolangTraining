# Protobuf

在 Go（Golang）中，Protocol Buffers（简称 Protobuf） 是一种由 Google 开发的 跨语言、跨平台的数据序列化协议。
它的作用可以简单理解为：定义数据结构 + 高效地序列化 / 反序列化数据。
相比 JSON、XML 这类格式，Protobuf 更快、更紧凑、更类型安全，非常适合 RPC 通信、微服务、存储结构化数据 等场景。

Protocol Buffers 支持在 C++、C#、Dart、Go、Java、Kotlin、Objective-C、Python 和 Ruby 中生成代码。使用 proto3，protoc --version
还可以使用 PHP。

成 Go 代码、配合 gRPC 用）。
我来给你一个完整、可执行的安装流程，包括 **环境安装 + 验证 + Go 插件配置**。

## 安装步骤（跨平台详细）

要在 Go 里用 Protobuf，实际上要安装两部分：

| 组件                         | 作用                                          |
|----------------------------|---------------------------------------------|
| **protoc**                 | Protobuf 官方编译器，用来读取 `.proto` 文件并生成代码        |
| **protoc-gen-go**          | Go 的 Protobuf 插件，让 `protoc` 能输出 `.pb.go` 文件 |
| **protoc-gen-go-grpc**（可选） | 如果你用 gRPC，这个插件用来生成 gRPC 相关接口代码              |

### Windows 系统

#### 安装 `protoc` 主程序

进入官方发布页：
[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

选择与系统对应的版本，例如：

```
protoc-28.0-win64.zip
```

下载后解压，把 `bin/protoc.exe` 放到任意目录，然后把该目录加入系统环境变量 `PATH`。

验证：

```bash
protoc --version
```

输出类似：

```
libprotoc 28.0
```

表示安装成功。

---

#### 安装 Go 插件

执行：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

这两个命令会把插件安装到：

```
%USERPROFILE%\go\bin
```

> 请确保这个路径也加到了系统环境变量 PATH 中，否则 `protoc` 找不到插件。

验证：

```bash
protoc-gen-go --version
```

---

### Linux / macOS

#### 安装 protoc

使用包管理器：

**Ubuntu / Debian:**

```bash
sudo apt install -y protobuf-compiler
```

**macOS（推荐 Homebrew）:**

```bash
brew install protobuf
```

验证：

```bash
protoc --version
```

---

#### 安装 Go 插件

和 Windows 一样：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Go 的 `$GOPATH/bin` 通常在：

```
~/go/bin
```

要确保它在 PATH 里：

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

### 验证安装是否成功

创建一个简单的测试文件 `test.proto`：

```proto
syntax = "proto3";

package test;

message Hello {
  string name = 1;
}
```

然后运行：

```bash
protoc --go_out=. test.proto
```

如果成功生成 `test.pb.go` 文件，说明环境配置正确。

---

## 常用命令速查表

| 目的                      | 命令                                                                |
|-------------------------|-------------------------------------------------------------------|
| 查看 protoc 版本            | `protoc --version`                                                |
| 安装 Go 生成插件              | `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`  |
| 安装 gRPC 生成插件            | `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` |
| 编译 proto 为 Go 文件        | `protoc --go_out=. *.proto`                                       |
| 编译 proto 为 Go + gRPC 文件 | `protoc --go_out=. --go-grpc_out=. *.proto`                       |

## 基本使用

在 Protobuf 世界里，所有的数据结构都是通过一个 `xxx.proto` 文件定义的。
这个文件描述了你想要传输或存储的数据格式，然后通过 `protoc` 编译生成 Go 代码。

### 最小示例

一个最小可用的 `student.proto` 文件示例：

```proto
syntax = "proto3";
package main;
option go_package = "./main"

// this is a comment
message Student {
  string name = 1;
  bool male = 2;
  repeated int32 scores = 3;
}
```

- protobuf 有2个版本，默认版本是 proto2，如果需要 proto3，则需要在非空非注释第一行使用 `syntax = "proto3"` 标明版本
- package，Protobuf 世界里的命名空间，即包名声明符，是可选的，用来防止不同的消息类型有命名冲突
- `option go_package`：生成 Go 文件的包路径，设置关乎生成go文件后存放的路径
- 消息类型 使用 message 关键字定义，Student 是类型名，name, male, scores 是该类型的 3 个字段，类型分别为 string, bool 和
  `[]int32`。字段可以是标量类型，也可以是合成类型。
- 每个字段的修饰符默认是 singular，一般省略不写，repeated 表示字段可重复，即用来表示 Go 语言中的数组类型
- 每个字符 =后面的数字称为标识符，每个字段都需要提供一个唯一的标识符。标识符用来在消息的二进制格式中识别各个字段，一旦使用就不能够再改变，标识符的取值范围为
  `[1, 2^29 - 1]`
- `.proto` 文件可以写注释，单行注释 `//`，多行注释 `/* ... */`
- 一个 .proto 文件中可以写多个消息类型，即对应多个结构体(struct)

对应生成的 Go 结构体大致为：

```
type User struct {
Id       int32  `protobuf:"varint,1,opt,name=id,proto3"`
Name     string `protobuf:"bytes,2,opt,name=name,proto3"`
IsActive bool   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3"`
}
```

### 字段编号（Field Number）

每个字段后面的 `= 1`, `= 2`, `= 3` 是 **唯一编号**，用于序列化时标识字段。
这个编号不能重复，并且一旦发布后，不建议随意修改。

建议：

* 1～15：常用字段（占一个字节）
* 16～2047：一般字段

> 删除字段时不要重复使用编号，否则会造成旧数据解析错误。

---

## 数据类型

### 标量类型(Scalar)

| proto类型  | 	go类型	  | 备注                 | 	proto类型  | 	go类型    | 	备注                 |
|----------|---------|--------------------|-----------|----------|---------------------|
| double	  | float64 | 		                 | float     | 	float32 |                     |
| int32	   | int32   | 		                 | int64     | 	int64   |                     |
| uint32	  | uint32  | 		                 | uint64    | 	uint64  |                     |
| sint32	  | int32   | 	适合负数              | sint64    | 	int64   | 	适合负数               |
| fixed32  | 	uint32 | 	固长编码，适合大于2^28的值   | 	fixed64  | 	uint64  | 	固长编码，适合大于2^56的值    |
| sfixed32 | 	int32  | 	固长编码              | 	sfixed64 | 	int64   | 	固长编码               |        
| bool	    | bool    | 		                 | string    | 	string  | 	UTF8 编码，长度不超过 2^32 |
| bytes    | 	[]byte | 	任意字节序列，长度不超过 2^32 |           |          |                     |

> 标量类型如果没有被赋值，则不会被序列化，解析时，会赋予默认值。
>- strings：空字符串
>- bytes：空序列
>- bools：false
>- 数值类型：0

---

### 嵌套与引用类型

Protobuf 支持嵌套类型（类似结构体嵌套）

```proto
message Address {
  string city = 1;
  string country = 2;
}

message User {
  int32 id = 1;
  string name = 2;
  Address address = 3; // 嵌套引用
}
```

生成的 Go 代码会自动把 `Address` 作为一个子结构体字段。

---

### 数组（repeated）

要定义一个列表，用 `repeated`：

```proto
message Order {
  int32 id = 1;
  repeated string items = 2; // 等价于 []string
}
```

生成的 Go 代码：

```
type Order struct {
Id    int32
Items []string
}
```

---

### 枚举(Enumerations)

枚举类型适用于提供一组预定义的值，选择其中一个。例如将性别定义为枚举类型。

```proto
message Student {
  string name = 1;
  enum Gender {
    FEMALE = 0;
    MALE = 1;
  }
  Gender gender = 2;
  repeated int32 scores = 3;
}
```

- 枚举类型的第一个选项的标识符必须是0，这也是枚举类型的默认值
- 别名（Alias），允许为不同的枚举值赋予相同的标识符，称之为别名，需要打开allow_alias选项

```proto
message EnumAllowAlias {
  enum Status {
    option allow_alias = true;
    UNKOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
```

> Go 代码会生成对应的枚举常量

### 任意类型(Any)

Any 可以表示不在 .proto 中定义任意的内置类型。

```proto
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

### oneof

```proto
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

### map

```proto
message MapRequest {
  map<string, int32> points = 1;
}
```

## 定义服务（services）

如果消息类型是用来远程通信的(Remote Procedure Call, RPC)，可以在 .proto 文件中定义 RPC 服务接口。例如定义了一个名为
SearchService 的 RPC 服务，提供了 Search 接口，入参是 SearchRequest 类型，返回类型是 SearchResponse

```
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```

> 官方仓库也提供了一个[插件列表](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md_)，帮助开发基于
> Protocol Buffer 的 RPC 服务。

## 推荐风格

1. 文件(Files)

    - 文件名使用小写下划线的命名风格，例如 lower_snake_case.proto
    - 每行不超过 80 字符
    - 使用 2 个空格缩进
2. 包(Packages)

    - 包名应该和目录结构对应，例如文件在my/package/目录下，包名应为 my.package
3. 消息和字段(Messages & Fields)

    - 消息名使用首字母大写驼峰风格(CamelCase)，例如`message StudentRequest { ... }`
    - 字段名使用小写下划线的风格，例如 `string status_code = 1`
    - 枚举类型，枚举名使用首字母大写驼峰风格，例如 `enum FooBar`，枚举值使用全大写下划线隔开的风格(
      CAPITALS_WITH_UNDERSCORES )，例如 FOO_DEFAULT=1
4. 服务(Services)
    - RPC 服务名和方法名，均使用首字母大写驼峰风格，例如`service FooService{ rpc GetSomething() }`

> GitHub地址：https://github.com/protocolbuffers/protobuf
>
> 官方文档：https://protobuf.dev/
