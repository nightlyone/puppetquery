package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/nightlyone/puppetquery"
)

func main() {
	var typ string
	flag.StringVar(&typ, "t", "", "type of resource")
	flag.Parse()
	tags := flag.Args()
	log.Print("INFO: Start query")
	cr, err := puppetquery.CollectResources(typ, tags...)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	} else {
		log.Printf("INFO: End query (received %d resources)\n", len(cr))
	}
	b, err := json.MarshalIndent(cr, "", "  ")
	if err != nil {
		log.Fatalln("ERROR: cannot marshal result into json: ", err)
	}

	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Fatalln("ERROR: cannot display result: ", err)
	}
}
