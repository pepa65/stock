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

const version = "0.5.0"
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
    -c            Console-only: no GUI notifications (notify-send)
    -h            Show this help text (exclusive)
  DESIGNATOR:  STOCK:EXCH (stock) or CUR-CUR (exch.rate) [default: %s]
`, baseintv, randintv, stock)

func helpexit(mess string, console bool) {
	fmt.Printf("%s", usage)
	if len(mess) > 0 {
		fmt.Printf("\nERROR: %s\n", mess)
		if !console {
			errorAlert := exec.Command("notify-send", "-w", "-t", "600000",
				fmt.Sprintf("ERROR alert '%s'", os.Args[0]),
				fmt.Sprintf("'%s' has exited, restart to continue monitoring!", os.Args[0]))
			errorAlert.Start()
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	var min float64
	var max float64
	var intv int
	var mrand int
	var console bool
	var help bool
	designator := stock
	flag.Float64Var(&min, "b", 0, "Bottom price monitored")
	flag.Float64Var(&max, "t", math.MaxFloat64, "Top price monitored")
	flag.IntVar(&intv, "i", baseintv, "Check interval in seconds (plus random add)")
	flag.IntVar(&mrand, "r", randintv, "Max.random add to check interval in seconds")
	flag.BoolVar(&console, "c", false, "Console-only: no GUI notifications")
	flag.BoolVar(&help, "h", false, "Show help text")
	flag.Parse()
	if help {
		helpexit("", true)
	}

	rest := flag.Args()
	if len(rest) > 1 {
		for _, r := range rest {
			if r[0] == '-' {
				helpexit("all flags (start with '-') should come before the designator", console)
			}
		}
		helpexit(fmt.Sprintf("only 1 designator allowed: %s", rest), console)
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
			helpexit("give designator as STOCK:EXCHANGE or as CUR-CUR", console)
		}
	}

	if min > max {
		helpexit(fmt.Sprintf("Bottom price (%f) is bigger than Top price (%f)", min, max), console)
	}

	designator = strings.ToUpper(designator)
	url := fmt.Sprintf("https://www.google.com/finance/quote/%s", designator)

	for {
		res, err := http.Get(url)
		if err != nil {
			helpexit("failure to read URL", console)
		}

		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		if len(data) == 0 {
			helpexit(fmt.Sprintf("invalid designator %s", designator), console)
		}

		// data-last-price="PRICE"
		start := strings.Split(string(data), `data-last-price="`)
		if len(start) < 2 {
			helpexit(fmt.Sprintf("designator %s not found", designator), console)
		}

		price := strings.Split(start[1], `"`)
		if len(price) < 2 {
			helpexit(fmt.Sprintf("designator %s not found", designator), console)
		}

		val, err := strconv.ParseFloat(price[0], 64)
		if err != nil {
			helpexit(fmt.Sprintf("cannot convert %s to float", price), console)
		}

		now := time.Now()
		fmt.Printf("%s  %s  %s %f\n", now.Format("2006-01-02_15:04:05"), designator, curr, val)
		if val < min || val > max {
			if !console {
				title := fmt.Sprintf("Price alert '%s'", os.Args[0])
				var msg string
				if val < min {
					msg = fmt.Sprintf("Price under Bottom (%.2f)", min)
				} else {
					msg = fmt.Sprintf("Price over Top (%.2f)", max)
				}
				message := fmt.Sprintf("%s\n%s is now %s %.2f\n", msg, designator, curr, val)
				alert := exec.Command("notify-send", "-w", "-t", "600000", title, message)
				alert.Start()
			}
		}
		time.Sleep(time.Duration((intv * 1000 + rand.Intn(mrand * 1000))) * time.Millisecond)
	}
}
