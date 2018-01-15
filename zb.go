package main

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

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

func getTicker(api, market string) *respTicker {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
	}).R().Get(api)
	var res respTicker
	json.Unmarshal(resp.Body(), &res)
	return &res
}
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
func trades(api, market string) *respTrades {
	resp, _ := dataClient.SetQueryParams(map[string]string{
		"market": market,
	}).R().Get(api)
	var res respTrades
	json.Unmarshal(resp.Body(), &res)
	return &res
}

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

// tradeType 交易类型1/0[buy/sell]
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

func cancelOrder(api, id, currency, sign string) bool {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"currency": currency,
		"method":   "cancelOrder",
		"id":       id,
		"sign":     sign,
	}).R().Get(api)
	return simpleResponse(resp)
}
