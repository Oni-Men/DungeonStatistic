package reinc

import (
	"fmt"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"jp/thelow/static/progress"
	"strconv"
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

var (
	count_msg = map[ReincType]string{
		SWORD: "SWORDの転生回数",
		MAGIC: "MAGICの転生回数",
		BOW:   "BOWの転生回数",
	}
)

func CountInMonth(q *fetch.QueryBuilder, year, month int) *ReincResult {
	r := NewResult(year, month)

	CountInMonthByType(*q, SWORD, r)
	CountInMonthByType(*q, MAGIC, r)
	CountInMonthByType(*q, BOW, r)
	CountInMonthByType(*q, ALL, r)

	return r
}

func CountInMonthByType(q fetch.QueryBuilder, reincType ReincType, r *ReincResult) {

	end := q.End
	q.SetTags(&map[string]string{"description": string(reincType)})

	monthDays := float64(end.Day())

	fmt.Printf("Count %s\n", reincType)
	bar := progress.NewProgressBar(string(reincType))

	for {
		p := 1.0 - (float64(end.Day()) / monthDays)
		bar.SetTitle(end.Format(format))
		bar.SetProgress(p)

		t := fetch.FetchTraces(&q)

		for _, trace := range t.Data {
			when := CountFromTrace(&trace, reincType, r)

			if when.Before(end) {
				end = when
			}
		}
		q.End = end

		if len(t.Data) < q.Limit {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	bar.CompleteProgress()
}

func CountFromTrace(t *model.Trace, reincType ReincType, r *ReincResult) time.Time {
	var oldest *time.Time = nil
	for _, span := range t.Spans {
		when := CountFromSpan(&span, reincType, r)

		if oldest == nil {
			oldest = &when
		} else if when.Before(*oldest) {
			oldest = &when
		}
	}

	return *oldest
}

func CountFromSpan(s *model.Span, reincType ReincType, r *ReincResult) time.Time {

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
		countStr := log.GetValue(count_msg[reincType])

		if description == nil || occurAtStr == nil || countStr == nil {
			continue
		}

		if *description != string(reincType) {
			continue
		}

		occurAt, err := time.Parse("2006/01/02 15:04:05", *occurAtStr)
		if err != nil {
			break
		}

		count, err := strconv.Atoi(*countStr)
		if err != nil {
			break
		}

		if occurAt.Year() == r.GetYear() && occurAt.Month() == r.GetMonth() {
			r.Increment(*uuid, *mcid, s.SpanID, reincType, occurAt, count)
		}

		break
	}

	return when
}
