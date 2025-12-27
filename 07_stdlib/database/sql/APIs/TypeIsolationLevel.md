# sql.IsolationLevel

`sql` 包中定义了 `IsolationLevel` 隔离级别枚举类型常量：

```
const (
	LevelDefault IsolationLevel = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
	LevelSerializable
	LevelLinearizable
)
```

参考隔离级别：https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels

## 主要方法

就一个 String 方法，返回字符串表现形式

```
func (i IsolationLevel) String() string
```


