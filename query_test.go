package puppetquery

import "testing"

func TestActiveNodes(t *testing.T) {
	q := ActiveNodes()
	want := `["=",["node","active"],true]`
	got := q.ToJson()
	if got != want {
		t.Error("got: ", got, " want: ", want)
	} else {
		t.Log("got: ", got, " want: ", want)
	}
}

func TestFactCompare(t *testing.T) {
	q := FactCompare("kernel", "=", "Linux")
	want := `["=",["fact","kernel"],"Linux"]`
	got := q.ToJson()
	if got != want {
		t.Error("got: ", got, " want: ", want)
	} else {
		t.Log("got: ", got, " want: ", want)
	}
}

func TestAnd(t *testing.T) {
	q := And(FactCompare("kernel", "=", "Linux"), FactCompare("uptime_days", ">", 30))
	want := `["and",["=",["fact","kernel"],"Linux"],["\u003e",["fact","uptime_days"],30]]`
	got := q.ToJson()
	if got != want {
		t.Error("got: ", got, " want: ", want)
	} else {
		t.Log("got: ", got, " want: ", want)
	}
}
func TestOr(t *testing.T) {
	q := Or(FactCompare("kernel", "=", "Linux"), FactCompare("uptime_days", ">", 30))
	want := `["or",["=",["fact","kernel"],"Linux"],["\u003e",["fact","uptime_days"],30]]`
	got := q.ToJson()
	if got != want {
		t.Error("got: ", got, " want: ", want)
	} else {
		t.Log("got: ", got, " want: ", want)
	}
}
