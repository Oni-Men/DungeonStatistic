package elastic

import (
	"io/ioutil"
	"strings"
	"time"
)

type ElasticQuery struct {
	Host      string
	QueryFile string
	Start     time.Time
	End       time.Time
}

type AggsReincarnation struct {
	Aggregations struct {
		PlayerUUIDGroup struct {
			Buckets []struct {
				Key         string `json:"key"`
				CountReincs struct {
					Value float64 `json:"value"`
				} `json:"count-reincs"`
			} `json:"buckets"`
		} `json:"player-uuid-group"`
	} `json:"aggregations"`
}

type AggsDungeonClear struct {
	Aggregations struct {
		DungeonGroup struct {
			Buckets []struct {
				Key         string `json:"key"`
				CountClears struct {
					Value float64
				} `json:"count-clears"`
			} `json:"buckets"`
		} `json:"dungeon-group"`
	} `json:"aggregations"`
}

const ISO_OFFSET = "2006-01-02T15:04:05+07:00"

func ReadQuery(filename string) (string, error) {
	data, err := ioutil.ReadFile("./elastic/queries/" + filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func FormatTime(src string, start, end time.Time) string {
	s := start.Format(ISO_OFFSET)
	e := end.Format(ISO_OFFSET)
	src = strings.ReplaceAll(src, "$start", s)
	src = strings.ReplaceAll(src, "$end", e)
	return src
}
