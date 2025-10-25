package main

import (
	"fmt"
	"path"
)

func main() {
	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("a/b/c/d/e/")
	split("a/b/c/d/e")
	split("static/myfile.css")
	split("myfile.css")
	split("")
	/*
		path.Split("static/myfile.css") = dir: "static/", file: "myfile.css"
		path.Split("myfile.css") = dir: "", file: "myfile.css"
		path.Split("") = dir: "", file: ""
	*/
}
