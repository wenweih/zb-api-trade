// Package cmd 命令行
package main

import (
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/go-resty/resty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configure struct {
	secretkey   string
	accesstoken string
	dataURL     string
	tradeURL    string
}

// Config 中币接口配置信息
var (
	db                      *bolt.DB
	cfgFile                 string
	config                  configure
	concurrency             *receiver
	dataClient, tradeClient httpClient

	rootCmd = &cobra.Command{
		Use:   "zb_trade",
		Short: "A tool for earning money",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start to trade...",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			concurrency.Add(1)
			go writeTaskRun(concurrency)
			time.Sleep(3)
			concurrency.Add(1)
			go readTaskRun(concurrency)
		},
	}
)

// Execute 命令行入口
func cmdExecute(r *receiver) {
	concurrency = r
	if err := rootCmd.Execute(); err != nil {
		Exit(err.Error())
	}
}

func init() {
	db = boltDB()
	rootCmd.AddCommand(startCmd)
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/zb.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigType("yaml")
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			Exit("Error: file no exist:", cfgFile)
		}

		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(HomeDir())
		viper.SetConfigName("zb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		StrLogger("Using Configure file:", viper.ConfigFileUsed())
	} else {
		Exit("Error: zb.yml not found in: ", HomeDir())
	}

	for key, value := range viper.AllSettings() {
		switch key {
		case "secretkey":
			config.secretkey = value.(string)
		case "accesstoken":
			config.accesstoken = value.(string)
		case "data_url":
			config.dataURL = value.(string)
		case "trade_url":
			config.tradeURL = value.(string)
		}
	}

	c1 := resty.New().SetDebug(false).SetHostURL(config.dataURL)
	c2 := resty.New().SetDebug(false).SetHostURL(config.tradeURL)
	dataClient = httpClient{c1}
	tradeClient = httpClient{c2}

	dataClient.handleQueryParams()
	tradeClient.handleQueryParams()
}
