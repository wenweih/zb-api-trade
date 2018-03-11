# TokenTrade
Golang 实现的数字货币 API 交易，由于中币收手续费了，现开源出来，策略需要你自己写。
### Install
假设你已经有了 Golang 开发环境。

```sh
dep ensure -update -v
go install

cp zb.yml to your home directory
```
### Usage
```sh
➜  TokenTrade git:(master) TokenTrade -h
A tool for earning money

Usage:
  TokenTrade [flags]
  TokenTrade [command]

Available Commands:
  help        Help about any command
  start       Start to trade...

Flags:
  -c, --config string   config file (default is $HOME/zb.yaml)
  -h, --help            help for TokenTrade

Use "TokenTrade [command] --help" for more information about a command.
```
