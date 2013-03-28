package main

import (
	"flag"
	"fmt"
	"github.com/nightlyone/puppetquery"
	"log"
	"strings"
)

func ParseEqualFact(arg string) puppetquery.QueryString {
	f := strings.SplitN(arg, "=", 2)
	if len(f) != 2 {
		log.Fatalln("ERROR: invalid argument: ", arg)
	}
	fact, value := f[0], f[1]
	return puppetquery.FactCompare(fact, "=", value)
}

func main() {
	flag.Parse()
	q := puppetquery.ActiveNodes()
	for i := 0; i < flag.NArg(); i++ {
		q = puppetquery.And(q, ParseEqualFact(flag.Arg(i)))
	}
	n, err := puppetquery.QueryNodes(q)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	}
	for _, node := range n {
		fmt.Println(node)
	}
}
