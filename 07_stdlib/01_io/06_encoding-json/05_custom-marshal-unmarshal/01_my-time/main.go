package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyTime time.Time

const layout = "2006-01-02 15:04:05"

// 格式化 MyTime 类型输出
func (t MyTime) String() string {
	return time.Time(t).Format(layout)

}

func (t MyTime) MarshalJSON() ([]byte, error) {
	// 将 time 转为格式化字符串，再使用 json.Marshal 序列化返回结果
	s := time.Time(t).Format(layout)
	return json.Marshal(s)
}

func (t *MyTime) UnmarshalJSON(b []byte) error {
	// 1. 允许 nil
	if string(b) == "null" {
		*t = MyTime(time.Time{})
		return nil
	}

	// 2. 尝试作为字符串解析
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		tt, err := time.ParseInLocation(layout, s, time.Local)
		if err != nil {
			return fmt.Errorf("MyTime: invalid time string %q", s)
		}
		*t = MyTime(tt)
		return nil
	}
	// 3. 尝试作为数字解析（秒/毫秒）
	var num float64
	if err := json.Unmarshal(b, &s); err == nil {
		// 判断是毫秒还是秒
		if num > 1e12 {
			*t = MyTime(time.UnixMilli(int64(num)))
		} else {
			*t = MyTime(time.Unix(int64(num), 0))
		}
		return nil
	}
	return fmt.Errorf("MyTime: unsupported format %s", string(b))
}

type Event struct {
	Name string `json:"name"`
	Time MyTime `json:"time"`
}

func main() {
	// 序列化测试
	e := Event{Name: "Launch", Time: MyTime(time.Date(2025, 10, 30, 15, 4, 5, 0, time.Local))}
	b, _ := json.Marshal(e)

	fmt.Println(string(b))

	// 反序列化测试
	//var e Event
	//json.Unmarshal([]byte(`{"name":"Test","time":"2025-11-01 12:00:00"}`), &e)
	//fmt.Printf("%T: %v\n", e.Time, e.Time)
}
