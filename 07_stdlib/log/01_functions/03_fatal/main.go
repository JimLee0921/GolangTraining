package main

import "log"

// fatal 输出后立即退出程序，不执行 defer
func main() {
	//log.Fatal("fatal error, exiting")
	//log.Fatalln("fatal:", "config missing")
	log.Fatalf("cannot connect to %s", "database")
}
