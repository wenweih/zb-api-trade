package main

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// 市场深度
// depth("depth", "btc_usdt", "20")
func depth(api, market, size string) *respDepth {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
		"size":   size,
	}).R().Get(api)

	var (
		res          respDepth
		mapInterface map[string]interface{}
	)
	json.Unmarshal(resp.Body(), &mapInterface)
	mapstructure.Decode(mapInterface, &res)
	return &res
}

// 行情
// getTicker("ticker", "btc_usdt")
func getTicker(api, market string) *respTicker {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
	}).R().Get(api)
	var res respTicker
	json.Unmarshal(resp.Body(), &res)
	return &res
}

// K线
// kline("kline", "btc_usdt", "1min", "10")
func kline(api, market, timeType, size string) *respKline {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
		"type":   timeType,
		"size":   size,
	}).R().Get(api)
	var (
		res          respKline
		mapInterface map[string]interface{}
	)
	json.Unmarshal(resp.Body(), &res)
	mapstructure.Decode(mapInterface, &res)
	return &res
}

// 历史成交
// trades("trades", "btc_usdt")
func trades(api, market string) *respTrades {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
	}).R().Get(api)
	var res respTrades
	json.Unmarshal(resp.Body(), &res)
	return &res
}

// 获取用户信息
// params := map[string]string{
// 	"accesskey": config.accesstoken,
// 	"method":    "getAccountInfo",
// }
// sorted := sortParams(params)
// sign := hmacSign(sorted)
// accountInfo("getAccountInfo", sign)
func accountInfo(api, sign string) *respAccountInfo {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"method": "getAccountInfo",
		"sign":   sign,
	}).R().Get(api)
	var (
		res          respAccountInfo
		mapInterface map[string]interface{}
	)
	json.Unmarshal(resp.Body(), &mapInterface)
	mapstructure.Decode(mapInterface, &res)
	return &res
}

// 委托下单
// createOrderParams := map[string]string{
// 	"accesskey": config.accesstoken,
// 	"amount":    "1",
// 	"currency":  "bts_usdt",
// 	"price":     "20",
// 	"tradeType": "0",
// 	"method":    "order",
// }
// createOrderSorted := sortParams(createOrderParams)
// createOrderSign := hmacSign(createOrderSorted)
// createOrder("order", "1", "bts_usdt", "0", "20", createOrderSign)
func createOrder(api, amount, currency, tradeType, price, sign string) bool {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"amount":    amount,
		"currency":  currency,
		"method":    "order",
		"price":     price,
		"tradeType": tradeType,
		"sign":      sign,
	}).R().Get(api)
	return simpleResponse(resp)
}

// 获取委托买单或卖单
// tradeType 交易类型1/0[buy/sell]

// orderParams := map[string]string{
// 	"accesskey": config.accesstoken,
// 	"currency":  "bts_usdt",
// 	"method":    "getOrdersNew",
// 	"pageIndex": "1",
// 	"pageSize":  "50",
// 	"tradeType": "0",
// }
// orderSorted := sortParams(orderParams)
// orderSign := hmacSign(orderSorted)
// getOrders("getOrdersNew", "bts_usdt", "0", orderSign)
func getOrders(api, currency, tradeType, sign string) *respOrders {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"currency":  currency,
		"method":    "getOrdersNew",
		"pageIndex": "1",
		"pageSize":  "50",
		"tradeType": tradeType,
		"sign":      sign,
	}).R().Get(api)
	var res respOrders
	json.Unmarshal(resp.Body(), &res)
	return &res
}

// 取消委托
// id := (*resss)[0].ID
// cancelParams := map[string]string{
// 	"accesskey": config.accesstoken,
// 	"currency":  "bts_usdt",
// 	"id":        id,
// 	"method":    "cancelOrder",
// }
// cancelSorted := sortParams(cancelParams)
// cancelSign := hmacSign(cancelSorted)
// cancelOrder("cancelOrder", id, "bts_usdt", cancelSign)
func cancelOrder(api, id, currency, sign string) bool {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"currency": currency,
		"method":   "cancelOrder",
		"id":       id,
		"sign":     sign,
	}).R().Get(api)
	return simpleResponse(resp)
}
