package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Env string
	Ver string
}

var loadAppConfig = sync.OnceValue(func() *Config {
	fmt.Println("init config")
	return &Config{Env: "Prod", Ver: "1.0.0"}
})

func main() {
	cfg1 := loadAppConfig()
	cfg2 := loadAppConfig()
	fmt.Println(cfg1 == cfg2) // true，指向同一个对象
}
