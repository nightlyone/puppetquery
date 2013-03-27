package main

import (
	"log"
	"flag"
	"github.com/nightlyone/puppetquery"
	"fmt"
	"strings"
)

func ParseEqualFact(arg string) puppetquery.QueryString {
	f := strings.SplitN(arg, "=", 2)
	fact, value := f[0], f[1]
	return puppetquery.FactCompare(fact, "=", value)
}

func main() {
	flag.Parse()
	q := puppetquery.QueryString{}
	for i := 0; i < flag.NArg(); i++ {
		q = puppetquery.And(q, ParseEqualFact(flag.Arg(i)))
	}
	log.Println("DEBUG: query: ", q.ToJson())
	n, err := puppetquery.QueryNodes(q)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	}
	for _, node := range n {
		fmt.Println(node)
	}
}
