package reinc

import (
	"encoding/json"
	"jp/thelow/static/elastic"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"jp/thelow/static/mojang"
	"sort"
	"time"
)

func Count(config *model.Config) (*ReincList, error) {
	q := &elastic.ElasticQuery{
		Host:      config.Elastic.Host,
		Start:     fetch.StartOfMonth(config.Year, config.Month),
		End:       fetch.EndOfMonth(config.Year, config.Month),
		QueryFile: "reincarnations.json",
	}

	data, err := elastic.Fetch(q)
	if err != nil {
		return nil, err
	}

	aggs := new(elastic.AggsReincarnation)
	if err = json.Unmarshal(data, aggs); err != nil {
		return nil, err
	}

	buckets := aggs.Aggregations.PlayerUUIDGroup.Buckets
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].CountReincs.Value > buckets[j].CountReincs.Value
	})

	n := 10
	if len(buckets) < 10 {
		n = len(buckets)
	}

	result := &ReincList{
		Data:    make(map[string]int),
		Ranking: make([]string, 10),
		Year:    config.Year,
		Month:   time.Month(config.Month),
	}

	for i := 0; i < n; i++ {
		bucket := buckets[i]
		uuid := bucket.Key
		name, err := mojang.FetchPlayerName(uuid)

		if err != nil {
			continue
		}

		result.Data[name] = int(bucket.CountReincs.Value)
		result.Ranking[i] = name
	}

	return result, nil
}
