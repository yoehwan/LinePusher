package main

import (
	"bytes"
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
	body, err := c.PostBody()
	if err != nil {
		return err
	}
	log.Println(string(body))
	req, err := http.NewRequest("POST", api, bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.Authorization)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)

		return err
	}
	log.Println(string(res))
	return nil
}
