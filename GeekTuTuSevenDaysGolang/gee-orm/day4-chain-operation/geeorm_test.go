package geeorm

import "testing"
import _ "modernc.org/sqlite"

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite", "gee.db")
	if err != nil {
		t.Fatal("failed to connect db", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
