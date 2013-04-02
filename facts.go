package puppetquery

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// this is the returned puppetdb API data for facts
type rawFacts struct {
	Facts map[string]string `json:"facts"`
}

// query facts api endpoint returning the list of facts for a node
func QueryFacts(node string) (facts map[string]string, err error) {
	req, err := http.NewRequest("GET", endpoint+"/facts/"+url.QueryEscape(node), nil)
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Print("FATAL: node ", node)
		log.Print("FATAL: request ", resp.Request)
		log.Fatal("FATAL: Status != 200, got ", resp.Status)
	}
	defer resp.Body.Close()

	var raw rawFacts

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	return raw.Facts, nil
}
