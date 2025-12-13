package demo

// User 业务结构体
// json.Unmarshal(json.Marshal(x)) == x 应该是不变量
type User struct {
	ID    int
	Name  string
	Email string
}
