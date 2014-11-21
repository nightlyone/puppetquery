package puppetquery

import "encoding/json"

// A Resource as returned by puppetdb resource endpoint
type Resource struct {
	Certname   string                 `json:"certname"`
	Type       string                 `json:"type"`
	Title      string                 `json:"title"`
	Exported   bool                   `json:"exported"`
	Tags       []string               `json:"tags"`
	Sourcefile string                 `json:"sourcefile"`
	Sourceline int                    `json:"sourceline"`
	Parameters map[string]interface{} `json:"parameters"`
}

// CollectResources delivers collected resources of typ matching all of the tags provided.
func CollectResources(typ string, tags ...string) (resources []Resource, err error) {
	q := make(QueryString, 0, len(tags)*2+2)
	q = And(q, ActiveNodes())
	q = And(q, Match("type", typ))
	q = And(q, Match("exported", true))
	for _, tag := range tags {
		q = And(q, Match("tag", tag))

	}
	b, err := do("resources", q)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resources)
	return resources, err
}
