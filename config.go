package main

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
)

type Config struct {
	Authorization string                   `mapstructure:"authorization"`
	Sender        *Sender                  `mapstructure:"sender"`
	To            string                   `mapstructure:"to"`
	Messages      []map[string]interface{} `mapstructure:"messages"`
}

type Sender struct {
	Name    string `json:"name",mapstructure:"name"`
	IconUrl string `json:"iconUrl",mapstructure:"iconUrl"`
}

func (c *Config) PostBody() ([]byte, error) {
	var msg []map[string]interface{}
	for _, item := range c.Messages {
		item["sender"] = c.Sender
		msg = append(msg, item)
	}

	body := map[string]interface{}{
		"to":       c.To,
		"messages": msg,
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
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
