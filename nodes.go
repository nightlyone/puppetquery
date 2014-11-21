package puppetquery

import "encoding/json"

// query node endpoint returning a list of nodes matching the query
func QueryNodes(query QueryString) (nodes []string, err error) {
	b, err := do("nodes", query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &nodes)
	return nodes, err
}
