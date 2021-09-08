package elastic

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Fetch(q *ElasticQuery) ([]byte, error) {
	u, err := url.Parse(q.Host)
	if err != nil {
		return nil, err
	}

	u.Path = "/thelow-server*/_search"

	query, err := ReadQuery(q.QueryFile)
	if err != nil {
		return nil, err
	}

	query = FormatTime(query, q.Start, q.End)

	req, err := http.NewRequest("GET", u.String(), bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, err
}
