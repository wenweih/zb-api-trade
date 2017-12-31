package main

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

func init() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

// Exit 程序推出函数
func Exit(str ...string) {
	log.Error(strings.Join(str, ""))
	os.Exit(1)
}

// StrLogger 字符串日志
func StrLogger(str ...string) {
	log.Info(strings.Join(str, ""))
}

// HomeDir 获取服务器当亲用户目录路径
func HomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}
