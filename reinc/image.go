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

	blacklist := map[string]bool{
		"7a44a801-652f-40fa-a65c-863c882935f9": true,
	}

	n := 0

	for i := 0; i < len; i++ {
		if _, ok := blacklist[ranking[i]]; ok {
			continue
		}

		name := r.GetMcid(ranking[i])
		count := strconv.Itoa(r.GetCount(ranking[i]))

		t = strings.Replace(t, fmt.Sprintf("{#mcid_%d}", n), name, 1)
		t = strings.Replace(t, fmt.Sprintf("{#count_%d}", n), count, 1)

		n++
		if n > 10 {
			break
		}
	}

	return t
}
