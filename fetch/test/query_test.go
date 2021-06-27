package query_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"thelow/survey/fetch"
)

func TestToQuery(t *testing.T) {

	start := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2021, time.June, 30, 23, 59, 59, 0, time.Local)

	q := fetch.Queries{
		Service:   "server1",
		Operation: "dungeon",
		Limit:     20,
		Offset:    0,
		Start:     start,
		End:       end,
	}

	expect := fmt.Sprintf("service=server1&operation=dungeon&limit=20&offset=0&start=%d&end=%d&minDuration&maxDuration&lookback", start.UnixNano()/1000000, end.UnixNano()/1000000)

	if q.ToQuery() != expect {
		log.Fatalf("expect: %s\nactual: %s", expect, q.ToQuery())
	}
}
