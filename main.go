package main

import "fmt"

func init() {
	cmdExecute()
}
func main() {
	params := map[string]string{
		"accesskey": config.accesstoken,
		"method":    "getAccountInfo",
	}
	sorted := sortParams(params)
	sign := hmacSign(sorted)
	accountInfo("getAccountInfo", sign)

	createOrderParams := map[string]string{
		"accesskey": config.accesstoken,
		"amount":    "1",
		"currency":  "bts_usdt",
		"price":     "20",
		"tradeType": "0",
		"method":    "order",
	}

	createOrderSorted := sortParams(createOrderParams)
	createOrderSign := hmacSign(createOrderSorted)
	crea := createOrder("order", "1", "bts_usdt", "0", "20", createOrderSign)
	fmt.Println(crea)

	orderParams := map[string]string{
		"accesskey": config.accesstoken,
		"currency":  "bts_usdt",
		"method":    "getOrdersNew",
		"pageIndex": "1",
		"pageSize":  "50",
		"tradeType": "0",
	}
	orderSorted := sortParams(orderParams)
	orderSign := hmacSign(orderSorted)
	resss := getOrders("getOrdersNew", "bts_usdt", "0", orderSign)
	fmt.Println(resss)

	id := (*resss)[0].ID
	cancelParams := map[string]string{
		"accesskey": config.accesstoken,
		"currency":  "bts_usdt",
		"id":        id,
		"method":    "cancelOrder",
	}
	cancelSorted := sortParams(cancelParams)
	cancelSign := hmacSign(cancelSorted)
	resp := cancelOrder("cancelOrder", id, "bts_usdt", cancelSign)
	fmt.Println(resp)

	depth("depth", "ltc_usdt", "20")

	trades("trades", "btc_usdt")
	kline("kline", "btc_usdt", "1min", "10")
	tick := getTicker("ticker", "btc_usdt")
	fmt.Println(tick.Ticker.Last)
}
