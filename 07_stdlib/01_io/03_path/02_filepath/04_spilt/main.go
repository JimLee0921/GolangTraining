package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Split("main.go"))      // ("", "main.go")
	fmt.Println(filepath.Split("/usr/local/"))  // ("/usr/local/", "")
	fmt.Println(filepath.Split("/usr/local/a")) // ("/usr/local/", "a")

	// Linux
	paths := "/usr/bin:/bin:/usr/local/bin"
	fmt.Println(filepath.SplitList(paths)) // [/usr/bin /bin /usr/local/bin]

	// Windows
	paths = `C:\Windows;C:\Program Files;D:\Go\bin`
	fmt.Println(filepath.SplitList(paths)) // [C:\Windows C:\Program Files D:\Go\bin]

}
