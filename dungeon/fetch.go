package dungeon

import (
	"fmt"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"jp/thelow/static/progress"
	"time"
)

const (
	format = "01-02 15:04"
)

func FetchCompletions(host string, year, month int) *DungeonResult {
	start := fetch.StartOfMonth(year, month)
	end := fetch.EndOfMonth(year, month)
	tags := map[string]string{"description": "ExpBlock取得"}

	monthDays := float64(end.Day())

	result := NewResult(year, month)

	fmt.Println("Count ダンジョン攻略")
	bar := progress.NewProgressBar("dungeon")

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

		p := 1.0 - (float64(end.Day()) / monthDays)
		bar.SetTitle(end.Format(format))
		bar.SetProgress(p)

		time.Sleep(1 * time.Second)
	}

	bar.CompleteProgress()

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
