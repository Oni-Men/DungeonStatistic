package reinc

import (
	"encoding/json"
	"jp/thelow/static/model"
	"sort"
	"time"
)

type ReincResult struct {
	model.Result
	total int
	IDs   map[string]bool
	datas map[string]*ReincData
}

func NewResult(year, month int) *ReincResult {
	r := &ReincResult{
		total: 0,
		IDs:   make(map[string]bool),
		datas: make(map[string]*ReincData),
	}

	r.SetYear(year)
	r.SetMonth(time.Month(month))

	return r
}

func (r *ReincResult) GetTotal() int {
	return r.total
}

func (r *ReincResult) Increment(uuid, mcid, spanID string, reincType ReincType, when time.Time, count int) {
	if r.checkForDuplicates(spanID) {
		return
	}

	data, ok := r.datas[uuid]
	if !ok {
		data = NewReincData(mcid)
		r.datas[uuid] = data
	}

	counter, ok := data.Counters[reincType]
	if !ok {
		counter = NewReincCounter(r.GetYear(), r.GetMonthInt())
		data.Counters[reincType] = counter
	}

	counter.Count(count, when)
	r.total++
}

func (r *ReincResult) GetCount(uuid string) int {
	data, ok := r.datas[uuid]

	if ok {
		count := 0

		for _, counter := range data.Counters {
			count += counter.Total
		}

		return count
	}

	return -1
}

func (r *ReincResult) GetMcid(uuid string) string {
	data, ok := r.datas[uuid]

	if !ok {
		return ""
	}

	return data.MCID
}

func (r *ReincResult) ToJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Total int
		Datas map[string]*ReincData
	}{
		Total: r.total,
		Datas: r.datas,
	})
}

func (r *ReincResult) CreateRanking() []string {
	ranking := make([]string, 0, len(r.datas))
	for k := range r.datas {
		ranking = append(ranking, k)
	}

	sort.Slice(ranking, func(i, j int) bool {
		a, ok := r.datas[ranking[i]]
		if !ok {
			a = nil
		}

		b, ok := r.datas[ranking[j]]
		if !ok {
			b = nil
		}

		if a == nil {
			return false
		}

		if b == nil {
			return true
		}

		return a.GetCount() > b.GetCount()
	})

	return ranking
}

// Returns true if duplicates has exist.
func (r *ReincResult) checkForDuplicates(id string) bool {
	_, ok := r.IDs[id]
	if ok {
		return true
	}
	r.IDs[id] = true

	return false
}