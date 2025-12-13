package parse

import "strconv"

func Parse(s string) (int, error) {
	return strconv.Atoi(s)
}
