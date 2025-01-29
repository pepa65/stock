[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/stock)](https://goreportcard.com/report/github.com/pepa65/stock)
[![GoDoc](https://godoc.org/github.com/pepa65/stock?status.svg)](https://godoc.org/github.com/pepa65/stock)
[![GitHub](https://img.shields.io/github/license/pepa65/stock.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/stock/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/stock/actions/workflows/ci.yml) 

# stock 0.1.1
**Monitor stock by scraping Google Finance**

Follows a single stock on a specified exchange, monitors a bottom & top price and
alerts (on Linux with notify-send) when the stock price goes outside of the range.
The Google Finance page for the stock + exchange is scraped
(where some stocks are real-time, but some stocks can have a 15 minute delay!).

* Repo: <https://github.com/pepa65/stock>
* Requires: libnotify-bin(notify-send)

## Install
### With golang installed
`go install github.com/pepa65/stock@latest`

## Usage
```
stock 0.1.1 - Monitor stock by scraping Google Finance
Usage: stock [options]
    -s <Stock symbol>       Stock symbol (case insensitive, default: NVDA)
    -e <Exchange symbol>    Exchange symbol (case insensitive, default: NASDAQ)
    -b <amount>             Bottom price monitored in USD
    -t <amount>             Top price monitored in USD
    -h                      Show this help text

```

### Examples
```
# This monitors Nvidia on NASDAQ exchange
stock

# This monitors Nvidia on NASDAQ exchange within the USD 100..200 range
stock -b 100 -t 200

# This monitors Siam Cement on BKK exchange
stock -s SCC -e BKK
```
