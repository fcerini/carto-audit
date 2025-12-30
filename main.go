package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var gloAPP Config

// se asigna en build.sh
var version string

func main() {

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println(version)
		return
	}
	gloAPP.load()

	log.Printf("START carto-audit %v", version)

}
