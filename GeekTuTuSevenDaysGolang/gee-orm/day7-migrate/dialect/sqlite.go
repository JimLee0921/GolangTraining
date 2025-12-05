package dialect

import (
	"fmt"
	"reflect"
	"time"

	_ "modernc.org/sqlite"
)

/*
针对 sqlite 数据库的 dialect
*/

type sqlite struct {
}

// 编译检查
var _ Dialect = (*sqlite)(nil)

// init 函数，包在第一次加载时会执行 init，将 sqlite 的 dialect 自动注册到全局
func init() {
	RegisterDialect("sqlite", &sqlite{})
}

func (s *sqlite) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (s *sqlite) TableExistSQL(tableName string) (string, []any) {
	args := []any{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
