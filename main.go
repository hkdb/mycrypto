package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func main() {

	// get path of binary
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := path.Dir(ex)

	// Get Environment Variables
	e := godotenv.Load(path + "/.config") //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	api := os.Getenv("CMC_API")
	user := os.Getenv("NAME")

	fmt.Println("")
	// Get latest data from CoinMarketCap.com
	res := getPrices(api)
	fmt.Println("")

	// BTC
	showPrice(res.Data.BTC)

	// ETH
	showPrice(res.Data.ETH)

	// Ask if the user wants to display his/her portfolio
	fmt.Print("Show Portfolio? (Y/n): ")
	choice := confirm()
	if choice == false {
		os.Exit(0)
	}
	fmt.Println("\n")

	// Convert Current Price String to Float
	btc_current := res.Data.BTC.Quote.USD.Price
	eth_current := res.Data.ETH.Quote.USD.Price

	// Read my transactions
	getCoins(user, "BTC", btc_current)
	fmt.Println("\n")
	getCoins(user, "ETH", eth_current)
	fmt.Println("\n")

}
