package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

// define text color
var colorReset = "\033[0m"
var colorYellow = "\033[33m"
var colorGreen = "\033[92m"
var colorRed = "\033[91m"
var colorCyan = "\033[36m"

// flags
var s = flag.String("s", "none", "USAGE: -s <option>:\n\t\tt: transactions\n\t\tp: price\n\n      ")
var u = flag.String("p", "none", "USAGE: -p <portfolio>\n\n      ")
var q = flag.String("q", "none", "USAGE: -q <coin symbol>\n\n      ")

var version = "v0.1.0"

func main() {

	fmt.Println("")
	fmt.Println(string(colorGreen), "####################", string(colorReset))
	fmt.Print(string(colorGreen), "  CryptoInfo ", string(colorReset))
	fmt.Println(version)
	fmt.Println(string(colorGreen), "####################", string(colorReset))

	// Handle flags
	flag.Parse()
	if *s != "t" && *s != "p" && *s != "none" {
		fmt.Println("Unknown show option...\n")
		os.Exit(3)
	}

	// get path of binary
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := path.Dir(ex)

	// Get Environment Variables
	e := godotenv.Load(path + "/settings.conf") //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	cpath := os.Getenv("PORTFOLIO_CONF_PATH")
	if cpath == "" {
		cpath = path
	}

	// Load portfolio.conf
	ce := godotenv.Load(cpath + "/portfolio.conf")
	if ce != nil {
		fmt.Print("\n\n")
		fmt.Println(ce)
		fmt.Println(string(colorRed), "\nCan't find portfolio.conf. Double check your settings.conf and ensure that a syntax error free portfolio.conf file is placed in the location as specified...\n\n", string(colorReset))
		os.Exit(0)
	}

	// get path of CSVs
	csv_path := os.Getenv("PORTFOLIO_PATH")
	if csv_path == "" {
		csv_path = path
	}

	// Get coins from settings.conf
	c = os.Getenv("COINS")
	c_array = strings.Split(c, delim)
	// Get CMC API key
	api := os.Getenv("CMC_API")

	fmt.Println("")
	// Get latest data from CoinMarketCap.com
	res := getPrices(api)
	fmt.Println("")

	// Disabling show prices
	if *s == "p" {
		if *q != "none" {
			showPrice(res.Data[strings.ToUpper(*q)])
		} else {
			for i := 0; i < len(c_array); i++ {
				showPrice(res.Data[c_array[i]])
			}
		}
		os.Exit(0)
	}
	user := os.Getenv("PORTFOLIO")

	if *u != "none" {
		user = *u
	}

	fmt.Print(string(colorGreen), "Portfolio: ", string(colorReset))
	fmt.Println(user)
	fmt.Println("")

	if user == "" {
		fmt.Println(string(colorRed), "User must not be empty. Either use the -p flag or enter a default portfolio name in settings.conf...\n\n", string(colorReset))
		os.Exit(0)
	}

	// Ask if the user wants to display his/her portfolio
	fmt.Print("Show Results? (Y/n): ")
	choice := confirm()
	if choice == false {
		os.Exit(0)
	}
	fmt.Println("\n")

	// Read in portfolio data
	if *q != "none" {
		sym := strings.ToUpper(*q)
		t := res.Data[sym]
		cp := t.Quote.USD.Price

		records, err := readData(csv_path + "/" + user + "-" + strings.ToLower(sym) + ".csv")

		if err == nil {
			showPrice(t)
			getCoins(*s, records, user, sym, cp, t.Quote.USD.MarketCap)
		}

		if *s == "t" {
			fmt.Println("\n---------------------------------------------------------------------------------------------------------------------\n\n")
		} else {
			fmt.Println("----------------------------\n\n")
		}
	} else {
		for i := 0; i < len(c_array); i++ {
			// Token
			t := res.Data[c_array[i]]
			// Convert Current Price String to Float
			cp := t.Quote.USD.Price

			records, err := readData(csv_path + "/" + user + "-" + strings.ToLower(c_array[i]) + ".csv")

			if err == nil {
				showPrice(t)
				getCoins(*s, records, user, c_array[i], cp, t.Quote.USD.MarketCap)
				// Read my transactions and show data
				fmt.Println("\n")
			}
		}

		showPortfolioTotal(user)
	}

}
