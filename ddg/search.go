package ddg

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var URL url.URL

func init() {
	loc, err := url.Parse("https://api.duckduckgo.com")
	if err != nil {
		panic("(bug) duckduckgo API url is invalid: " + err.Error())
	}

	URL = *loc
}

func Search(q string) (*Info, error) {
	v := url.Values{}
	// We only work with sane JSON
	v.Add("format", "json")
	// This is to avoid having html in the JSON result
	v.Add("no_html", "1")
	v.Add("q", q)

	u := URL
	u.RawQuery = v.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var info Info

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
