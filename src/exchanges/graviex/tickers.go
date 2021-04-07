package graviex

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {

	raw := utils.MakeRequest(baseUrl + "/tickers.json")

	data := raw.(map[string]interface{}) 

	for _, i := range data {
		d := i.(map[string]interface{})

		coin := strings.ToUpper(d["base_unit"].(string))
		cur := strings.ToUpper(d["quote_unit"].(string))
		bid, _ := strconv.ParseFloat(d["buy"].(string), 64)
		ask, _ := strconv.ParseFloat(d["sell"].(string), 64)
		baseVol, _ := strconv.ParseFloat(d["volume"].(string), 64)
		quoteVol, _ := strconv.ParseFloat(d["volume2"].(string), 64)

		t := models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bid,
			BidQty: 0.0,
			AskPrice: ask,
			AskQty: 0.0,
			BaseVolume: baseVol,
			QuoteVolume: quoteVol,
			Exchange: "graviex",
			Timestamp: int(time.Now().Unix()),
		}

		tickers = append(tickers, t)
	}

	return tickers
}