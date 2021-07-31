package reinc

import (
	"fmt"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"jp/thelow/static/progress"
	"time"
)

type ReincType string

const (
	SWORD  ReincType = "転生:SWORD"
	MAGIC  ReincType = "転生:MAGIC"
	BOW    ReincType = "転生:BOW"
	ALL    ReincType = "全転生"
	format           = "01-02 15:04"
)

func CountInMonth(q *fetch.QueryBuilder, year, month int) *ReincResult {
	r := NewResult(year, month)

	CountInMonthByType(*q, SWORD, r)
	CountInMonthByType(*q, MAGIC, r)
	CountInMonthByType(*q, BOW, r)
	CountInMonthByType(*q, ALL, r)

	return r
}

func CountInMonthByType(q fetch.QueryBuilder, typ ReincType, r *ReincResult) {
	end := q.End
	q.SetTags(&map[string]string{"description": string(typ)})

	monthDays := float64(end.Day())

	fmt.Printf("Count %s\n", typ)
	bar := progress.NewProgressBar(string(typ))

	for {
		p := 1.0 - (float64(end.Day()) / monthDays)
		bar.SetTitle(end.Format(format))
		bar.SetProgress(p)

		t := fetch.FetchTraces(&q)

		for _, trace := range t.Data {
			when := CountFromTrace(&trace, typ, r)

			if when.Before(end) {
				end = when
			}
		}
		q.End = end

		if len(t.Data) < 1000 {
			break
		}

		time.Sleep(1 * time.Second)
	}

	bar.CompleteProgress()
}

func CountFromTrace(t *model.Trace, typ ReincType, r *ReincResult) time.Time {
	var oldest *time.Time = nil
	for _, span := range t.Spans {
		when := CountFromSpan(&span, typ, r)

		if oldest == nil {
			oldest = &when
		} else if when.Before(*oldest) {
			oldest = &when
		}
	}

	return *oldest
}

func CountFromSpan(s *model.Span, typ ReincType, r *ReincResult) time.Time {

	uuid := s.GetTagValue("uuid")
	mcid := s.GetTagValue("mcid")
	when := time.Unix(s.StartTime/1000000, 0)

	if uuid == nil || mcid == nil {
		return when
	}

	for _, log := range s.Logs {
		if !log.ContainsKey("description") || !log.ContainsKey("time") {
			continue
		}

		description := log.GetValue("description")
		occurAtStr := log.GetValue("time")

		if description == nil || occurAtStr == nil {
			continue
		}

		if *description != string(typ) {
			continue
		}

		occurAt, err := time.Parse("2006/01/02 15:04:05", *occurAtStr)
		if err != nil {
			break
		}

		if occurAt.Year() == r.GetYear() && occurAt.Month() == r.GetMonth() {
			r.Increment(*uuid, *mcid, s.SpanID)
		}

		break
	}

	return when
}
