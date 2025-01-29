[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/stock)](https://goreportcard.com/report/github.com/pepa65/stock)
[![GoDoc](https://godoc.org/github.com/pepa65/stock?status.svg)](https://godoc.org/github.com/pepa65/stock)
[![GitHub](https://img.shields.io/github/license/pepa65/stock.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/stock/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/stock/actions/workflows/ci.yml) 

# stock 0.1.0
**Monitor stock by scraping Google Finance**

Follows a single stock on a specified exchange, monitors a bottom & top price and
alerts (on Linux with notify-send) when the stock price goes outside of the range.
The Google Finance page for the stock + exchange is scraped
(where some stocks are real-time, but some stocks can have a 15 minute delay!).

* Repo: <https://github.com/pepa65/stock>
* Requires: libnotify-bin(notify-send)

## Install
### With golang installed
`go install github.com/pepa65/stock`

## Usage
```
```

### Example
`stock -b 100 -t 200`

This monitors NVDA on NASDAQ within the USD 100..200 range
