package fetch

import (
	"encoding/json"
	"io/ioutil"
	"jp/thelow/static/model"
	"log"
	"net/http"
	"time"
)

func FetchTraces(q *QueryBuilder) *model.Traces {
	req := q.Build()
	res := new(model.Traces)
	body := get(req)
	if err := json.Unmarshal(body, res); err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func get(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("failed to request: %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("failed to read body: %s", err)
	}
	return body
}

func StartOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

func EndOfMonth(year, month int) time.Time {
	next := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	return next.Add(-1 * time.Second)
}
