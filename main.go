// stock - main.go

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const version = "0.4.0"
const stock = "NVDA:NASDAQ"
const baseintv = 300 // seconds
const randintv = 60 // seconds

var usage = "stock v" + version +
	fmt.Sprintf(` - Monitor stock or exchange rate by scraping Google Finance
Usage:  stock [OPTIONS] DESIGNATOR
  OPTIONS:
    -b <Price>    Bottom price monitored in USD (optional)
    -t <Price>    Top price monitored in USD (optional)
    -i <Seconds>  Check interval (plus random add) [default: %d]
    -r <Seconds>  Max.random add to check interval [default: %d]
    -h            Show this help text (exclusive)
  DESIGNATOR:  STOCK:EXCH (stock) or CUR-CUR (exch.rate) [default: %s]
`, baseintv, randintv, stock)

func helpexit(mess string) {
	fmt.Printf("%s", usage)
	if len(mess) > 0 {
		fmt.Printf("\nERROR: %s\n", mess)
		errorAlert := exec.Command("notify-send", "-w", "-t", "600000",
			fmt.Sprintf("'%s' has exited, restart to continue monitoring!", os.Args[0]))
		errorAlert.Start()
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	var min float64
	var max float64
	var intv int
	var mrand int
	var help bool
	designator := stock
	flag.Float64Var(&min, "b", 0, "Bottom price monitored")
	flag.Float64Var(&max, "t", math.MaxFloat64, "Top price monitored")
	flag.IntVar(&intv, "i", baseintv, "Check interval in seconds (plus random add)")
	flag.IntVar(&mrand, "r", randintv, "Max.random add to check interval in seconds")
	flag.BoolVar(&help, "h", false, "Show help text")
	flag.Parse()
	if help {
		helpexit("")
	}

	rest := flag.Args()
	if len(rest) > 1 {
		for _, r := range rest {
			if r[0] == '-' {
				helpexit("all flags (start with '-') should come before the designator")
			}
		}
		helpexit(fmt.Sprintf("only 1 designator allowed: %s", rest))
	}

	if len(rest) == 1 {
		designator = rest[0]
	}

	curr := "USD"
	currs := strings.Split(designator, "-")
	if len(currs) == 2 {
		curr = currs[1]
	} else if len(currs) == 1 {
		currs = strings.Split(designator, ":")
		if len(currs) != 2 {
			helpexit("give designator as STOCK:EXCHANGE or as CUR-CUR")
		}
	}

	if min > max {
		helpexit(fmt.Sprintf("Bottom price (%f) is bigger than Top price (%f)", min, max))
	}

	designator = strings.ToUpper(designator)
	url := fmt.Sprintf("https://www.google.com/finance/quote/%s", designator)

	for {
		res, err := http.Get(url)
		if err != nil {
			helpexit("failure to read URL")
		}

		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		if len(data) == 0 {
			helpexit(fmt.Sprintf("invalid designator %s", designator))
		}

		// data-last-price="PRICE"
		start := strings.Split(string(data), `data-last-price="`)
		if len(start) < 2 {
			helpexit(fmt.Sprintf("designator %s not found", designator))
		}

		price := strings.Split(start[1], `"`)
		if len(price) < 2 {
			helpexit(fmt.Sprintf("designator %s not found", designator))
		}

		val, err := strconv.ParseFloat(price[0], 64)
		if err != nil {
			helpexit(fmt.Sprintf("cannot convert %s to float", price))
		}

		now := time.Now()
		fmt.Printf("%s  %s  %s %f\n", now.Format("2006-01-02_15:04:05"), designator, curr, val)
		if val < min || val > max {
			var title string
			if val < min {
				title = fmt.Sprintf("Price under Bottom (%.2f)", min)
			} else {
				title = fmt.Sprintf("Price over Top (%.2f)", max)
			}
			message := fmt.Sprintf("%s\n%s is now %s %.2f\n", title, designator, curr, val)
			alert := exec.Command("notify-send", "-w", "-t", "600000", message)
			alert.Start()
			fmt.Printf(message)
		}
		time.Sleep(time.Duration((intv * 1000 + rand.Intn(mrand * 1000))) * time.Millisecond)
	}
}
