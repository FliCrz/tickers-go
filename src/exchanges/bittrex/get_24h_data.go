package bittrex

import (
	"strconv"
	"tickers/src/utils"
)

func get24hData () map[string][]float64 {
	url := "https://api.bittrex.com/v3/markets/summaries"

	data := utils.MakeRequest(url)

	parsed := data.([]interface{})

	// log.Println(parsed)

	var rsl = make(map[string][]float64)

	for _, i := range parsed {
		d := i.(map[string]interface{})
		baseVol, _ := strconv.ParseFloat(d["volume"].(string), 64)
		quoteVol, _ := strconv.ParseFloat(d["quoteVolume"].(string), 64)
		rsl[d["symbol"].(string)] = []float64{baseVol, quoteVol}
	}

	return rsl
}