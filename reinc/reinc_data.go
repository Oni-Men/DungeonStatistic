package reinc

type ReincData struct {
	MCID     string
	Counters map[ReincType]*ReincCounter
}

func NewReincData(mcid string) *ReincData {
	return &ReincData{
		MCID:     mcid,
		Counters: make(map[ReincType]*ReincCounter),
	}
}

func (d *ReincData) GetCount() int {
	count := 0

	for _, counter := range d.Counters {
		count += counter.Total
	}

	return count
}
