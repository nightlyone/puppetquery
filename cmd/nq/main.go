package main

import (
	"flag"
	"fmt"
	"github.com/nightlyone/puppetquery"
	"log"
	"sort"
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

func simpleAndQuery() {
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

type pair struct {
	key, value string
}

type byKey []pair

func (p byKey) Len() int           { return len(p) }
func (p byKey) Less(i, j int) bool { return p[i].key < p[j].key }
func (p byKey) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func factsOfNodeQuery(node string) {
	facts, err := puppetquery.QueryFacts(node)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	}
	linearFacts := []pair{}
	for k, v := range facts {
		linearFacts = append(linearFacts, pair{key: k, value: v})
	}
	sort.Sort(byKey(linearFacts))
	for _, v := range linearFacts {
		fmt.Printf("%v=%v\n", v.key, v.value)
	}
}

func main() {
	var listFacts string
	flag.StringVar(&listFacts, "l", "", "list all facts of a node")
	flag.Parse()

	if len(listFacts) > 0 {
		factsOfNodeQuery(listFacts)
	} else {
		simpleAndQuery()
	}
}
