package dungeon

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateImage(r *DungeonResult, t string) string {

	t = strings.Replace(t, "{#month}", strconv.Itoa(r.GetMonthInt()), 1)
	t = strings.Replace(t, "{#month_literal}", r.GetMonthLiteral(), 1)

	total := fmt.Sprintf("%.2f", float64(r.GetTotal())/1000.0)
	t = strings.Replace(t, "{#total}", total, 1)

	ranking := r.CreateRanking()
	len := len(ranking)
	if len > 10 {
		len = 10
	}

	for i := 0; i < len; i++ {
		name := ranking[i]
		count := strconv.Itoa(r.GetCount(name))

		t = strings.Replace(t, fmt.Sprintf("{#dungeon_name_%d}", i), name, 1)
		t = strings.Replace(t, fmt.Sprintf("{#count_%d}", i), count, 1)
	}

	return t
}
