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

func hasOp(op string, tree QueryString) bool {
	if len(tree) > 0 {
		o, ok := tree[0].(string)
		return ok && o == op
	}
	return false
}

func BinOp(binop string, left, right QueryString) QueryString {
	switch {
	case len(left) == 0:
		return right
	case len(right) == 0:
		return left
	case len(left) > 2 && hasOp(binop, left):
		return append(left, right)
	case len(right) > 2 && hasOp(binop, right):
		return append(right, left)
	}
	return QueryString{Any(binop), left, right}
}

func And(left, right QueryString) QueryString {
	return BinOp("and", left, right)
}

func Or(left, right QueryString) QueryString {
	return BinOp("or", left, right)
}

func Not(tree QueryString) QueryString {
	switch {
	case len(tree) == 0:
		return tree
	case len(tree) == 2 && hasOp("not", tree):
		return tree[1].(QueryString)
	}
	return QueryString{Any("not"), tree}
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
