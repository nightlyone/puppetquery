package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/nightlyone/puppetquery"
)

func ParseEqualFact(arg string) puppetquery.QueryString {
	f := strings.SplitN(arg, "=", 2)
	if len(f) != 2 {
		log.Fatalln("ERROR: invalid argument: ", arg)
	}
	fact, value := f[0], f[1]
	return puppetquery.FactCompare(fact, "=", value)
}

func simpleAndQuery() []string {
	q := puppetquery.ActiveNodes()
	for i := 0; i < flag.NArg(); i++ {
		q = puppetquery.And(q, ParseEqualFact(flag.Arg(i)))
	}
	n, err := puppetquery.QueryNodes(q)
	if err != nil {
		log.Fatalln("ERROR: cannot query puppetdb: ", err)
	}
	return n
}

type pair struct {
	key, value string
}

// These four are required for sorting
type byKey []pair

func (p byKey) Len() int           { return len(p) }
func (p byKey) Less(i, j int) bool { return p[i].key < p[j].key }
func (p byKey) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// print all facts (maybe limited to allowedFacts) of nodes.
// Return, whether we printed any lines at all
func factsOfNodesQuery(nodes []string, allowedFacts map[string]bool, nodeprefix bool) bool {
	output := false
	for _, node := range nodes {
		facts, err := puppetquery.QueryFacts(node)
		if err != nil {
			log.Fatalln("ERROR: cannot query puppetdb: ", err)
		}
		linearFacts := []pair{}
		for k, v := range facts {
			if allowedFacts != nil {
				// only record facts we like to register
				_, exists := allowedFacts[k]
				if !exists {
					continue
				}
			}

			linearFacts = append(linearFacts, pair{key: k, value: v})
		}
		sort.Sort(byKey(linearFacts))
		for _, v := range linearFacts {
			output = true
			if nodeprefix {
				fmt.Printf("%v:%v=%v\n", node, v.key, v.value)
			} else {
				fmt.Printf("%v=%v\n", v.key, v.value)
			}
		}
	}
	return output
}

func main() {
	var listFacts string
	var limitFacts string
	var verbose bool
	flag.StringVar(&listFacts, "l", "", "list all facts of a node")
	flag.StringVar(&limitFacts, "f", "", "limit display to comma separated list of facts")
	flag.BoolVar(&verbose, "v", false, "more verbose output")
	flag.Parse()

	switch {
	case listFacts != "":
		res := factsOfNodesQuery([]string{listFacts}, nil, false)
		if !res && verbose {
			log.Print("Query returned no results")
		}
	case limitFacts != "":
		allowedFacts := map[string]bool{}

		for _, allow := range strings.Split(limitFacts, ",") {
			allowedFacts[allow] = true
		}
		res := factsOfNodesQuery(simpleAndQuery(), allowedFacts, true)
		if !res && verbose {
			log.Print("Query returned no results. Check your fact limits.")
		}
	default:
		for _, node := range simpleAndQuery() {
			fmt.Println(node)
		}
	}
}
