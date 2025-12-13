package main

/*
Encode/Decode 经常成对出现
Decode(Encode(x) == x) 不变量
*/

func Encode(s string) []byte {
	return []byte(s)
}

func Decode(b []byte) string {
	return string(b)
}
