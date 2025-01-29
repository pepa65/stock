[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/stock)](https://goreportcard.com/report/github.com/pepa65/stock)
[![GoDoc](https://godoc.org/github.com/pepa65/stock?status.svg)](https://godoc.org/github.com/pepa65/stock)
[![GitHub](https://img.shields.io/github/license/pepa65/stock.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/stock/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/stock/actions/workflows/ci.yml) 

# stock v0.2.0
**Monitor exchange rate by scraping Google Finance**

Follows a single stock on a specified exchange or a currency exchange rate.
Optionally monitors a bottom & top price and then alerts
(on Linux with notify-send) when the price goes outside of the range.
The Google Finance page for the stock or the exchange rate is scraped.

* Repo: <https://github.com/pepa65/stock>
* Requires: libnotify-bin(notify-send)[for alerts]

## Install
### With golang installed
`go install github.com/pepa65/stock@latest`

## Usage
```
stock v0.2.0 - Monitor exchange rate by scraping Google Finance
Usage:  stock [OPTIONS] DESIGNATOR
  OPTIONS:
    -b <Bottom>    Bottom price monitored in USD (optional)
    -t <Top>       Top price monitored in USD (optional)
    -h             Show this help text (exclusive)
  DESIGNATOR:      STOCK:EXCHANGE (stock) or CUR-CUR (exchange rate)
```

### Examples
```
# Monitor Nvidia on NASDAQ exchange
stock NVDA:NASDAQ

# Monitor Nvidia on NASDAQ within the USD 100..200 range
stock -b 100 -t 200 NVDA:NASDAQ

# Monitor Siam Cement on BKK exchange with a USD 100 Bottom
stock - 100 SCC:BKK

# Monitor BTC-USD exchange rate (price of BTC in USD) with a USD 110000 Top
stock -t 110000 BTC-USD
```
