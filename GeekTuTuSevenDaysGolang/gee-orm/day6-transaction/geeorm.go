package geeorm

import (
	"database/sql"
	"day1-database-sql/dialect"
	"day1-database-sql/mylog"
	"day1-database-sql/session"
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
