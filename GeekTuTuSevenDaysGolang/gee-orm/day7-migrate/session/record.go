package session

import (
	"day1-database-sql/clause"
	"errors"
	"reflect"
)

/*
实现记录增删改查相关
*/

func (s *Session) Insert(values ...any) (int64, error) {
	recordValues := make([]any, 0)
	for _, value := range values {
		s.CallMethod(BeforeInsert, value)
		table := s.Model(value).RefTable()
		// 多次调用 clause.Set() 构造好每一个子句
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}
	s.clause.Set(clause.VALUES, recordValues...)
	// 生成最终的 SQL 语句
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	// 执行 SQL
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterInsert, nil)
	return result.RowsAffected()
}

func (s *Session) Find(values any) error {
	s.CallMethod(BeforeQuery, nil)
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)

	rows, err := s.Raw(sql, vars...).Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []any
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		s.CallMethod(AfterQuery, dest.Addr().Interface())
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

/*
1. destSlice.Type().Elem() 获取切片的单个元素的类型 destType，使用 reflect.New() 方法创建一个 destType 的实例，作为 Model() 的入参，映射出表结构 RefTable()
2. 根据表结构，使用 clause 构造出 SELECT 语句，查询到所有符合条件的记录 rows
3. 遍历每一行记录，利用反射创建 destType 的实例 dest，将 dest 的所有字段平铺开，构造切片 values
4. 调用 rows.Scan() 将该行记录每一列的值依次赋值给 values 中的每一个字段
5. 将 dest 添加到切片 destSlice 中。循环直到所有的记录都添加到切片 destSlice 中
*/

// Update 支持两种参数：map[string]any 或 kv 列表：["name": "jim", "age": 12, ...]
func (s *Session) Update(kv ...any) (int64, error) {
	s.CallMethod(BeforeUpdate, nil)
	m, ok := kv[0].(map[string]any)
	if !ok {
		m = make(map[string]any)
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterUpdate, nil)
	return result.RowsAffected()
}

func (s *Session) Delete() (int64, error) {
	s.CallMethod(BeforeDelete, nil)
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterDelete, nil)
	return result.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()
	var temp int64
	if err := row.Scan(&temp); err != nil {
		return 0, err
	}
	return temp, nil
}

// First 只返回一条记录
// 根据传入的类型，利用反射构造切片，调用 Limit(1) 限制返回的行数，调用 Find 方法获取到查询结果。
func (s *Session) First(value any) error {
	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}

// 下面三个不执行 sql 只是往 clause 中保存对应的子句内容

// Limit 保存 limit 子句
func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

// Where 保存 where 条件子句
func (s *Session) Where(desc string, args ...any) *Session {
	var vars []any
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s
}

// OrderBy 保存 orderby 排序规则子句
func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}
