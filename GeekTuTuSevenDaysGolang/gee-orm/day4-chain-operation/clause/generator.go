package clause

import (
	"fmt"
	"strings"
)

/*
把 SQL 的某个子句（Clause）封装成一个可以被调用的函数
*/

// 生成 SQL 字符串 + 参数列表
type generator func(values ...any) (string, []any)

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
	generators[UPDATE] = _update
	generators[DELETE] = _delete
	generators[COUNT] = _count
}

// 拼接参数
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

// _insert 接收表名和字段列表 ["User", []string{"Name", "Age"}]
func _insert(values ...any) (string, []any) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []any{}
}

// _values []interface{}{"Tom", 18}, []interface{}{"Sam", 25}
func _values(values ...any) (string, []any) {
	var bindStr string      // 存储占位符
	var sql strings.Builder // 构造 VALUES 子句
	var vars []any          // 存储所有参数
	sql.WriteString("VALUES ")
	for i, value := range values {
		// values 是 ...any 类型，每个 value 实际上是 []any 需要作类型断言填回切片
		v := value.([]any)
		// 第一次循环生成占位符
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		// 循环拼接
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	// 返回子句和参数
	return sql.String(), vars
}

// _select(tableName, []string{"Name", "Age"})
func _select(values ...any) (string, []any) {
	// SELECT $fields FROM $tableName
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []any{}
}

// _limit(10)
func _limit(values ...any) (string, []any) {
	// LIMIT %num
	return "LIMIT ?", values
}

// _where("Name = ?", "Tom")
func _where(values ...any) (string, []any) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

// _orderBy("Age DESC")
func _orderBy(values ...any) (string, []any) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []any{}
}

// _update("User", map[string]any{"age":30,"name":"jim",})
func _update(values ...any) (string, []any) {
	tableName := values[0]
	m := values[1].(map[string]any)
	var keys []string
	var vars []any
	for k, v := range m {
		keys = append(keys, k+" = ?")
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), vars
}

// _delete("User")
func _delete(values ...any) (string, []any) {
	return fmt.Sprintf("DELETE From %s", values[0]), []any{}
}

// _count("User") 单纯查询条数，借用 _select 结构
func _count(values ...any) (string, []any) {
	return _select(values[0], []string{"count(*)"})
}
