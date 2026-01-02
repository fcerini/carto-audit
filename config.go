package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Configuracion: se carga desde config.json
type Config struct {
	DbUser     string // Usuario de PostgreSQL
	DbPassword string // Password de PostgreSQL
	DbHost     string // Host de PostgreSQL
	DbPort     string // Puerto de PostgreSQL
	DbName     string // Nombre de PostgreSQL

	TablaCalles string // Tabla de vectores de calles para geocodificacion
	SQLCartas   string // Consulta SQL para obtener las cartas a procesar
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
