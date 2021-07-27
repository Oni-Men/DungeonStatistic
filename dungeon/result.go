package dungeon

import "encoding/json"

type DungeonResult struct {
	total     int
	completes map[string]int
}

func NewResult() *DungeonResult {
	return &DungeonResult{
		total:     0,
		completes: make(map[string]int),
	}
}

func (r *DungeonResult) Increment(name string) {
	r.completes[name]++
	r.total++
}

func (r *DungeonResult) GetCount(name string) int {
	count, ok := r.completes[name]

	if ok {
		return count
	}

	return -1
}

func (r *DungeonResult) GetTotal() int {
	return r.total
}

func (r *DungeonResult) ToJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Total     int
		Completes map[string]int
	}{
		Total:     r.total,
		Completes: r.completes,
	})
}
