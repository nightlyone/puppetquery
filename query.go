package puppetquery

import (
	"encoding/json"
	"log"
)

//NOTE(nightlyone): any might be replaceable by string
type any interface{}

//internal format to manage and massage Query
type QueryString []any

// Match returns a query matching for existance of certain (key,value) pairs
func Match(key string, value interface{}) QueryString {
	return QueryString{any("="), any(key), any(value)}
}

// Returns a query for active nodes only
func ActiveNodes() QueryString {
	return QueryString{any("="), QueryString{any("node"), any("active")}, any(true)}
}

func hasOp(op string, tree QueryString) bool {
	if len(tree) > 0 {
		o, ok := tree[0].(string)
		return ok && o == op
	}
	return false
}

//returns Query <left> <binop> <right> and reduces left or right, if it contains binop already
// e.g. (a and b) and c becomes and(a,b,c)
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
	return QueryString{any(binop), left, right}
}

// constructs and(left, right) from left and right
func And(left, right QueryString) QueryString {
	return BinOp("and", left, right)
}

// constructs or(left, right) from left and right
func Or(left, right QueryString) QueryString {
	return BinOp("or", left, right)
}

//returns Query "not" <tree> <binop> <right> and reduces not(not(tree)) to tree again
func Not(tree QueryString) QueryString {
	switch {
	case len(tree) == 0:
		return tree
	case len(tree) == 2 && hasOp("not", tree):
		return tree[1].(QueryString)
	}
	return QueryString{any("not"), tree}
}

// returns a fact comparison query
func FactCompare(name, op string, value any) QueryString {
	return QueryString{op, QueryString{any("fact"), any(name)}, value}
}

func (q *QueryString) ToJson() string {
	b, err := json.Marshal(q)
	if err != nil {
		log.Println("ERROR: cannot marshal: ", err)
		return ""
	}
	return string(b)
}
