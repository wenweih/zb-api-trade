package main

import (
	"encoding/json"

	"github.com/go-resty/resty"
	"github.com/mitchellh/mapstructure"
)

var apiClient, tradeClient *resty.Client

func init() {
	apiClient = resty.New().SetDebug(false).SetHostURL(config.dataURL)
	tradeClient = resty.New().SetDebug(false).SetHostURL(config.tradeURL)

	handleQueryParams(tradeClient)
	handleQueryParams(apiClient)
}

func depth(api, market, size string) *respDepth {
	resp, _ := apiClient.SetQueryParams(map[string]string{
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

func trades(api, market string) *respTrades {
	resp, _ := apiClient.SetQueryParams(map[string]string{
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
