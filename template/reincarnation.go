package template

import (
	"jp/thelow/static/model"
)

func GenImgReincarnations(cfg model.Config, year, month int, text string) string {
	// ctx := context.Background()
	// client := sheet.NewSheetClient(ctx, cfg.SheetID)

	// values, err := client.Get(getSheetName(year, month) + "!D1:E11")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// literal := GetMonthLiteral(month)

	// text = strings.Replace(text, "{#month}", strconv.Itoa(month), 1)
	// text = strings.Replace(text, "{#month_literal}", literal, 1)

	// for i := 0; i < len(values); i++ {
	// 	mcid := values[i][0]
	// 	count := values[i][1]

	// 	text = strings.Replace(text, fmt.Sprintf("{#mcid_%d}", i-1), mcid.(string), 1)
	// 	text = strings.Replace(text, fmt.Sprintf("{#count_%d}", i-1), count.(string), 1)
	// }

	return text
}
