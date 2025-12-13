package main

import "testing"

func FuzzEncodeDecode(f *testing.F) {
	f.Add("hello")
	f.Add("")
	f.Add("你好")
	f.Add("🙂")

	f.Fuzz(func(t *testing.T, s string) {
		out := Decode(Encode(s))
		if out != s {
			t.Fatalf("encode/decode failed %q -> %q", s, out)
		}
	})
}
