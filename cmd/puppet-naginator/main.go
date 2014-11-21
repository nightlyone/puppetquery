package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/nightlyone/puppetquery"
)

func CollectNagiosResource(typ string, resp chan<- *bytes.Buffer, tags ...string) {
	log.Println("INFO: Start query for", typ)
	cr, err := puppetquery.CollectResources(nagiosPrefix+typ, tags...)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	} else {
		log.Printf("INFO: End query for %s (received %d resources)\n", typ, len(cr))
	}
	b := new(bytes.Buffer)
	log.Print("INFO: generating resources for", typ)
	err = generate(b, time.Now(), cr)
	if err != nil {
		log.Print("ERROR: generating resources for", typ)
		resp <- nil
	} else {
		log.Print("INFO: done generating resources for", typ)
		resp <- b
	}
}

var nagiosTypes = strings.Fields(`command contact contactgroup host hostdependency hostescalation hostextinfo
    hostgroup service servicedependency serviceescalation serviceextinfo servicegroup timeperiod`)

func main() {
	var typ string
	flag.StringVar(&typ, "t", "", "type of nagios resource")
	flag.Parse()
	tags := flag.Args()
	types := []string{typ}
	if typ == "" {
		types = nagiosTypes
		log.Print("INFO: generating all resources in a single file")
	} else {
		sort.Strings(nagiosTypes)
		if sort.SearchStrings(nagiosTypes, typ) < 0 {
			log.Fatalln("ERROR: invalid nagios type: ", typ)
		}
	}
	buffers := make(chan *bytes.Buffer, len(types))
	for _, t := range types {
		go CollectNagiosResource(t, buffers, tags...)
	}
	for _, _ = range types {
		b := <-buffers
		if b == nil {
			log.Fatalln()
			return
		}
		_, err := io.Copy(os.Stdout, b)
		if err != nil {
			log.Fatalln("ERROR: cannot display result: ", err)
		} else {
			log.Print("INFO: written file")
		}
	}
}
