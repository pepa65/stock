// stock - main.go

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const version = "0.1.0"

var usage = "stock " + version + ` - Monitor stock by scraping Google Finance
Usage: stock [options]
    -s <Stock symbol>       Stock symbol (case insensitive, default: NVDA)
    -e <Exchange symbol>    Exchange symbol (case insensitive, default: NASDAQ)
    -b <amount>             Bottom price monitored in USD
    -t <amount>             Top price monitored in USD
` // end usage

func man(mess string) {
	fmt.Printf("%s", usage)
	if len(mess) > 0 {
		fmt.Printf("\nERROR: %s\n", mess)
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		man("")
	}

	var stock string
	var exchange string
	var min float64
	var max float64
	flag.StringVar(&stock, "s", "NVDA", "Stock symbol")
	flag.StringVar(&exchange, "e", "NASDAQ", "Exchange symbol")
	flag.Float64Var(&min, "b", -0.1, "Bottom price monitored")
	flag.Float64Var(&max, "t", -0.1, "Top price monitored")
	flag.Parse()
	if stock == "" || exchange == "" || min == -0.1 || max == -0.1 {
		man("all arguments are mandatory")
	}

	if min > max {
		man(fmt.Sprintf("Bottom price (%f) is bigger than Top price (%f)", min, max))
	}

	exchange = strings.ToUpper(exchange)
	stock = strings.ToUpper(stock)
	url := fmt.Sprintf("https://www.google.com/finance/quote/%s:%s", stock, exchange)
	errorAlert := exec.Command("notify-send", "'stock' has exited, restart to continue monitoring")
	var priceregex = regexp.MustCompile(`data-last-price="[^"]*`)

	for {
		res, err := http.Get(url)
		if err != nil {
			errorAlert.Run()
			man("failure to read URL")
		}

		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		if len(data) == 0 {
			errorAlert.Run()
			man(fmt.Sprintf("invalid Exchange symbol (%s) or Stock symbol (%s)", exchange, stock))
		}

		// data-last-price="PRICE"
		price := strings.Split(string(priceregex.Find(data)), `"`)
		if len(price) < 2 {
			errorAlert.Run()
			man(fmt.Sprintf("Stock %s not found in exchange %s", stock, exchange))
		}

		val, err := strconv.ParseFloat(price[1], 64)
		if err != nil {
			errorAlert.Run()
			man(fmt.Sprintf("cannot convert %s to float", price))
		}

		now := time.Now()
		fmt.Printf("%s  %s:%s  USD %f\n", now.Format("2006-01-02 15:04:05"), stock, exchange, val)
		if val < min || val > max {
			var title string
			if val < min {
				title = fmt.Sprintf("Price under Bottom (%.2f)", min)
			} else {
				title = fmt.Sprintf("Price over Top (%.2f)", max)
			}
			message := fmt.Sprintf("%s\n%s:%s is now USD %.2f\n", title, stock, exchange, val)
			alert := exec.Command("notify-send", message)
			alert.Run()
			fmt.Printf(message)
		}

		// 180s + 0..30s
		time.Sleep(time.Duration((180000 + rand.Intn(30000))) * time.Millisecond)
	}
}
