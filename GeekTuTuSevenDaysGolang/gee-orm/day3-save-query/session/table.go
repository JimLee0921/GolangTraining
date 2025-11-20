package session

import (
	"day1-database-sql/mylog"
	"day1-database-sql/schema"
	"fmt"
	"reflect"
	"strings"
)

/*
数据库操作表相关
*/

// Model 给 refTable 赋值
func (s *Session) Model(value any) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 返回 refTable 的值
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		mylog.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable 创建表
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, filed := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", filed.Name, filed.Type, filed.Tag))

	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

// DropTable 删除表
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable 查询表是否存在
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
