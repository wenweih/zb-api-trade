package main

import (
	"encoding/json"
	"strconv"
	"time"

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
var apiClient, tradeClient *resty.Client

func init() {
	apiClient = resty.New().SetDebug(false).
		SetHostURL(config.dataURL).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})

	tradeClient = resty.New().SetDebug(false).
		SetHostURL(config.tradeURL).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		})
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
		"accesskey": config.accesstoken,
		"method":    "getAccountInfo",
		"sign":      sign,
		"reqTime":   strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
	}).R().Get(api)

	var (
		res          respAccountInfo
		mapInterface map[string]interface{}
	)
	json.Unmarshal(resp.Body(), &mapInterface)
	mapstructure.Decode(mapInterface, &res)
	return &res
}
