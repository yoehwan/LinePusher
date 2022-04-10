package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var config *Config
	//var auth *string

	configPath := flag.String("config", "./config.yaml", "Config file path")
	flag.Parse()

	if *configPath == "" {
		//auth = flag.String("auth", "", "Channel Access Token")

	} else {
		var err error
		config, err = LoadFromPath(*configPath)
		if err != nil {
			log.Println(err)
		}
	}
	pushMessage(config)

}

const api = "https://api.line.me/v2/bot/message/push"

func pushMessage(c *Config) error {
	res, err := http.PostForm(api, c.PostBody())
	if err != nil {
		return err
	}
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+c.Authorization)
	return nil
}
