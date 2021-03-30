package okex

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

var baseURL = "https://www.okex.com"

// GetTickers ...
func GetTickers () []models.Ticker {
	
	var tickers []models.Ticker

	url := baseURL + "/api/spot/v3/instruments/ticker"

	data := utils.MakeRequest(url)

	parsed := data.([]interface{})

	var coin string
	var cur string

	for n := range parsed {
		d := parsed[n].(map[string]interface{})
		coin = strings.Split(d["product_id"].(string), "-")[0]
		cur = strings.Split(d["product_id"].(string), "-")[1]
		bp, _ := strconv.ParseFloat(d["best_bid"].(string), 64)
		bps, _ := strconv.ParseFloat(d["best_bid_size"].(string), 64)
		ap, _ := strconv.ParseFloat(d["best_ask"].(string), 64)
		aps, _ := strconv.ParseFloat(d["best_ask_size"].(string), 64)
		
		baseVol := d["base_volume_24"]
		if baseVol != nil {
			baseVol, _ = strconv.ParseFloat(baseVol.(string), 64)
		} else {baseVol = 0.0}

		quoteVol := d["quote_volume_24h"]
		if quoteVol != nil {
			quoteVol, _ = strconv.ParseFloat(quoteVol.(string), 64)
		} else { quoteVol = 0.0}

		tickers = append(tickers, models.Ticker{
			Coin: coin,
			Currency: cur,
			Symbol: coin + cur,
			BidPrice: bp,
			BidQty: bps,
			AskPrice: ap,
			AskQty: aps,
			BaseVolume: baseVol.(float64),
			QuoteVolume: quoteVol.(float64),
			Exchange: "okex",
			Timestamp: int(time.Now().Unix())})
	}
	return tickers
}