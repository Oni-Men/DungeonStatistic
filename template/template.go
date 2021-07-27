package template

import (
	"fmt"
	"strconv"
)

var monthLiterals = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func GetYearLiteral(year int) string {
	return strconv.Itoa(year)
}

func GetMonthLiteral(month int) string {
	return monthLiterals[month-1]
}

func getSheetName(year, month int) string {
	return fmt.Sprintf("'%d/%d月'", year, month)
}
