package puppetquery

import (
	"encoding/json"
	"log"
)

type Any interface{}
type QueryString []Any

func ActiveNodes() QueryString {
	return QueryString{Any("="), QueryString{Any("node"), Any("active")}, Any(true)}
}

func And(left, right QueryString) QueryString {
	return QueryString{Any("and"), left, right}
}

func Or(left, right QueryString) QueryString {
	return QueryString{Any("or"), left, right}
}

func FactCompare(name, op string, value Any) QueryString {
	return QueryString{op, QueryString{Any("fact"), Any(name)}, value}
}

func (q *QueryString) ToJson() string {
	b, err := json.Marshal(q)
	if err != nil {
		log.Println("ERROR: cannot marshal: ", err)
		return ""
	}
	return string(b)
}
