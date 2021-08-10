package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"

	"jp/thelow/static/dungeon"
	"jp/thelow/static/fetch"
	"jp/thelow/static/model"
	"jp/thelow/static/player"
	"jp/thelow/static/reinc"
	"jp/thelow/static/template"
)

var (
	config = new(model.Config)
	year   *int
	month  *int
	mcid   *string
	svg    *bool
	output = "."
)

func main() {
	now := time.Now()
	year = flag.Int("year", now.Year(), "集計対象の年(デフォルトは今年)")
	month = flag.Int("month", int(now.Month()), "集計対象の月(デフォルトは今月)")
	mcid = flag.String("mcid", "", "mcid")
	svg = flag.Bool("image", false, "テンプレートからSVGを生成するか")

	flag.Parse()

	model.LoadConfig(config)
	prepareOutputDir(*year, *month)

	if *mcid != "" {
		fmt.Printf("Fetch %s's data...\n", *mcid)
		player.FetchPlayersDungeonData(config.Host, *mcid, "クラバスタ")
	} else {
		saveDungeonCompletions()
		saveReincarnations()
	}
}

func saveDungeonCompletions() {
	q := fetch.NewQueryBuilder(config.Host)
	q.SetStart(fetch.StartOfMonth(*year, *month))
	q.SetEnd(fetch.EndOfMonth(*year, *month))
	q.SetLimit(500)

	dungeonResult := dungeon.CountCompletions(q, *year, *month)
	writeJSON("completions.json", dungeonResult)

	if *svg {

		t := string(template.ReadTemplate(template.DUNGEON))
		t = dungeon.CreateImage(dungeonResult, t)

		if err := ioutil.WriteFile(output+string(template.DUNGEON), []byte(t), fs.ModePerm); err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func saveReincarnations() {
	q := fetch.NewQueryBuilder(config.Host)
	q.SetStart(fetch.StartOfMonth(*year, *month))
	q.SetEnd(fetch.EndOfMonth(*year, *month))
	q.SetLimit(100)

	reincResult := reinc.CountInMonth(q, *year, *month)
	writeJSON("reincs.json", reincResult)

	if *svg {

		t := string(template.ReadTemplate(template.REINCARNATION))
		t = reinc.CreateImage(reincResult, t)

		svgFile := output + string(template.REINCARNATION)

		if err := ioutil.WriteFile(svgFile, []byte(t), fs.ModePerm); err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func writeJSON(filenam string, data interface{ ToJSON() ([]byte, error) }) {
	json, err := data.ToJSON()

	if err != nil {
		log.Fatalf(err.Error())
	}

	ioutil.WriteFile(output+filenam, json, fs.ModePerm)
}

func prepareOutputDir(year, month int) {
	year_str := template.GetYearLiteral(year)
	month_str := template.GetMonthLiteral(month)

	output = fmt.Sprintf("./data/%s/%d_%s/", year_str, month, month_str)

	if err := os.MkdirAll(output, os.ModePerm); err != nil {
		log.Fatalf(err.Error())
	}
}
