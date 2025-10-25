package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	fmt.Printf("Path separator: %q\n", string(filepath.Separator))
	fmt.Printf("List separator: %q\n", string(filepath.ListSeparator))

	/*
		Linux/Macos
			Path separator: '/'
			List separator: ':'

		Windows
			Path separator: '\'
			List separator: ';'
	*/
}
