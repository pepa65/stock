[![Go Report Card](https://goreportcard.com/badge/github.com/pepa65/stock)](https://goreportcard.com/report/github.com/pepa65/stock)
[![GoDoc](https://godoc.org/github.com/pepa65/stock?status.svg)](https://godoc.org/github.com/pepa65/stock)
[![GitHub](https://img.shields.io/github/license/pepa65/stock.svg)](LICENSE)
[![run-ci](https://github.com/pepa65/stock/actions/workflows/ci.yml/badge.svg)](https://github.com/pepa65/stock/actions/workflows/ci.yml) 

# stock v0.7.0
**Monitor stock or exchange rate by scraping Google Finance**

* The Google Finance page for the stock or the exchange rate is scraped.
* Follows a single stock on a specified exchange or a (crypto) currency exchange rate.
* Optionally monitors a bottom and/or a top price and then alerts when the price goes outside of the range.
* The alerts can be terminal only, or use a GUI on Linux, BSD and MacOS.
* Repo: <https://github.com/pepa65/stock>
* Required:
  - libnotify-bin(notify-send) [for GUI alerts on Linux/BSD]
  - terminal-notifier [for GUI alerts on MacOS]

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
(Or `GOOS` can be `darwin` for MacOS, or
`dragonfly`, `freebsd`, `netbsd` or `openbsd` for BSD.)

## Usage
```
stock v0.7.0 - Monitor stock or exchange rate by scraping Google Finance
Usage:  stock [OPTIONS] DESIGNATOR
  OPTIONS:
    -b <Price>    Bottom price monitored in USD (optional)
    -t <Price>    Top price monitored in USD (optional)
    -i <Seconds>  Check interval (plus random add) [default: 300]
    -r <Seconds>  Max.random add to check interval [default: 60]
    -c            Console-only: no GUI notifications (notify-send)
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

# Monitor Nvidia on NASDAQ within USD 120..140 without GUI
stock -b 120 -t 140 -c
```
