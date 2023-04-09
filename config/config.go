package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadCfg() *Config {
	filePath := "./.config/config.json"
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
