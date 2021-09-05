package template

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"time"
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

	ErrMonthOutOfRange = errors.New("specified month is out of range 1-12")
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

func GetMonthLiteral(month int) (string, error) {
	if month < 1 || 12 < month {
		return "", ErrMonthOutOfRange
	}

	return time.Month(month).String(), nil
}
