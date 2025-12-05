package session

import (
	"day1-database-sql/mylog"
	"testing"
)

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	mylog.Info("before insert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	mylog.Info("after query", account)
	account.Password = "**********"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{
		ID:       1,
		Password: "15994958",
	}, &Account{
		ID:       2,
		Password: "wdjawjj1",
	})

	u := &Account{}
	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "**********" {
		t.Fatal("failed to call hooks after query, got", u)
	}
}
