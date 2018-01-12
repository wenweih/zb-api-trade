package main

import (
	"encoding/json"

	"github.com/go-resty/resty"
	"github.com/mitchellh/mapstructure"
)

//==================================================//
type depthOrder []float64
type respDepth struct {
	Timestamp float64      `mapstructure:"timestamp"`
	Asks      []depthOrder `mapstructure:"asks"`
	Bids      []depthOrder `mapstructure:"bids"`
}

//==================================================//
type respTrades []trade
type trade struct {
	Amount    string `json:"amount"`
	Date      int64  `json:"date"`
	Price     string `json:"price"`
	Tid       int64  `json:"tid"`
	TradeType string `json:"trade_type"`
	TxType    string `json:"type"`
}

//==================================================//
type accountInfoCoin struct {
	EnName        string `mapstructure:"enName"`
	Freez         string `mapstructure:"freez"`
	UnitDecimal   string `mapstructure:"unitDecimal"`
	CnName        string `mapstructure:"cnName"`
	IsCanRecharge bool   `mapstructure:"isCanRecharge"`
	UnitTag       string `mapstructure:"unitTag"`
	IsCanWithdraw bool   `mapstructure:"isCanWithdraw"`
	Available     string `mapstructure:"available"`
	Key           string `mapstructure:"key"`
}
type accountInfoBase struct {
	Username             string `mapstructure:"username"`
	TradePasswordEnabled bool   `mapstructure:"trade_password_enabled"`
	AuthGoogleEnabled    bool   `mapstructure:"auth_google_enabled"`
	AuthMobileEnabled    bool   `mapstructure:"auth_mobile_enabled"`
}
type accountInfoResult struct {
	Coins    []accountInfoCoin `mapstructure:"coins"`
	BaseInfo accountInfoBase   `mapstructure:"base"`
}
type respAccountInfo struct {
	Result accountInfoResult `mapstructure:"result"`
}

//==================================================//
type respOrders []order
type order struct {
	Currency    string  `json:"currency"`
	ID          string  `json:"id"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	TradeAmount float64 `json:"trade_amount"`
	TradeDate   int64   `json:"trade_date"`
	TradeMoney  string  `json:"trade_money"`
	TradePrice  float64 `json:"trade_price"`
	OrderType   int     `json:"type"`
}

//==================================================//
type respSimple struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//==================================================//
var apiClient, tradeClient *resty.Client

func init() {
	apiClient = resty.New().SetDebug(false).SetHostURL(config.dataURL)
	tradeClient = resty.New().SetDebug(false).SetHostURL(config.tradeURL)

	handleQueryParams(tradeClient)
	handleQueryParams(apiClient)
}

func depth(api string, market string, size string) *respDepth {
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

func trades(api string, market string) *respTrades {
	resp, _ := apiClient.SetQueryParams(map[string]string{
		"market": market,
	}).R().Get(api)
	var res respTrades
	json.Unmarshal(resp.Body(), &res)
	return &res
}

func accountInfo(api string, sign string) *respAccountInfo {
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
func getOrders(api string, currency string, tradeType string, sign string) *respOrders {
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

func cancelOrder(api string, id string, currency string, sign string) bool {
	resp, _ := tradeClient.SetQueryParams(map[string]string{
		"currency": currency,
		"method":   "cancelOrder",
		"id":       id,
		"sign":     sign,
	}).R().Get(api)
	return simpleResponse(resp)
}
