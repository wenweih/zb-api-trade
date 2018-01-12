package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"os"
	"sort"
	"strings"

	"github.com/go-resty/resty"
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
		Exit(err.Error())
	}
	return home
}

// SHA1 加密
func digest() string {
	hash := sha1.New()
	hash.Write([]byte(config.secretkey))
	return hex.EncodeToString(hash.Sum(nil))
}

// hmac MD5
func hmacSign(message string) string {
	hmac := hmac.New(md5.New, []byte(digest()))
	hmac.Write([]byte(message))
	return hex.EncodeToString(hmac.Sum(nil))
}

// 参数按照字母排序
func sortParams(params map[string]string) string {
	var buffer bytes.Buffer
	sortKey := make([]string, 0, len(params))

	for k := range params {
		sortKey = append(sortKey, k)
	}
	sort.Strings(sortKey)

	for _, k := range sortKey {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(params[k])
		buffer.WriteString("&")
	}
	return strings.TrimSuffix(buffer.String(), "&")
}

func resetQueryParams(client *resty.Client) {
	client.OnAfterResponse(func(client *resty.Client, req *resty.Response) error {
		for k := range client.QueryParam {
			delete(client.QueryParam, k)
		}
		return nil
	})
}
