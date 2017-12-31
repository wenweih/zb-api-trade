package main

import (
	"encoding/json"

	"github.com/go-resty/resty"
	"github.com/mitchellh/mapstructure"
)

type order []float64

// RespDepth 深度接口返回结构体
type RespDepth struct {
	Timestamp float64 `mapstructure:"timestamp"`
	Asks      []order `mapstructure:"asks"`
	Bids      []order `mapstructure:"bids"`
}

var (
	apiClient, tradeClient *resty.Client
)

func init() {
	apiClient = resty.New().SetDebug(true).
		SetHostURL("http://api.zb.com/data/v1/").
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "https://huangwenwei.com",
		})
}

func depth(api string, market string, size string) *RespDepth {
	resp, _ := apiClient.SetQueryParams(map[string]string{
		"market": market,
		"size":   size,
	}).R().Get(api)

	var (
		res RespDepth
		f   interface{}
	)
	json.Unmarshal(resp.Body(), &f)
	mapstructure.Decode(f.(map[string]interface{}), &res)
	return &res
}
