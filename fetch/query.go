package fetch

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type LookBackType string

const ONE_HOUR LookBackType = "1h"
const TWO_HOUR LookBackType = "2h"
const THREE_HOUR LookBackType = "3h"
const SIX_HOUR LookBackType = "6h"
const TWELVE_HOUR LookBackType = "12h"
const ONE_DAY LookBackType = "1d"
const TWO_DAY LookBackType = "2d"

type QueryBuilder struct {
	Host        string             `label:"host"`
	Service     string             `label:"service"`
	Operation   string             `label:"operation"`
	Limit       int                `label:"limit"`
	Offset      int                `label:"offset"`
	Start       time.Time          `label:"start"`
	End         time.Time          `label:"end"`
	Tags        *map[string]string `label:"tags"`
	MinDuration time.Duration      `label:"minDuration"`
	MaxDuration time.Duration      `label:"maxDuration"`
	LookBack    LookBackType       `label:"lookback"`
}

func NewQueryBuilder(host string) *QueryBuilder {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	q := &QueryBuilder{
		Host:      host,
		Service:   "server1",
		Operation: "",
		Limit:     100,
		Offset:    0,
		Start:     StartOfMonth(year, month),
		End:       EndOfMonth(year, month),
		Tags:      &map[string]string{},
	}
	return q
}

func (q *QueryBuilder) Build() string {
	res := make([]string, 0, 10)
	res = append(res, fmt.Sprintf("service=%s", q.Service))
	if q.Operation != "" {
		res = append(res, fmt.Sprintf("operation=%s", q.Operation))
	}
	res = append(res, fmt.Sprintf("limit=%d", q.Limit))
	res = append(res, fmt.Sprintf("offset=%d", q.Offset))
	res = append(res, fmt.Sprintf("start=%d", q.Start.UnixNano()/1000))
	res = append(res, fmt.Sprintf("end=%d", q.End.UnixNano()/1000))
	if q.MinDuration != 0 {
		res = append(res, fmt.Sprintf("minDuration=%d", q.MinDuration.Milliseconds()))
	} else {
		res = append(res, "minDuration")
	}

	if q.MaxDuration != 0 {
		res = append(res, fmt.Sprintf("maxDuration=%d", q.MaxDuration.Milliseconds()))
	} else {
		res = append(res, "maxDuration")
	}

	if q.LookBack != "" {
		res = append(res, fmt.Sprintf("lookback=%s", q.LookBack))
	} else {
		res = append(res, "lookback=custom")
	}

	if q.Tags != nil {
		tags := make([]string, 0, len(*q.Tags))
		for k, v := range *q.Tags {
			tags = append(tags, fmt.Sprintf("\"%s\":\"%s\"", k, v))
		}
		escaped := url.QueryEscape("{" + strings.Join(tags, ",") + "}")
		res = append(res, "tags="+escaped)
	}

	return q.Host + "/api/traces?" + strings.Join(res, "&")
}

func (q *QueryBuilder) SetService(service string) {
	q.Service = service
}

func (q *QueryBuilder) SetOperation(operation string) {
	q.Operation = operation
}

func (q *QueryBuilder) SetLimit(limit int) {
	q.Limit = limit
}

func (q *QueryBuilder) SetStart(start time.Time) {
	q.Start = start
}

func (q *QueryBuilder) SetEnd(end time.Time) {
	q.End = end
}

func (q *QueryBuilder) SetTags(tags *map[string]string) {
	q.Tags = tags
}

func (q *QueryBuilder) AddTag(key, value string) {
	(*q.Tags)[key] = value
}

func (q *QueryBuilder) SetHost(host string) {
	q.Host = host
}
