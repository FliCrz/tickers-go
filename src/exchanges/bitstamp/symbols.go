package bitstamp

import (
	"strings"
)

func getCurrencies () (curList []string) {
	curList = append(curList, "eur", "usd", "gbp", "pax", "usdc", "btc", "eth")
	return curList
}

func getSymbols() (symbolsList []string) {
	symbolsString := "btcusd, btceur, btcgbp, btcpax, btcusdc, gbpusd, gbpeur, eurusd, xrpusd, xrpeur, xrpbtc, xrpgbp, xrppax, ltcusd, ltceur, ltcbtc, ltcgbp, ethusd, etheur, ethbtc, ethgbp, ethpax, ethusdc, bchusd, bcheur, bchbtc, bchgbp, paxusd, paxeur, paxgbp, xlmbtc, xlmusd, xlmeur, xlmgbp, linkusd, linkeur, linkgbp, linkbtc, linketh, omgusd, omgeur, omggbp, omgbtc, usdcusd, usdceur, daiusd, kncusd, knceur, kncbtc, mkrusd, mkreur, mkrbtc, zrxusd, zrxeur, zrxbtc, gusdusd"
	symbolsList = append(symbolsList, strings.Split(symbolsString, ", ")...)

	return symbolsList
}