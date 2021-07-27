package fetch_test

import (
	"fmt"
	"jp/thelow/static/fetch"
	"log"
	"testing"
	"time"
)

func TestToQuery(t *testing.T) {

	start := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2021, time.June, 30, 23, 59, 59, 0, time.Local)

	tags := map[string]string{
		"description": "ExpBlock取得",
	}

	q := fetch.Queries{
		Service:   "server1",
		Operation: "dungeon",
		Limit:     20,
		Offset:    0,
		Start:     start,
		End:       end,
		Tags:      &tags,
	}

	expect := fmt.Sprintf("service=server1&operation=dungeon&limit=20&offset=0&start=%d&end=%d&minDuration&maxDuration&lookback&tags=%%7B\"description\"%%3A\"ExpBlock取得\"%%7D", start.UnixNano()/1000000, end.UnixNano()/1000000)

	if q.Serialize() != expect {
		log.Fatalf("expect: %s\nactual: %s", expect, q.Serialize())
	}
}
