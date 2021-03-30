package bequant

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseUrl = "https://api.bequant.io"

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	url := baseUrl + "/api/2/public/ticker"

	data := utils.MakeRequest(url)

	for _, i := range data.([]interface{}) {
		d := i.(map[string]interface{})

		var bid float64
		var ask float64

		if d["bid"] != nil {
			bid, _ = strconv.ParseFloat(d["bid"].(string), 64)
		} else { bid = 0.0 }
		if d["ask"] != nil {
			ask, _ = strconv.ParseFloat(d["ask"].(string), 64)
		} else { ask = 0.0}
		
		baseVol, _ := strconv.ParseFloat(d["volume"].(string), 64)
		quoteVol, _ := strconv.ParseFloat(d["volumeQuote"].(string), 64)

		
		t := models.Ticker {
			Coin: strings.SplitN(d["symbol"].(string), "", 2)[0],
			Currency: strings.SplitN(d["symbol"].(string), "", 2)[1],
			Symbol: d["symbol"].(string),
			BidPrice: bid,
			BidQty: 0.0,
			AskPrice: ask,
			AskQty: 0.0,
			BaseVolume: baseVol,
			QuoteVolume: quoteVol,
			Exchange: "bequant",
			Timestamp: int(time.Now().Unix()),
		}

		tickers = append(tickers, t)
	}

	return tickers
}