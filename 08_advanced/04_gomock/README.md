# gomock

`Go Test` 是 Go 语言中单元测试的常用方法，包括子测试(subtests)、表格驱动测试(table-driven tests)、帮助函数(helpers)
、网络测试和基准测试(Benchmark)等。

`gomock` 是 `mock/stub` 测试：当待测试的函数/对象的依赖关系很复杂，并且有些依赖不能直接创建，例如数据库连接、文件I/O等。
这种场景就非常适合使用 `mock/stub` 测试。简单来说，就是用 `mock` 对象模拟依赖项的行为。

## 简介

`gomock` 是 Go 官方维护的 Mock 框架，用于在单元测试中模拟接口（interface）行为。

- 它帮你假装一个依赖对象
- 不需要真正调用数据库、HTTP、第三方服务
- 只验证代码是否正确地调用了接口

### 应用场景

| 场景                | 示例                     |
|-------------------|------------------------|
| 🔌 测试依赖外部接口的逻辑    | 模拟 HTTP 客户端、数据库、RPC 服务 |
| 🧪 测试函数调用次数、顺序、参数 | 验证被测函数是否按预期调用依赖        |
| ⚙️ 断开复杂依赖         | 解耦模块测试，避免集成测试复杂度       |

### 核心组件

| 组件        | 包路径                       | 功能                |
|-----------|---------------------------|-------------------|
| `gomock`  | `go.uber.org/mock/gomock` | 控制 Mock 生命周期、断言调用 |
| `mockgen` | 命令行工具                     | 自动根据接口生成 Mock 代码  |

### 下载安装

```sh
# 安装库
go get github.com/golang/mock/gomock

# 安装命令行工具
go install github.com/golang/mock/mockgen@latest


```

安装完成后运行 `mockgen -h` 如果能看到帮助说明就表示安装成功

### 生成 mock

**source 模式（最常用）**

```
mockgen -source=your_file.go -destination=mock_your_file_test.go -package=yourpkg
```

| 参数             | 含义                         |
|----------------|----------------------------|
| `-source`      | 指定要扫描接口的源文件                |
| `-destination` | 输出文件路径                     |
| `-package`     | 生成 mock 文件使用的包名（建议与测试文件同包） |

**reflect 模式（包路径 + 接口名）**

```
mockgen pkg/path InterfaceName
```

> 例子：mockgen database/sql/driver Conn,Driver > mock_driver_test.go

- 从已编译的包 database/sql/driver 中
- 反射出接口 Conn 和 Driver
- 并生成相应的 mock 实现
- 常用于 第三方包 或 标准库接口

**`//go:generate` 自动生成**

可以在接口文件头部加一行：`//go:generate mockgen -source=user.go -destination=mock_user_test.go -package=user`

之后只需执行： `go generate ./...`

Go 会自动调用 `mockgen`，生成 mock 文件。 适合团队协作与 CI/CD 自动化。

**参数**

| 选项                             | 含义                                        |
|--------------------------------|-------------------------------------------|
| `-source`                      | 指定源文件（source 模式）                          |
| `-destination`                 | 输出文件路径                                    |
| `-package`                     | 输出文件的包名                                   |
| `-imports`                     | 自定义导入别名映射                                 |
| `-aux_files`                   | 提供额外依赖文件（跨文件接口定义时使用）                      |
| `-copyright_file`              | 为生成文件加版权头                                 |
| `-self_package`                | 当前包的导入路径（解决 import 冲突时用）                  |
| `-write_package_comment=false` | 不输出 package 注释                            |
| `-mock_names`                  | 重命名生成的 mock，例如 `Interface=CustomMockName` |

## 编写可 mock 代码

想要能 mock 的代码，核心是可替换依赖。在 Go 里，gomock 只能 mock
接口（interface），不能直接替换具体类型、函数、全局变量。因此代码需要围绕以下原则来写（设计可测试/可替换的缝 seams）。

### 依赖倒置

面向接口编程 + 依赖注入，在使用方定义小接口（而不是在提供方/第三方里定义），把真正需要的那几个方法抽出来。
通过构造函数注入依赖（而不是在函数内部创建具体实现或用全局单例）。

```
// 使用方所在包定义接口（小而专一）
type UserStore interface {
    GetUser(ctx context.Context, id int64) (*User, error)
}

type UserService struct {
    store UserStore // 依赖接口，而不是具体 *sql.DB
}

func NewUserService(store UserStore) *UserService { // 构造函数注入
    return &UserService{store: store}
}

func (s *UserService) GetName(ctx context.Context, id int64) (string, error) {
    u, err := s.store.GetUser(ctx, id)
    if err != nil {
        return "", err
    }
    return u.Name, nil
}
```

### 小接口优先（Interface Segregation）

把依赖拆成最小必要方法集，避免一个大而全的接口到处传，mock 时更简单。

典型例子：

- io.Reader / io.Writer / io.ReadWriter
- interface{ Do(*http.Request) (*http.Response, error) }（见下）

### 为第三方/系统资源添加适配层

第三方库/具体类型不可直接 `mock`，但是可以包一层自己的接口，业务只依赖自己编写的接口

```text
// 定义自己的最小接口（而不是直接用 *http.Client）
type HTTPDoer interface {
    Do(req *http.Request) (*http.Response, error)
}

type APIClient struct {
    doer HTTPDoer
}

func NewAPIClient(doer HTTPDoer) *APIClient { return &APIClient{doer: doer} }

func (c *APIClient) GetUser(ctx context.Context, id int64) (*User, error) {
    req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "...", nil)
    resp, err := c.doer.Do(req)
    // ...
    return &User{}, err
}
```

> 生产环境传 *http.Client（它满足 Do 方法）；测试用 gomock 生成 HTTPDoer 的 mock


**时机/随机数**

```
type Clock interface{ Now() time.Time }
type realClock struct{}
func (realClock) Now() time.Time { return time.Now() }

type RNG interface{ Int63() int64 }
type realRNG struct{}
func (realRNG) Int63() int64 { return rand.Int63() }

// 注入
type Service struct{ clk Clock; rng RNG }
```

## Mock 使用

### mock 对象的结构

假设有接口：

```
type UserStore interface {
    GetUser(id int) (*User, error)
    SaveUser(u *User) error
}
```

mockgen 会生成：

```
type MockUserStore struct {
    ctrl     *gomock.Controller
    recorder *MockUserStoreMockRecorder
}

type MockUserStoreMockRecorder struct {
    mock *MockUserStore
}
```

- MockUserStore：实际在测试里被调用的 mock 对象
- MockUserStoreMockRecorder：用于配置“期望调用”的辅助对象，通过 .EXPECT() 访问

### 调用方法总览

| 方法 / 调用                                 | 作用                        |
|-----------------------------------------|---------------------------|
| `gomock.NewController(t)`               | 创建控制器，管理所有 mock 对象的生命周期   |
| `defer ctrl.Finish()`                   | 测试结束时自动检查是否所有期望被满足        |
| `mock.EXPECT()`                         | 进入录制模式，设置期望调用             |
| `.Return(...)`                          | 设置函数返回值                   |
| `.Times(n)`                             | 限制调用次数（必须恰好 n 次）          |
| `.AnyTimes()`                           | 允许被调用任意次数                 |
| `.MinTimes(n)` / `.MaxTimes(n)`         | 设置最少/最多调用次数               |
| `.Do(func(...){...})`                   | 调用时执行自定义函数逻辑              |
| `.After(call)`                          | 设置调用顺序依赖（必须在某个 call 之后发生） |
| `.InOrder(calls...)`                    | 声明多个期望的顺序                 |
| `gomock.Eq(x)` / `gomock.Any()`         | 参数匹配器（严格匹配 / 任意值）         |
| `gomock.Not(x)`                         | 参数不等于                     |
| `gomock.Nil()`                          | 参数为 nil                   |
| `gomock.AssignableToTypeOf(x)`          | 参数类型匹配                    |
| `.SetArg(i, value)`                     | 修改第 i 个入参（常用于指针或引用）       |
| `.DoAndReturn(func(...){ return ... })` | 用函数返回值动态决定结果              |

### 基础用法

#### 设置返回值

```
mockStore.EXPECT().
    GetUser(1).
    Return(&User{ID: 1, Name: "Alice"}, nil)
```

调用时：

```
u, _ := mockStore.GetUser(1)
fmt.Println(u.Name) // Alice
```

#### 设置调用次数

```
mockStore.EXPECT().
    SaveUser(gomock.Any()).
    Return(nil).
    Times(2)
```

> 只能被调用两次，否则测试失败

```
mockStore.EXPECT().SaveUser(gomock.Any()).AnyTimes()
```

> 可被调用任意次

#### 允许任何参数

```
mockStore.EXPECT().
    GetUser(gomock.Any()).
    Return(&User{Name: "Default"}, nil)
```

#### 多次不同返回

```
mockStore.EXPECT().
    GetUser(gomock.Any()).
    Return(&User{Name: "Alice"}, nil).
    Times(1)
mockStore.EXPECT().
    GetUser(gomock.Any()).
    Return(&User{Name: "Bob"}, nil).
    Times(1)
```

#### 自定义执行逻辑 .Do()

Mock 方法被调用时，要执行的操作，忽略返回值

```
mockStore.EXPECT().
    SaveUser(gomock.Any()).
    Do(func(u *User) {
        fmt.Println("Saving user:", u.Name)
    }).
    Return(nil)
```

#### 动态返回 .DoAndReturn()

可以动态地控制返回值

```
mockStore.EXPECT().
    GetUser(gomock.Any()).
    DoAndReturn(func(id int) (*User, error) {
        if id == 1 {
            return &User{Name: "Admin"}, nil
        }
        return nil, errors.New("not found")
    })
```

#### 指定调用顺序

```
first := mockStore.EXPECT().GetUser(1).Return(&User{Name: "A"}, nil)
second := mockStore.EXPECT().SaveUser(gomock.Any()).Return(nil).After(first)
```

或

```
gomock.InOrder(
    mockStore.EXPECT().GetUser(1).Return(&User{Name: "A"}, nil),
    mockStore.EXPECT().SaveUser(gomock.Any()).Return(nil),
)
```

#### 检查未被调用或超出调用

- 如果期望的 .Times(n) 没有满足 -> 测试失败
- 如果调用了未期望的方法 -> 测试失败

`ctrl.Finish()` 会自动检查所有这些。

#### 参数匹配器示例

| 匹配器                            | 示例                                              | 含义      |
|--------------------------------|-------------------------------------------------|---------|
| `gomock.Any()`                 | `.GetUser(gomock.Any())`                        | 任意参数    |
| `gomock.Eq(x)`                 | `.GetUser(gomock.Eq(10))`                       | 参数等于 10 |
| `gomock.Not(x)`                | `.GetUser(gomock.Not(5))`                       | 参数不等于 5 |
| `gomock.Nil()`                 | `.SaveUser(gomock.Nil())`                       | 参数为 nil |
| `gomock.AssignableToTypeOf(x)` | `.SaveUser(gomock.AssignableToTypeOf(&User{}))` | 参数类型匹配  |
| `gomock.Len(n)`                | `.SaveUser(gomock.Len(3))`                      | 参数长度为 3 |

### 使用流程

1. 创建控制器

    ```
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    ```

2. 创建 mock 对象

    ```
    mockStore := NewMockUserStore(ctrl)
    ```

3. 设置期望

    ```
    mockStore.EXPECT().GetUser(1).Return(&User{Name: "Alice"}, nil)
    ```

4. 调用被测逻辑：

    ```
    service := UserService{store: mockStore}
    name, _ := service.GetName(1)
    ```

5. 自动验证期望：

   如果期望没满足，或调用次数不符 -> 测试失败

- `.EXPECT()` 进入“期望设置模式”
- `.Return()` / `.Do()` / `.Times()` 等控制行为
- 匹配器（`gomock.Any()` 等）灵活匹配参数
- `gomock.InOrder()` 控制调用顺序
- 测试结束时自动验证是否符合预期

## 打桩(stubs)

```
m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
```

Get() 的参数为 Tom，则返回 error，这称之为打桩(stub)，有明确的参数和返回值是最简单打桩方式。
除此之外，检测调用次数、调用顺序，动态设置返回值等方式也经常使用。

### 参数(Eq, Any, Not, Nil)

```
m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
m.EXPECT().Get(gomock.Any()).Return(630, nil)
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil) 
m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil")) 
```

- Eq(value) 表示与 value 等价的值
- Any() 可以用来表示任意的入参
- Not(value) 用来表示非 value 以外的值
- Nil() 表示 None 值

### 返回值(Return, DoAndReturn)

```
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
m.EXPECT().Get(gomock.Any()).Do(func(key string) {
    t.Log(key)
})
m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
    if key == "Sam" {
        return 630, nil
    }
    return 0, errors.New("not exist")
})
```

- `Return`：返回确定的值
- `Do`：Mock 方法被调用时，要执行的操作，忽略返回值
- `DoAndReturn`：可以动态地控制返回值

### 调用次数(Times)

```
unc TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")
}
```

- Times() 断言 Mock 方法被调用的次数
- MaxTimes() 最大次数
- MinTimes() 最小次数
- AnyTimes() 任意次数（包括 0 次）

### 调用顺序(InOrder)

```
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}
```

