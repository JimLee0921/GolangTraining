package session

import (
	"database/sql"
	"day1-database-sql/clause"
	"day1-database-sql/dialect"
	"day1-database-sql/mylog"
	"day1-database-sql/schema"
	"strings"
)

// 实现与数据库的交互

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect // sql 方言
	refTable *schema.Schema  // 映射字段
	sql      strings.Builder // 累计将要执行的 sql 语句字符串
	sqlVars  []any           // 保存 SQL 语句中的占位符变量（如 ?、$1 等）对应的值
	clause   clause.Clause
	tx       *sql.Tx // 支持事务
}

// CommonDB 接口，统一操作 DB / TX，把 sql.DB 和 sql.Tx 共同拥有的方法抽象成一个接口
type CommonDB interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// 编译期静态检查
var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

// DB 如果 tx 存在则返回事务 tx，否则返回 db，兼容之前代码
func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

// New 创建 Session 实例
func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

// Clear 清空 Session 的 SQL 缓存
func (s *Session) clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

// Raw 把用户传入的 SQL 和变量值记录下来，返回 session 本身（这样可以链式调用）
func (s *Session) Raw(sql string, values ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

/*
封装 Exec(), Query() 和 QueryRaw()
可以做到统一打印日志（包括 执行的SQL 语句和错误日志）
并且执行完毕后自动调用 clear 方法做到会话复用
*/

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.clear()

	mylog.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		mylog.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.clear()
	mylog.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) Query() (rows *sql.Rows, err error) {
	defer s.clear()
	mylog.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		mylog.Error(err)
	}
	return
}
