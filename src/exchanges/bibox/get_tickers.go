package bibox

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

func getData (s string) models.Ticker {
	
	raw := utils.MakeRequest("https://api.bibox.com/v3/mdata/ticker?pair=" + s)

	data := raw.(map[string]interface{})["result"]

	bid := 0.0
	bidQty := 0.0
	ask := 0.0
	askQty := 0.0
	vol := 0.0
	
	if data.(map[string]interface{})["buy"] != nil {
		bid, _ = strconv.ParseFloat(data.(map[string]interface{})["buy"].(string), 64)
	}

	if data.(map[string]interface{})["buy_amount"] != nil {
		bidQty, _ = strconv.ParseFloat(data.(map[string]interface{})["buy_amount"].(string), 64)
	}

	if data.(map[string]interface{})["sell"] != nil {
		ask, _ = strconv.ParseFloat(data.(map[string]interface{})["sell"].(string), 64)
	}

	if data.(map[string]interface{})["sell_amount"] != nil {
		askQty, _ = strconv.ParseFloat(data.(map[string]interface{})["sell_amount"].(string), 64)
	}

	if data.(map[string]interface{})["vol"] != nil {
		vol, _ = strconv.ParseFloat(data.(map[string]interface{})["vol"].(string), 64)
	}
	
	t := models.Ticker {
		Coin: strings.Split(s, "_")[0],
		Currency: strings.Split(s, "_")[1],
		Symbol: strings.Replace(s, "_", "", 1),
		BidPrice: bid,
		BidQty: bidQty,
		AskPrice: ask,
		AskQty: askQty,
		BaseVolume: vol,
		QuoteVolume: 0.0,
		Exchange: "bibox",
		Timestamp: int(time.Now().Unix()),
	}

	return t
}

// GetTickrs ...
func GetTickers () (tickers []models.Ticker) {

	symbols := getSymbols()

	for _, s := range symbols {

		c := make(chan models.Ticker)

		go func () { c <- getData(s) } ()
		
		t := <- c

		tickers = append(tickers, t)
	}
	return tickers
}