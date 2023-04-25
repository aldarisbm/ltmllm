package main

import (
	"fmt"
	"github.com/aldarisbm/ltmllm/config"
)

func main() {
	conf := config.NewConfig()
	fmt.Printf("%v\n", conf)
}
