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
	"jp/thelow/static/model"
	"jp/thelow/static/template"
)

var (
	config = new(model.Config)
	year   *int
	month  *int
	output = "."
)

func main() {
	now := time.Now()
	year = flag.Int("year", now.Year(), "集計対象の年(デフォルトは今年)")
	month = flag.Int("month", int(now.Month()), "集計対象の月(デフォルトは今月)")
	flag.Parse()

	model.LoadConfig(config)
	prepareOutputDir(*year, *month)

	fetchAll()
	//createImages()
}

// func createImages() {

// 	gen := func(fileName string, executor func(model.Config, int, int, string) string) {
// 		data, err := ioutil.ReadFile("./template/" + fileName)
// 		if err != nil {
// 			log.Fatalf(err.Error())
// 		}
// 		text := executor(*config, year, month, string(data))
// 		if err = ioutil.WriteFile(output+fileName, []byte(text), fs.ModePerm); err != nil {
// 			log.Fatalf(err.Error())
// 		}
// 	}

// 	gen("dungeon.svg", template.GenImgCompletedDungeons)
// 	gen("reincarnation.svg", template.GenImgReincarnations)

// }

func fetchAll() {
	result := dungeon.FetchCompletions(config.Host, *year, *month)
	data, err := result.ToJSON()
	if err != nil {
		log.Fatalf(err.Error())
	}
	ioutil.WriteFile(output+"completes.json", data, fs.ModePerm)
}

func prepareOutputDir(year, month int) {
	year_str := template.GetYearLiteral(year)
	month_str := template.GetMonthLiteral(month)

	output = fmt.Sprintf("./data/%s/%s/", year_str, month_str)

	if err := os.MkdirAll(output, os.ModePerm); err != nil {
		log.Fatalf(err.Error())
	}
}
