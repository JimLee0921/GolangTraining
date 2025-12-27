# time.Location

`time.Location` 表示一个 时区规则，包括：

- 该时区与 UTC 的偏移量（例如 +08:00）
- 是否需要考虑夏令时（DST）
- 时区的名称（如 `Asia/Shanghai`、`America/New_York`）

Go 的时间 `time.Time` 结构中始终携带一个 Location，否则就无法正确表示本地时间的意义，位置信息用于在打印的时间值中提供时区，以及用于涉及可能跨越夏令时边界的时间间隔的计算

## 常见 Location

| 名称                                   | 说明                   |
|--------------------------------------|----------------------|
| `time.UTC`                           | 世界协调时（零时区）           |
| `time.Local`                         | 当前系统所在本地时区（依赖系统环境变量） |
| `time.LoadLocation("Asia/Shanghai")` | 明确指定某个时区（推荐）         |

但是在 Go 中建议使用 LoadLocation 而不是 Local，因为 Local 依赖服务器环境：

- Linux 用 /etc/localtime
- Windows 与系统区域设置有关
- Docker 容器里经常是 UTC 导致乱套

所以 跨系统/跨服务器 时，Local 不可靠

## 主要方法和函数

### time.FixedZone()

FixedZone 主要用于创建一个固定偏移量的时区：

- 不会有夏令时
- 偏移量永远不变
- 不依赖系统时区数据库

```
func FixedZone(name string, offset int) *Location
```

**参数**

- name：传入时区名称（展示使用）
- offset：相对于 UTC 的秒数偏移量

**返回值**

- `*time.Location`：永不返回 nil 或 error

**注意事项**

- 没有 DST 夏令时
- name 只是标签，并不参与计算
- 不要替代真实国家时区
- 主要用于 API 返回固定 UTC 偏移的时间或测试/Mock 使用

### time.LoadLocation

非常重要，从 `IANA Time Zone Database（tzdata）` 加载一个真实世界时区：

- 支持夏令时
- 支持历史规则
- 支持未来规则

```
func LoadLocation(name string) (*time.Location, error)
```

**参数**

- name：IANA 时区名，如 `Asia/Shanghai`

**返回值**

- `*Location`：成功加载的时区
- `error`：找不到指定时区名/数据无效

**注意事项**

- 依赖系统 `tzdata`
    - Linux：`/usr/share/zoneinfo`
    - masOS：系统自带
    - Windows：Go 内部携带了一份精简 `tzdata`（`time/tzdata` 包）
- IANA 时区名必须符合规范

### time.LoadLocationFromTZData

偏高级/底层用法，从二进制 `tzdata` 数据中手动加载时区

- 不依赖操作系统
- 完全自定义
- LoadLocation 底层能力，开发几乎用不到

```
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
```

**参数**

- `name`： 时区名称（标识用）
- `data`： `[]byte` | tzfile 格式的二进制数据

**返回值**

- `*Location`：解析后的时区
- `error`：数据非法/格式错误

### String()

返回该 Location 的名称字符串表示

```
func (l *time.Location) String() string
```

- 仅用于展示 / 调试
- 不保证唯一性
- 不参与时间计算逻辑