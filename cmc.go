package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type status struct {
	TimeStamp   string `json:"timestamp"`
	ErrorCode   int    `json:"error_code"`
	ErrorMsg    string `json:"error_message"`
	Elapsed     int    `json:"elapsed"`
	CreditCount int    `json:"credit_count"`
	Notice      string `json:"notice"`
}

type price struct {
	Price    float64 `json:"price"`
	Change1H float64 `json:"percent_change_1h"`
	Change1D float64 `json:"percent_change_24h"`
	Change7D float64 `json:"percent_change_7d"`
	Last     string  `json:"last_updated"`
}

type quote struct {
	USD price `json:"USD"`
}

type coin struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Slug        string `json:"slug"`
	Rank        int    `json:"rank"`
	LastUpdated string `json:"last_updated"`
	Quote       quote  `json:"quote"`
}

var c = "BTC,ETH"

type Data struct {
	BTC coin `json:"BTC"`
	ETH coin `json:"ETH"`
}

type Response struct {
	Status status `json:"status"`
	Data   Data   `json:"data"`
}

type MyCoins struct {
	Date   string
	Coin   string
	Cost   string
	Fee    string
	Total  string
	Amount string
	Wallet string
}

// define text color
var colorReset = "\033[0m"
var colorYellow = "\033[33m"
var colorGreen = "\033[92m"
var colorRed = "\033[91m"

// Get CoinMarketCap Data
func getPrices(api string) Response {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("symbol", c)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", api)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println("Retrieving Data from CoinMarketCap.com")
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))

	// Ouptut Response to STDOUT
	var response Response
	json.Unmarshal([]byte(string(respBody)), &response)
	status := "Error"
	if response.Status.ErrorCode == 0 {
		status = "Success"
	}
	fmt.Println("Status: ", status)
	return response

}

// Show CoinMarketCap Data
func showPrice(c coin) {

	fmt.Print("Coin: ")
	fmt.Print(string(colorYellow), c.Symbol, string(colorReset))
	fmt.Print(" | ")
	fmt.Print("Price: ")
	fmt.Print(string(colorGreen), c.Quote.USD.Price, string(colorReset))
	fmt.Print(" | ")
	fmt.Print("Change (1H): ")
	fmt.Print(string(colorYellow), c.Quote.USD.Change1H, string(colorReset))
	fmt.Print(" | ")
	fmt.Print("Change (24H): ")
	fmt.Print(string(colorYellow), c.Quote.USD.Change1D, string(colorReset))
	fmt.Print(" | ")
	fmt.Print("Change (7D): ")
	fmt.Print(string(colorYellow), c.Quote.USD.Change7D, string(colorReset))
	fmt.Print(" | ")
	fmt.Print("Last Updated: ")
	fmt.Print(string(colorYellow), c.Quote.USD.Last+"\n\n", string(colorReset))

}
