package model

import "time"

type Result struct {
	total int
	year  int
	month time.Month
}

func (r *Result) IncrementTotal() {
	r.total++
}

func (r *Result) GetTotal() int {
	return r.total
}

func (r *Result) GetYear() int {
	return r.year
}

func (r *Result) SetYear(year int) {
	r.year = year
}

func (r *Result) GetMonth() time.Month {
	return r.month
}

func (r *Result) SetMonth(month time.Month) {
	r.month = month
}

func (r *Result) GetMonthInt() int {
	return int(r.GetMonth())
}

func (r *Result) GetMonthLiteral() string {
	return r.month.String()
}
