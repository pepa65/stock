[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/stock)](https://goreportcard.com/report/github.com/pepa65/stock)
[![GoDoc](https://godoc.org/github.com/pepa65/stock?status.svg)](https://godoc.org/github.com/pepa65/stock)
[![GitHub](https://img.shields.io/github/license/pepa65/stock.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/stock/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/stock/actions/workflows/ci.yml) 

# stock v0.4.0
**Monitor stock or exchange rate by scraping Google Finance**

Follows a single stock on a specified exchange or a currency exchange rate.
Optionally monitors a bottom & top price and then alerts
(on Linux with notify-send) when the price goes outside of the range.
The Google Finance page for the stock or the exchange rate is scraped.

* Repo: <https://github.com/pepa65/stock>
* Requires: libnotify-bin(notify-send)[for alerts]

## Install
### Download
```
sudo wget -O /usr/local/bin/stock 4e4.in/stock
sudo chmod +x /usr/local/bin/stock
```

### With golang installed
`go install github.com/pepa65/stock@latest`

### With golang from github repo
```
git clone https://github.com/pepa65/stock
cd stock
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o stock
sudo cp stock /usr/local/bin/
sudo chown root:root /usr/local/bin/stock
```

## Usage
```
stock v0.4.0 - Monitor stock or exchange rate by scraping Google Finance
Usage:  stock [OPTIONS] DESIGNATOR
  OPTIONS:
    -b <Price>    Bottom price monitored in USD (optional)
    -t <Price>    Top price monitored in USD (optional)
    -i <Seconds>  Check interval (plus random add) [default: 300]
    -r <Seconds>  Max.random add to check interval [default: 60]
    -h            Show this help text (exclusive)
  DESIGNATOR:  STOCK:EXCH (stock) or CUR-CUR (exch.rate) [default: NVDA:NASDAQ]
```

### Examples
```
# Monitor Microsoft on NASDAQ exchange
stock MSFT:NASDAQ

# Monitor Nvidia on NASDAQ within the USD 100..200 range
stock -b 100 -t 200 NVDA:NASDAQ

# Monitor Siam Cement on BKK exchange with a USD 100 Bottom
stock -b 100 SCC:BKK

# Monitor BTC-USD exchange rate (price of BTC in USD) with a USD 110000 Top
stock -t 110000 BTC-USD

# Monitor Nvidia on NASDAQ with 10 minute (plus max. 1 minute) interval
stock -i 600

# Monitor ETH-EUR rate (price of ETH in EUR) with max. 10 seconds random add
stock -r 10 ETH-EUR
```
