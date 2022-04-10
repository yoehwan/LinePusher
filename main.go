package main

import (
	"flag"
	"io/ioutil"
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
	err := pushMessage(config)
	if err != nil {
		log.Println(err)
	}

}

const api = "https://api.line.me/v2/bot/message/push"

func pushMessage(c *Config) error {
	client := http.Client{}
	req, err := http.NewRequest("POST", api, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.Authorization)
	/// add body
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(bytes))
	return nil
}
