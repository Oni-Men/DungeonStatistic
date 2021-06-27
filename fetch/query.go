package fetch

import (
	"fmt"
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

type Queries struct {
	Service     string        `label:"service"`
	Operation   string        `label:"operation"`
	Limit       int           `label:"limit"`
	Offset      int           `label:"offset"`
	Start       time.Time     `label:"start"`
	End         time.Time     `label:"end"`
	MinDuration time.Duration `label:"minDuration"`
	MaxDuration time.Duration `label:"maxDuration"`
	LookBack    LookBackType  `label:"lookback"`
}

func (q *Queries) ToQuery() string {
	res := make([]string, 0, 9)
	res = append(res, fmt.Sprintf("service=%s", q.Service))
	res = append(res, fmt.Sprintf("operation=%s", q.Operation))
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
		res = append(res, "lookback")
	}

	return strings.Join(res, "&")
}
