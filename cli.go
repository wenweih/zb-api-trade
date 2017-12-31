// Package cmd 命令行
package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config 中币接口配置信息
var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "zb_trade",
		Short: "A tool for earning money",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// Execute 命令行入口
func cmdExecute() {
	if err := rootCmd.Execute(); err != nil {
		Exit(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.zb.yaml)")
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
		viper.SetConfigName(".zb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		StrLogger("Using Configure file:", viper.ConfigFileUsed())
	} else {
		Exit("Error: zb.yml not found in: ", HomeDir())
	}
}
