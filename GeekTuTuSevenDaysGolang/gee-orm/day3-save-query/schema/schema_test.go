package schema

import (
	"day1-database-sql/dialect"
	"reflect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.FieldNames) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	// 构造 User 实例
	u := &User{
		Name: "JimLee",
		Age:  20,
	}

	// 调用 RecordValue
	values := schema.RecordValues(u)
	t.Logf("%v", values)
	want := []any{"JimLee", 20}
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("recordValues failed, got %v,  expected: %v", values, want)
	}
}
