package dungeon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jp/thelow/static/model"
	"log"
	"sort"
	"time"
)

type DungeonResult struct {
	model.Result
	Completes map[string]int
}

func NewResult(year, month int) *DungeonResult {
	r := &DungeonResult{
		Completes: make(map[string]int),
	}

	r.SetYear(year)
	r.SetMonth(time.Month(month))

	return r
}

func FromFile(year, month int) *DungeonResult {
	filename := fmt.Sprintf("./data/%d/%s/completions.json", year, time.Month(month).String())
	read, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf(err.Error())
	}

	r := new(DungeonResult)
	if err := json.Unmarshal(read, r); err != nil {
		log.Fatalf(err.Error())
	}

	return r
}

func (r *DungeonResult) Increment(name string) {
	r.Completes[name]++
	r.IncrementTotal()
}

func (r *DungeonResult) GetCount(name string) int {
	count, ok := r.Completes[name]

	if ok {
		return count
	}

	return -1
}

func (r *DungeonResult) ToJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Total     int
		Completes map[string]int
	}{
		Total:     r.GetTotal(),
		Completes: r.Completes,
	})
}

func (r *DungeonResult) CreateRanking() []string {
	ranking := make([]string, 0, len(r.Completes))
	for k := range r.Completes {
		ranking = append(ranking, k)
	}

	sort.Slice(ranking, func(i, j int) bool {
		a, ok := r.Completes[ranking[i]]
		if !ok {
			a = -1
		}

		b, ok := r.Completes[ranking[j]]
		if !ok {
			b = -1
		}

		return a > b
	})

	return ranking
}
