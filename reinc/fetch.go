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
	format           = "01-02 15:04"
)

func CountInMonth(host string, year, month int) *ReincResult {
	r := NewResult(year, month)

	CountInMonthByType(host, SWORD, year, month, r)
	CountInMonthByType(host, MAGIC, year, month, r)
	CountInMonthByType(host, BOW, year, month, r)

	return r
}

func CountInMonthByType(host string, typ ReincType, year, month int, r *ReincResult) {
	start := fetch.StartOfMonth(year, month)
	end := fetch.EndOfMonth(year, month)

	q := &fetch.Queries{
		Service: "server1",
		Limit:   100,
		Offset:  0,
		Start:   start,
		End:     end,
		Tags:    &map[string]string{"description": string(typ)},
	}

	monthDays := float64(end.Day())

	fmt.Printf("Count %s\n", typ)
	bar := progress.NewProgressBar(string(typ))

	for {
		t := fetch.FetchTraces(host, q)

		for _, trace := range t.Data {
			when := CountFromTrace(&trace, typ, r)

			if when.Before(end) {
				end = when
			}
		}

		if len(t.Data) < 100 {
			break
		}

		p := 1.0 - (float64(end.Day()) / monthDays)
		bar.SetTitle(end.Format(format))
		bar.SetProgress(p)

		q.End = end

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
		if !log.ContainsKey("description") {
			continue
		}

		description := log.GetValue("description")

		if description == nil {
			continue
		}

		if *description != string(typ) {
			continue
		}

		r.Increment(*uuid, *mcid)
		break
	}

	return when
}
