package puppetquery

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpError struct {
	code  int
	query QueryString
	body  string
}

func (e HttpError) Error() string {
	return fmt.Sprintf("http status '%s', query %s, body %s", http.StatusText(e.code), e.query.ToJson(), e.body)
}

func do(service string, query QueryString) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint+"/"+service+"?query="+url.QueryEscape(query.ToJson()), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		consume := func(r io.Reader) string {
			b, _ := ioutil.ReadAll(r)
			if len(b) > 0 {
				return string(b)
			} else {
				return ""
			}
		}
		return nil, HttpError{
			code:  resp.StatusCode,
			query: query,
			body:  consume(resp.Body),
		}
	}
	return ioutil.ReadAll(resp.Body)
}
