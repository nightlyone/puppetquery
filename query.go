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

func BinOp(binop string, left, right QueryString) QueryString {
	if len(left) > 2 {
		if op, ok := left[0].(string); ok && op == binop {
			return append(left, right)
		}
	}
	if len(right) > 2 {
		if op, ok := right[0].(string); ok && op == binop {
			return append(right, left)
		}
	}
	return QueryString{Any(binop), left, right}
}

func And(left, right QueryString) QueryString {
	return BinOp("and", left, right)
}

func Or(left, right QueryString) QueryString {
	return BinOp("or", left, right)
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
