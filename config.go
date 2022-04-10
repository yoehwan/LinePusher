package main

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
)

type Config struct {
	Authorization string
	UserID        string
	Sender        *Sender
	Message       *Message
}

type Sender struct {
	Name    string
	IconUrl string
}

type Message struct {
	Data map[string]interface{}
}

func (c *Config) PostBody() ([]byte, error) {
	body := map[string]interface{}{
		"to": c.UserID,
		"messages": []map[string]interface{}{
			{
				"type": "text",
				"text": "Hello",
			},
		},
	}
	res, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func LoadFromPath(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return readYaml(&data)
}

func readYaml(data *[]byte) (*Config, error) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(*data))
	if err != nil {
		return nil, err
	}
	return &Config{
		Authorization: viper.GetString("authorization"),
		UserID:        viper.GetString("user_id"),
		Sender: &Sender{
			Name:    viper.GetString("sender.name"),
			IconUrl: viper.GetString("sender.icon_url"),
		},
		Message: &Message{Data: viper.GetStringMap("message")},
	}, nil
}
