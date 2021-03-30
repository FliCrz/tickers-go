package stakecube

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	baseUrl := "https://stakecube.io/api/v2/exchange/spot/markets?baseMarket="
	
	currencies := []string{
		"BTC", "SCC", "LTC", "DASH", "DOGE",
	}

	for _, c := range currencies {
		url := baseUrl + c

		data := utils.MakeRequest(url)

		// log.Println(data)

		parsed := data.(map[string]interface{})["result"]

		reparsed := parsed.(map[string]interface{})

		for k, v := range reparsed {
			
			coin := strings.Split(k, "_")[0]
			cur := strings.Split(k, "_")[1]

			d := v.(map[string]interface{})

			baseVol, ok := d["volumeBase24h"].(float64)
			 if !ok {
				 baseVol, _ = strconv.ParseFloat(d["volumeBase24h"].(string), 64)
			 }
			 
			quoteVol, ok := d["volumeTrade24h"].(float64)
			if !ok {
				quoteVol, _ = strconv.ParseFloat(d["volumeTrade24h"].(string), 64)
			}

			bid, ok := d["bestBid"].(float64)
			if !ok {
				quoteVol, _ = strconv.ParseFloat(d["bestBid"].(string), 64)
			}

			ask, ok := d["bestAsk"].(float64)
			if !ok {
				quoteVol, _ = strconv.ParseFloat(d["bestAsk"].(string), 64)
			}

			new_ticker := models.Ticker {
				Coin: coin,
				Currency: cur,
				Symbol: coin + cur,
				BidPrice: bid,
				BidQty: 0.0,
				AskPrice: ask,
				AskQty: 0.0,
				BaseVolume: baseVol,
				QuoteVolume: quoteVol,
				Exchange: "stakecube",
				Timestamp: int(time.Now().Unix()),
			}
			
			// log.Println(new_ticker)
			tickers = append(tickers, new_ticker)
		}
	}
	
	return tickers
}