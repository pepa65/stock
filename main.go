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

const version = "0.2.2"

var usage = "stock v" + version + ` - Monitor exchange rate by scraping Google Finance
Usage:  stock [OPTIONS] DESIGNATOR
  OPTIONS:
    -b <Bottom>    Bottom price monitored in USD (optional)
    -t <Top>       Top price monitored in USD (optional)
    -h             Show this help text (exclusive)
  DESIGNATOR:      STOCK:EXCHANGE (stock) or CUR-CUR (exchange rate)
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
	var min float64
	var max float64
	var help bool
	designator := "NVDA:NASDAQ"
	flag.Float64Var(&min, "b", 0, "Bottom price monitored")
	flag.Float64Var(&max, "t", math.MaxFloat64, "Top price monitored")
	flag.BoolVar(&help, "h", false, "Show help text")
	flag.Parse()
	if help {
		man("")
	}

	rest := flag.Args()
	if len(rest) > 1 {
		for _, r := range rest {
			if r[0] == '-' {
				man("all flags (start with '-') should come before the designator")
			}
		}
		man(fmt.Sprintf("only 1 designator allowed: %s", rest))
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
			man("give designator as STOCK:EXCHANGE or as CUR-CUR")
		}
	}

	if min > max {
		man(fmt.Sprintf("Bottom price (%f) is bigger than Top price (%f)", min, max))
	}

	designator = strings.ToUpper(designator)
	url := fmt.Sprintf("https://www.google.com/finance/quote/%s", designator)
	errorAlert := exec.Command("notify-send", "'stock' has exited, restart to continue monitoring")

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
			man(fmt.Sprintf("invalid designator %s", designator))
		}

		// data-last-price="PRICE"
		start := strings.Split(string(data), `data-last-price="`)
		if len(start) < 2 {
			errorAlert.Run()
			man(fmt.Sprintf("designator %s not found", designator))
		}

		price := strings.Split(start[1], `"`)
		if len(price) < 2 {
			errorAlert.Run()
			man(fmt.Sprintf("designator %s not found", designator))
		}

		val, err := strconv.ParseFloat(price[0], 64)
		if err != nil {
			errorAlert.Run()
			man(fmt.Sprintf("cannot convert %s to float", price))
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
			alert := exec.Command("notify-send", message)
			alert.Run()
			fmt.Printf(message)
		}

		// 180s + 0..30s
		time.Sleep(time.Duration((180000 + rand.Intn(30000))) * time.Millisecond)
	}
}
