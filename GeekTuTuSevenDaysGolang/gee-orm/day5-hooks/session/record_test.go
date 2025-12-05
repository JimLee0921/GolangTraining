package session

import "testing"

var (
	u1 = &User{
		Name: "JimLee",
		Age:  12,
	}
	u2 = &User{
		Name: "JamesBond",
		Age:  18,
	}
	u3 = &User{
		Name: "ElonMusk",
		Age:  22,
	}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(u1, u2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(u3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create new record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
	for _, user := range users {
		t.Logf("%v", user)
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "JimLee").Update("Age", 666)
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)
	if affected != 1 || u.Age != 666 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("name = ?", "JimLee").Delete()
	count, _ := s.Count()
	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}
