package main

import (
	"flag"
	"fmt"
)

func main() {
	configPath := flag.String("config", "", "Config file path")
	auth := flag.String("auth", "", "Channel Access Token")

	flag.Parse()

	fmt.Println(auth)

	fmt.Println(configPath)

}
