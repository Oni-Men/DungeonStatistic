package dungeon

import (
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"time"
)

func FetchCompletions(host string, year, month int) *DungeonResult {
	start := fetch.StartOfMonth(year, month)
	end := fetch.EndOfMonth(year, month)
	tags := map[string]string{"description": "ExpBlock取得"}

	result := NewResult()

	for {
		t := fetch.FetchTraces(host, &fetch.Queries{
			Service:   "server1",
			Operation: "dungeon",
			Limit:     100,
			Offset:    0,
			Start:     start,
			End:       end,
			Tags:      &tags,
		})

		for _, trace := range t.Data {
			for _, span := range trace.Spans {
				s := time.Unix(span.StartTime/1000000, 0)
				if s.Before(end) {
					end = s
				}
			}
		}

		for _, t := range t.Data {
			CountCompletesFromTrace(&t, result)
		}

		if len(t.Data) < 100 {
			break
		}

		println(end.Format(time.RFC1123Z))
		time.Sleep(250 * time.Millisecond)
	}

	return result
}

func CountCompletesFromTrace(t *model.Trace, result *DungeonResult) {
	for _, span := range t.Spans {
		CountCompletesFromSpan(&span, result)
	}
}

func CountCompletesFromSpan(s *model.Span, result *DungeonResult) {

	for _, log := range s.Logs {

		if !log.ContainsKey("description") {
			continue
		}

		description := log.GetValue("description")
		dungeonName := log.GetValue("ダンジョン名")

		if description == nil || dungeonName == nil {
			continue
		}

		if *description == "ExpBlock取得" {
			result.Increment(*dungeonName)
		}
	}
}
