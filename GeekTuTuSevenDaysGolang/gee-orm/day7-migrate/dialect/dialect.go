package dialect

import "reflect"

/*
不同数据库字段类型和命令可能各不相同，使用 dialect 包进行解耦
*/

// 存储不同数据库 dialect
var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string            // 将 Go 语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []any) // 返回某个表是否存在的 SQL 语句
}

// RegisterDialect 注册 Dialect 实例
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取 Dialect 实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
