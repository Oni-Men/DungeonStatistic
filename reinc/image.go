package reinc

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateImage(r *ReincResult, t string) string {

	t = strings.Replace(t, "{#month}", strconv.Itoa(r.GetMonthInt()), 1)
	t = strings.Replace(t, "{#month_literal}", r.GetMonthLiteral(), 1)

	ranking := r.CreateRanking()
	len := len(ranking)
	if len > 10 {
		len = 10
	}

	for i := 0; i < len; i++ {
		name := r.GetMcid(ranking[i])
		count := strconv.Itoa(r.GetCount(ranking[i]))

		t = strings.Replace(t, fmt.Sprintf("{#mcid_%d}", i), name, 1)
		t = strings.Replace(t, fmt.Sprintf("{#count_%d}", i), count, 1)
	}

	return t
}
