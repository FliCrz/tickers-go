package binance

import (
	"strconv"
	"tickers/src/utils"
)

func get24HData () map[string][]float64 {
	url := APIURL + "/ticker/24hr"
	data := utils.MakeRequest(url)

	var rsl = make(map[string][]float64)

	for _, i := range data.([]interface{}) {
		d := i.(map[string]interface{})
		baseVol, _ := strconv.ParseFloat(d["volume"].(string), 64)
		quoteVol, _ := strconv.ParseFloat(d["quoteVolume"].(string), 64)
		rsl[d["symbol"].(string)] = []float64{baseVol, quoteVol}
	}
	return rsl
}