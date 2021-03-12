package utils

import (
	"tickers/src/libs/exchanges"
	"tickers/src/models"
)

// TickersFuncs ...
var TickersFuncs = map[string]func() []models.Ticker {
	"binance": exchanges.GetBinanceTickers,
	"bitfinex": exchanges.GetBitfinexTickers,
	"bittrex": exchanges.GetBittrexTickers,
	"btcpop": exchanges.GetBtcPopTickers,
	"huobi": exchanges.GetHuobiTickers,
	"kraken": exchanges.GetKrakenTickers,
	"kucoin": exchanges.GetKucoinTickers,
	"okex": exchanges.GetOkexTickers,
	"poloniex": exchanges.GetPoloniexTickers,
	"liquid": exchanges.GetLiquidTickers}