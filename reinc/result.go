package reinc

import (
	"time"
)

type ReincList struct {
	Data    map[string]int
	Ranking []string
	Year    int
	Month   time.Month
}
