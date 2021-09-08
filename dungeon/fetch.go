package dungeon

import (
	"encoding/json"
	"jp/thelow/static/elastic"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"sort"
)

func Count(config *model.Config) (*DungeonList, error) {
	q := elastic.ElasticQuery{
		Host:      config.Elastic.Host,
		Start:     fetch.StartOfMonth(config.Year, config.Month),
		End:       fetch.EndOfMonth(config.Year, config.Month),
		QueryFile: "dungeon_clears.json",
	}

	data, err := elastic.Fetch(&q)
	if err != nil {
		return nil, err
	}

	println(string(data))

	aggs := new(elastic.AggsDungeonClear)
	if err = json.Unmarshal(data, aggs); err != nil {
		return nil, err
	}

	buckets := aggs.Aggregations.DungeonGroup.Buckets
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].CountClears.Value > buckets[j].CountClears.Value
	})

	list := &DungeonList{
		Year:    config.Year,
		Month:   config.Month,
		Ranking: make([]DungeonClear, 0, len(buckets)),
	}

	for _, bucket := range buckets {
		count := int(bucket.CountClears.Value)
		d := DungeonClear{
			Name:  bucket.Key,
			Count: count,
		}
		list.Ranking = append(list.Ranking, d)
		list.Total += count
	}

	return list, nil
}
