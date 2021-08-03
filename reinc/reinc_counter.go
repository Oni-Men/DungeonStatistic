package reinc

import (
	"jp/thelow/static/fetch"
	"time"
)

type ReincCounter struct {
	CountLatest int
	CountOldest int
	Latest      time.Time
	Oldest      time.Time
	Total       int
}

func NewReincCounter(year, month int) *ReincCounter {
	data := &ReincCounter{
		Latest: fetch.StartOfMonth(year, month),
		Oldest: fetch.EndOfMonth(year, month),
	}

	return data
}

func (c *ReincCounter) Count(count int, when time.Time) {

	if when.After(c.Latest) {
		c.Latest = when
		c.CountLatest = count
	}

	if when.Before(c.Oldest) {
		c.Oldest = when
		c.CountOldest = count
	}

	if c.Latest.Equal(c.Oldest) {
		c.Total = 1
	} else {
		c.Total = c.CountLatest - c.CountOldest
	}

}
