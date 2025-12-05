package clause

import (
	"strings"
)

/*
实现结构体 Clause 拼接各个独立的子句
*/

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]any
}

// Type 定义子句常量
type Type int

// 新增其它子句
const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

// Set 根据传入的 Type 调用对应的 generator，生成子句对应的 sql 语句
func (c *Clause) Set(name Type, vars ...any) {
	// 延迟初始化
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]any)
	}

	sql, vars := generators[name](vars...)

	c.sql[name] = sql
	c.sqlVars[name] = vars
}

// Build 根据传入 Type 的顺序构造出最终完整的 sql 语句
func (c *Clause) Build(orders ...Type) (string, []any) {
	var sqls []string
	var vars []any

	for _, order := range orders {
		if sql, ok := c.sql[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
