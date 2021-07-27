package fetch

import (
	"fmt"
	"io/ioutil"
	"jp/thelow/static/model"
	"log"
	"net/http"
	"strings"
	"time"
)

func FetchTraces(host string, q *Queries) *model.Traces {
	return q.Execute(host)
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

func endpoint(host string, query ...string) string {
	return fmt.Sprintf("%s/api/%s", host, strings.Join(query, "/"))
}

func StartOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

func EndOfMonth(year, month int) time.Time {
	next := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	next.Add(-1 * time.Second)
	return next
}
