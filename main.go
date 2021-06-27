package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"thelow/survey/fetch"
	"thelow/survey/model"
)

var (
	config = new(model.Config)
	year   *int
	month  *int
	output = "."
)

var monthNames = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func main() {
	mode := flag.String("mode", "all", "all/fetch/image")
	month = flag.Int("month", int(time.Now().Month()), "処理対象の月(1-12)")
	year = flag.Int("year", time.Now().Year(), "処理対象の年")
	flag.Parse()

	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = json.Unmarshal(data, config); err != nil {
		log.Fatalf(err.Error())
	}

	output = "./data/" + strconv.Itoa(*year) + "/" + monthNames[*month-1] + "/"

	switch *mode {
	case "all":
		completes := modeFetch(config.Host)
		modeImage(completes)
	case "fetch":
		modeFetch(config.Host)
	case "image":
		if read, err := ioutil.ReadFile(output + "completes.json"); err == nil {
			completes := make(map[string]int)
			if err = json.Unmarshal(read, &completes); err != nil {
				log.Fatalf(err.Error())
			}
			modeImage(completes)
		}
	}
}

func modeImage(completes map[string]int) {
	data, err := ioutil.ReadFile("./template.svg")
	if err != nil {
		log.Fatalf(err.Error())
	}
	text := string(data)

	sorted := toDungeonCompletedDatas(completes)

	literal := monthNames[*month]

	text = strings.Replace(text, "{#month}", strconv.Itoa(*month), 1)
	text = strings.Replace(text, "{#month_literal}", literal, 1)

	total := fmt.Sprintf("%.2fk", float32(completes["__completes__"])/1000.0)
	text = strings.Replace(text, "{#total}", total, 1)

	for i := 2; i < 12; i++ {
		d := sorted[i]
		text = strings.Replace(text, fmt.Sprintf("{#dungeon_name_%d}", i-2), d.Name, 1)
		text = strings.Replace(text, fmt.Sprintf("{#count_%d}", i-2), strconv.Itoa(d.Count), 1)
	}
	if err = ioutil.WriteFile(output+"completes.svg", []byte(text), fs.ModePerm); err != nil {
		log.Fatalf(err.Error())
	}
}

func modeFetch(host string) map[string]int {
	start := startOfMonth(*month)
	end := endOfMonth(*month)

	completes := make(map[string]int)

	for {
		t := fetch.FetchTraces(host, start, end)

		if len(t.Data) == 0 {
			break
		}

		sorByTimestamp(t)

		spans := t.Data[0].Spans
		if len(spans) != 0 {
			end = time.Unix(spans[0].StartTime/1000000, 0)
			println(end.Format(time.RFC1123Z))
		}

		list := toDungeonStrategyDatas(t)

		for _, data := range *list {
			if data.Complete {
				completes[data.Name]++
				completes["__completes__"]++
			}
			completes["__tries__"]++
		}

		time.Sleep(1 * time.Second)
	}

	if data, err := json.Marshal(completes); err == nil {
		ioutil.WriteFile(output+"completes.json", data, fs.ModePerm)
	}

	return completes
}

func toDungeonStrategyDatas(t *model.Traces) *[]*model.DungeonStrategyData {
	list := make([]*model.DungeonStrategyData, 0, 100)

	for _, trace := range t.Data {
		for _, span := range trace.Spans {
			if span.OperationName != "dungeon" {
				continue
			}

			data := new(model.DungeonStrategyData)

			if !model.TagContainsKey(span.Tags, "ダンジョン名") {
				continue
			}

			data.Name = *model.GetValueFromTags(span.Tags, "ダンジョン名")

			for _, log := range span.Logs {

				if !model.TagContainsKey(log.Fields, "description") {
					continue
				}

				desc := *model.GetValueFromTags(log.Fields, "description")

				if desc == "ExpBlock取得" {
					data.Complete = true
					break
				}
			}

			list = append(list, data)
		}
	}

	return &list
}

func toDungeonCompletedDatas(completes map[string]int) []model.DungeonCompletesData {
	sorted := make([]model.DungeonCompletesData, 0, len(completes))

	for name, count := range completes {
		sorted = append(sorted, model.DungeonCompletesData{
			Name:  name,
			Count: count,
		})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count > sorted[j].Count
	})

	return sorted
}

func sorByTimestamp(t *model.Traces) {
	sort.Slice(t.Data, func(i, j int) bool {
		return t.Data[i].Spans[0].StartTime < t.Data[j].Spans[0].StartTime
	})
}

func startOfMonth(month int) time.Time {
	return time.Date(*year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

func endOfMonth(month int) time.Time {
	next := time.Date(*year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	next.Add(-1 * time.Second)
	return next
}
