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

func CountCompletions(q *fetch.QueryBuilder, year, month int) *DungeonResult {
	q.SetOperation("dungeon")
	q.SetTags(&map[string]string{"description": "ExpBlock取得"})

	end := q.End
	monthDays := float64(end.Day())
	result := NewResult(year, month)

	fmt.Println("Count ダンジョン攻略")
	bar := progress.NewProgressBar("dungeon")

	for {
		p := 1.0 - (float64(end.Day()) / monthDays)
		bar.SetTitle(end.Format(format))
		bar.SetProgress(p)

		t := fetch.FetchTraces(q)

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

		if len(t.Data) < 1000 {
			break
		}

		time.Sleep(2 * time.Second)
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
