package zbg

import (
	"strconv"
	"tickers/src/utils"
)

func getSymbols () (map[int64]string) {
	url := baseUrl + "/exchange/api/v1/common/symbols"

	data := utils.MakeRequest(url)

	parsed := data.(map[string]interface{})["datas"]

	symbols := make(map[int64]string)

	for _, i := range parsed.([]interface{}) {
		d := i.(map[string]interface{})
		_id, _ := strconv.ParseInt(d["id"].(string), 0, 64)
		symbols[_id] = d["symbol"].(string)
	}

	return symbols
}