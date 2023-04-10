package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Environment        string `json:"environment"`
	Port               int    `json:"port"`
	PostgresConnection string `json:"postgresConnection"`
}

func ReadCfg() *Config {
	filePath := "./.config/local.json"
	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var cfg Config
	err = json.NewDecoder(bytes.NewBuffer(fileByte)).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return &cfg
}
