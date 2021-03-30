package zbg

import (
	"strconv"
	"strings"
	"tickers/src/models"
	"tickers/src/utils"
	"time"
)

// GetTickers ...
func GetTickers () (tickers []models.Ticker) {
	
	url := "https://kline.zbg.com/api/data/v1/tickers"

	data := utils.MakeRequest(url)

	parsed := data.(map[string]interface{})["datas"]

	symbols := getSymbols()

	// log.Println(symbols)

	for _, i := range parsed.([]interface{}) {
		d := i.([]interface{})

		rawSymbol, _ := strconv.ParseInt(d[0].(string), 0, 64)
		coin := strings.ToUpper(strings.Split(symbols[rawSymbol], "_")[0])
		
		if coin != "" {
			bid, _ := strconv.ParseFloat(d[7].(string), 64)
			ask, _ := strconv.ParseFloat(d[8].(string), 64)
			
			t := models.Ticker{
				Coin: coin,
				Currency: strings.ToUpper(strings.Split(symbols[rawSymbol], "_")[1]),
				Symbol: strings.ToUpper(strings.Replace(symbols[rawSymbol], "_", "", 1)),
				BidPrice: bid,
				BidQty: 0.0,
				AskPrice: ask,
				AskQty: 0.0,
				Exchange: "zbg",
				Timestamp: int(time.Now().Unix()),
			}
			tickers = append(tickers, t)
		}
	}

	return tickers
}