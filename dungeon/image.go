package dungeon

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"jp/thelow/static/template"
	"log"
	"strconv"
	"strings"
	"time"
)

func CreateRankingImage(r *DungeonList, output string) {
	t := string(template.ReadTemplate(template.DUNGEON))
	t = strings.Replace(t, "{#month}", strconv.Itoa(r.Month), 1)
	t = strings.Replace(t, "{#month_literal}", time.Month(r.Month).String(), 1)

	total := fmt.Sprintf("%.2f", float64(r.Total)/1000.0)
	t = strings.Replace(t, "{#total}", total, 1)

	ranking := r.Ranking
	len := len(ranking)
	if len > 10 {
		len = 10
	}

	for i := 0; i < len; i++ {
		name := ranking[i].Name
		count := strconv.Itoa(ranking[i].Count)

		t = strings.Replace(t, fmt.Sprintf("{#dungeon_name_%d}", i), name, 1)
		t = strings.Replace(t, fmt.Sprintf("{#count_%d}", i), count, 1)
	}

	if err := ioutil.WriteFile(output+string(template.DUNGEON), []byte(t), fs.ModePerm); err != nil {
		log.Fatalf(err.Error())
	}
}
