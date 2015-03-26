// The puppet-naginator command collects nagios_XXXX resources and writes them to stdout.
// Using the parameter -t TYPE only outputs nagios config defines of TYPE.
// The TYPE is the noun you write after "define" in your nagios config.
//
// Example:
// 	puppet-naginator -t host
// Output:
//	define host {
//		host_name myhost.example.org
//	}
package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"github.com/nightlyone/puppetquery"
)

func CollectNagiosResource(typ string, resp chan<- *bytes.Buffer, tags ...string) {
	log.Print("INFO: Start PuppetDB query for resources of type nagios_", typ)
	cr, err := puppetquery.CollectResources(nagiosPrefix+typ, tags...)
	if err != nil {
		log.Fatalln("ERROR: cannot query PuppetDB: ", err)
	} else {
		log.Printf("INFO: End query for %s (received %d resources)\n", typ, len(cr))
	}
	b := new(bytes.Buffer)
	err = generate(b, time.Now(), cr)
	if err != nil {
		log.Println("ERROR: generating resources for", typ)
		resp <- nil
	} else {
		log.Printf("INFO: done generating %d %s definitions\n", len(cr), typ)
		resp <- b
	}
}

func main() {
	var typ string
	flag.StringVar(&typ, "t", "", "type of nagios resource (the noun after 'define' in your nagios config)")
	flag.Parse()
	tags := flag.Args()
	types := []string{typ}
	if typ == "" {
		types = nagiosTypes
		log.Print("INFO: generating all resources in a single file")
	} else {
		sort.Strings(nagiosTypes)
		if i := sort.SearchStrings(nagiosTypes, typ); i < 0 || nagiosTypes[i] != typ {
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
		}
	}
}
