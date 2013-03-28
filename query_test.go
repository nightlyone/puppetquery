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

var opTests = []struct {
	in   QueryString
	want string
}{
	{
		in:   And(FactCompare("kernel", "=", "Linux"), FactCompare("uptime_days", ">", 30)),
		want: `["and",["=",["fact","kernel"],"Linux"],["\u003e",["fact","uptime_days"],30]]`,
	},
	{
		in:   And(ActiveNodes(), And(FactCompare("kernel", "=", "Linux"), FactCompare("uptime_days", ">", 30))),
		want: `["and",["=",["fact","kernel"],"Linux"],["\u003e",["fact","uptime_days"],30],["=",["node","active"],true]]`,
	},
	{
		in:   Or(FactCompare("kernel", "=", "Linux"), FactCompare("uptime_days", ">", 30)),
		want: `["or",["=",["fact","kernel"],"Linux"],["\u003e",["fact","uptime_days"],30]]`,
	},
}

func TestOps(t *testing.T) {
	for i, test := range opTests {
		got := test.in.ToJson()
		if got != test.want {
			t.Error(i, ": got: ", got, " want: ", test.want)
		} else {
			t.Log(i, ": got: ", got, " want: ", test.want)
		}
	}
}
