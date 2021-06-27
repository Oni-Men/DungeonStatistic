package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"thelow/survey/model"
	"time"
)

func FetchTraces(host string, start, end time.Time) *model.Traces {
	q := Queries{
		Service:   "server1",
		Operation: "dungeon",
		Limit:     100,
		Offset:    0,
		Start:     start,
		End:       end,
	}

	return traces(host, q)
}

func traces(host string, q Queries) *model.Traces {
	req := endpoint(host, "traces") + "?" + q.ToQuery()
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

func endpoint(host string, query ...string) string {
	return fmt.Sprintf("%s/api/%s", host, strings.Join(query, "/"))
}
