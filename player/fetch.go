package player

import (
	"fmt"
	"jp/thelow/static/fetch"
	"time"
)

func FetchPlayersDungeonData(host, mcid, dungeon string) {
	q := fetch.NewQueryBuilder(host)
	q.SetOperation(mcid)
	q.SetLimit(3)
	q.SetTags(&map[string]string{
		"description": "ExpBlock取得",
		"mcid":        mcid,
		"ダンジョン名":      dungeon,
	})

	now := time.Now()
	weekAgo := now.Add(-7 * 24 * time.Hour)

	q.SetStart(weekAgo)
	q.SetEnd(now)

	traces := fetch.FetchTraces(q)

	var durationSum int64 = 0
	var total int64 = 0

	for _, t := range traces.Data {
		for _, s := range t.Spans {
			for _, log := range s.Logs {

				name := log.GetValue("ダンジョン名")
				desc := log.GetValue("description")

				if name == nil || desc == nil {
					continue
				}

				if *name != dungeon || *desc != "ExpBlock取得" {
					continue
				}

				durationSum += s.Duration / 1000000
				total++
			}

		}
	}

	average := durationSum / total
	fmt.Printf("%sの攻略に成功した回数: %d\n", dungeon, total)
	fmt.Printf("平均攻略時間: %02d:%02d\n", average/60, average%60)
}
