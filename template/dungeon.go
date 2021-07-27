package template

import (
	"jp/thelow/static/model"
)

func GenImgCompletedDungeons(cfg model.Config, year, month int, text string) string {
	// ctx := context.Background()
	// client := sheet.NewSheetClient(ctx, cfg.SheetID)

	// values, err := client.Get(getSheetName(year, month) + "!A1:B11")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// literal := GetMonthLiteral(month)

	// text = strings.Replace(text, "{#month}", strconv.Itoa(month), 1)
	// text = strings.Replace(text, "{#month_literal}", literal, 1)

	// f, err := strconv.ParseFloat(values[0][1].(string), 64)
	// if err != nil {
	// 	f = 0.0
	// }

	// total := fmt.Sprintf("%.2f", f/1000.0)
	// text = strings.Replace(text, "{#total}", total, 1)

	// for i := 1; i < len(values); i++ {
	// 	dungeonName := values[i][0]
	// 	dungeonCompletes := values[i][1]

	// 	text = strings.Replace(text, fmt.Sprintf("{#dungeon_name_%d}", i-1), dungeonName.(string), 1)
	// 	text = strings.Replace(text, fmt.Sprintf("{#count_%d}", i-1), dungeonCompletes.(string), 1)
	// }

	return text
}
