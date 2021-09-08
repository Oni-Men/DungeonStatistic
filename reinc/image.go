package reinc

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"jp/thelow/static/template"
	"log"
	"strconv"
	"strings"
)

func CreateRankingImage(r *ReincList, output string) {
	t := string(template.ReadTemplate(template.REINCARNATION))
	t = strings.Replace(t, "{#month}", strconv.Itoa(int(r.Month)), 1)
	t = strings.Replace(t, "{#month_literal}", r.Month.String(), 1)

	n := 0

	for _, name := range r.Ranking {
		count := r.Data[name]

		t = strings.Replace(t, fmt.Sprintf("{#mcid_%d}", n), name, 1)
		t = strings.Replace(t, fmt.Sprintf("{#count_%d}", n), strconv.Itoa(count), 1)

		n++
		if n > 10 {
			break
		}
	}

	svgFile := output + string(template.REINCARNATION)

	if err := ioutil.WriteFile(svgFile, []byte(t), fs.ModePerm); err != nil {
		log.Fatalf(err.Error())
	}
}
