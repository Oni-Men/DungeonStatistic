package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"jp/thelow/static/dungeon"
	"jp/thelow/static/model"
	"jp/thelow/static/player"
	"jp/thelow/static/reinc"
	"jp/thelow/static/template"
)

var (
	config = new(model.Config)
	mcid   *string
	image  *bool
	output = "."
)

func main() {
	loadConfig()

	year := flag.Int("year", config.Year, "集計対象の年(デフォルトは今年)")
	month := flag.Int("month", config.Month, "集計対象の月(デフォルトは今月)")
	mcid = flag.String("mcid", "", "mcid")
	image = flag.Bool("image", false, "テンプレートからSVGを生成するか")

	flag.Parse()

	config.Year = *year
	config.Month = *month

	if config.Year < 0 {
		log.Printf("specified year is less than 0: %d\n", config.Year)
	}

	if config.Month < 1 || 12 < config.Month {
		log.Printf("specified month is out of range: %d\n", config.Month)
		os.Exit(1)
	}

	createOutputDir(config.Year, config.Month)

	if *mcid != "" {
		fmt.Printf("Fetch %s's data...\n", *mcid)
		player.FetchPlayersDungeonData(config.Host, *mcid, "クラバスタ")
	} else {
		clears, err := dungeon.Count(config)
		if err == nil {
			writeJSON("dungeon_clears.json", clears)
		} else {
			log.Fatalf(err.Error())
		}

		reincs, err := reinc.Count(config)

		if err == nil {
			writeJSON("reincs.json", reincs)
		} else {
			log.Fatalf(err.Error())
		}

		if *image {
			dungeon.CreateRankingImage(clears, output)
			reinc.CreateRankingImage(reincs, output)
		}

	}
}

func loadConfig() {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = json.Unmarshal(data, config); err != nil {
		log.Fatalf(err.Error())
	}
}

func writeJSON(filenam string, data interface{}) {
	json, err := json.Marshal(data)

	if err != nil {
		log.Fatalf(err.Error())
	}

	ioutil.WriteFile(output+filenam, json, fs.ModePerm)
}

func createOutputDir(year, month int) error {
	year_lit := template.GetYearLiteral(year)
	month_lit, err := template.GetMonthLiteral(month)
	if err != nil {
		return err
	}

	output = fmt.Sprintf("./data/%s/%d_%s/", year_lit, month, month_lit)
	return os.MkdirAll(output, os.ModePerm)
}
