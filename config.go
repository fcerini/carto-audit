package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Configuracion: se carga desde config.json
type Config struct {
	SQL string //postgresql://user:pass@Host:port/Base
}

// load config
func (c *Config) load() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &c)
	if err2 != nil {
		log.Fatal(err2)
	}

}
