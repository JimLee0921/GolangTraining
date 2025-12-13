package demo

import (
	"encoding/json"
	"testing"
)

func FuzzJSONUser(f *testing.F) {
	f.Add(`{"ID":1,"Name":"Tom","Email":"a@b.com"}`)
	f.Add(`{}`)
	f.Add(`{"ID":-1,"Name":"","Email":""}`)

	f.Fuzz(func(t *testing.T, input string) {
		var u User
		if err := json.Unmarshal([]byte(input), &u); err != nil {
			return
		}

		data, err := json.Marshal(u)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}

		var u2 User
		if err := json.Unmarshal(data, &u2); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
	})
}
