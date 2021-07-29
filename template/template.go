package template

import (
	"io/ioutil"
	"log"
	"strconv"
)

type Template string

var (
	monthLiterals = []string{
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
	DUNGEON       Template = "dungeon.svg"
	REINCARNATION Template = "reincarnation.svg"
)

func ReadTemplate(t Template) string {
	read, err := ioutil.ReadFile(GetTemplateLocation(t))

	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(read)
}

func GetTemplateLocation(t Template) string {
	return "./template/" + string(t)
}

func GetYearLiteral(year int) string {
	return strconv.Itoa(year)
}

func GetMonthLiteral(month int) string {
	return monthLiterals[month-1]
}
