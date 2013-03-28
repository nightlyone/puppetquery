package puppetquery

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func QueryNodes(query QueryString) (nodes []string, err error) {
	query_string := "query=" + query.ToJson()
	req, err := http.NewRequest("GET", endpoint+"/nodes"+"?"+url.QueryEscape(query_string), nil)
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Print("FATAL: request ", resp.Request)
		log.Fatal("FATAL: Status != 200, got ", resp.Status)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&nodes)
	return nodes, err
}
