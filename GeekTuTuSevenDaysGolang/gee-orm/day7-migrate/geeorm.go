package geeorm

import (
	"database/sql"
	"day1-database-sql/dialect"
	"day1-database-sql/mylog"
	"day1-database-sql/session"
	"fmt"
	"strings"
)

/*
geeorm 入口文件
*/

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		mylog.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		mylog.Error(err)
		return
	}

	// 确保 dialect 存在
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		mylog.Errorf("dialect %s not found", driver)
	}

	e = &Engine{db: db, dialect: dial}
	mylog.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		mylog.Error("Failed to close database", err)
	}
	mylog.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

type TxFunc func(*session.Session) (any, error)

func (engine *Engine) Transaction(f TxFunc) (result any, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p) // 回滚后抛出异常
		} else if err != nil {
			_ = s.Rollback() // err 不为空回滚
		} else {
			err = s.Commit() // err 为空，进行事务提交，如果提交返回错误更新 err
		}
	}()
	return f(s)
}

// migrate 表修改

// difference 对比字段
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

// Migrate 迁移表
func (engine *Engine) Migrate(value any) error {
	_, err := engine.Transaction(func(s *session.Session) (result any, err error) {
		// 1. 表不存在则直接建表
		if !s.Model(value).HasTable() {
			mylog.Infof("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}

		table := s.RefTable()

		// 2. 查询一行，拿当前表字段
		rows, err := s.
			Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).
			Query()
		if err != nil {
			return nil, err
		}

		columns, err := rows.Columns()
		// ⚠️ 一定要在 DDL 之前就关闭 rows，释放对 User 表的锁
		rows.Close()
		if err != nil {
			return nil, err
		}

		addCols := difference(table.FieldNames, columns) // 需要新增的列
		delCols := difference(columns, table.FieldNames) // 需要删除的列
		mylog.Infof("added cols %v, delete cols %v", addCols, delCols)

		// 3. 逐个新增缺失字段
		for _, col := range addCols {
			f := table.GetField(col)
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table.Name, f.Name, f.Type)
			if _, err = s.Raw(sqlStr).Exec(); err != nil {
				return nil, err
			}
		}

		// 4. 没有要删的列就结束
		if len(delCols) == 0 {
			return nil, nil
		}

		// 5. 需要删列：临时表方案
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")

		// (1) 创建临时表，只保留需要的字段
		if _, err = s.Raw(
			fmt.Sprintf("CREATE TABLE %s AS SELECT %s FROM %s;", tmp, fieldStr, table.Name),
		).Exec(); err != nil {
			return nil, err
		}

		// (2) 删除旧表
		if _, err = s.Raw(
			fmt.Sprintf("DROP TABLE %s;", table.Name),
		).Exec(); err != nil {
			return nil, err
		}

		// (3) 临时表改名
		if _, err = s.Raw(
			fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmp, table.Name),
		).Exec(); err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}
