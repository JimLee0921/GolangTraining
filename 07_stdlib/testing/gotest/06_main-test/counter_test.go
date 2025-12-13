package counter

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	testStorePath string
	fileStore     Store
)

// ---- 普通单元测试 ----
func TestCounter_AddAndInc(t *testing.T) {
	c := NewCounter(1)
	t.Run("inc", func(t *testing.T) {
		if got := c.Inc(); got != 2 {
			t.Fatalf("Inc() = %d; want %d", got, 2)
		}
	})
	t.Run("add", func(t *testing.T) {
		if got := c.Add(41); got != 43 {
			t.Fatalf("Add() = %d; want %d", got, 43)
		}
	})
}

// ---- 表驱动子测试 ----
func TestCounter_Add_Cases(t *testing.T) {
	cases := []struct {
		name  string
		init  int
		delta int
		want  int
	}{
		{"positive", 1, 2, 3},
		{"zero", 1, 0, 1},
		{"negitive", 10, -3, 7},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			c := NewCounter(tc.init)
			if got := c.Add(tc.delta); got != tc.want {
				t.Fatalf("Add(%d)=%d, want %d", tc.delta, got, tc.want)
			}
		})
	}
}

// ---- 依赖外部资源的测试：需要 store 初始化 ----
func TestFileStore_LoadFromTestMainSeed(t *testing.T) {
	v, err := fileStore.Load()
	if err != nil {
		t.Fatalf("Load error: %v", err)
	}
	if v != 10 {
		t.Fatalf("Load()=%d, want 10 (seeded in TestMain)", v)
	}
}

func TestFileStore_SaveAndLoad(t *testing.T) {
	if err := fileStore.Save(99); err != nil {
		t.Fatalf("Save() error: %v", err)
	}
	v, err := fileStore.Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if v != 99 {
		t.Fatalf("Load()=%d, want 99", v)
	}
}

// ---- Example：展示在文档中；带 Output 才会执行校验 ----

func ExampleCounter_Inc() {
	c := NewCounter(41)
	fmt.Println(c.Inc())
	// Output:
	// 42
}

// ---- Benchmark：也会受 TestMain 影响（比如依赖 store/环境） ----

func BenchmarkCounter_Inc(b *testing.B) {
	c := NewCounter(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Inc()
	}
}

func BenchmarkFileStore_Save(b *testing.B) {
	// 举例：某些 benchmark 依赖 OS/文件系统；不满足条件就 skip
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		b.Skip("no filesystem available")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fileStore.Save(i)
	}
}

// ---- 可选：Fuzz（需要 go test -fuzz 才会跑） ----

func FuzzFileStore_SaveLoad(f *testing.F) {
	// seed
	f.Add(0)
	f.Add(42)
	f.Add(-7)

	f.Fuzz(func(t *testing.T, v int) {
		if err := fileStore.Save(v); err != nil {
			t.Fatalf("Save(%d) error: %v", v, err)
		}
		got, err := fileStore.Load()
		if err != nil {
			t.Fatalf("Load() error: %v", err)
		}
		if got != v {
			t.Fatalf("Load()=%d, want %d", got, v)
		}
	})
}

// TestMain 测试入口（整个 package 只会被调用一次）
func TestMain(m *testing.M) {
	// ==== setup 阶段，整个包只执行一次 ====
	tmpDir, err := os.MkdirTemp("", "counter-test-*")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "setup failed", err)
		os.Exit(1)
	}
	testStorePath = filepath.Join(tmpDir, "counter.txt")
	fileStore = FileStore{Path: testStorePath}

	// 给测试问了中写入一个已知的初始状态
	_ = fileStore.Save(10)

	// ==== run，必须要有才会调用测试 ====
	code := m.Run()

	// ==== teardown 资源清理阶段 ====
	_ = os.RemoveAll(tmpDir)

	// ==== 返回 Run 运行结束的 code 来观察哪些测试失败  =====
	os.Exit(code)
}
