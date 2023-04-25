package main

import (
	"LTMLLM/config"
	"fmt"
)

func main() {
	conf := config.NewConfig()
	fmt.Printf("%v\n", conf)
}
