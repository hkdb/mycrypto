package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
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

	// Convert Current Price String to Float
	btc_current := res.Data.BTC.Quote.USD.Price
	eth_current := res.Data.ETH.Quote.USD.Price

	// Read my transactions
	getCoins(user, "BTC", btc_current)
	fmt.Println("\n")
	getCoins(user, "ETH", eth_current)
	fmt.Println("\n")

}

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

func getCoins(name string, symbol string, current float64) {

	// get path of binary
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := path.Dir(ex)

	records, err := readData(path + "/" + name + "-" + strings.ToLower(symbol) + ".csv")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name + "'s " + symbol + " Transaction History:")
	fmt.Println("------------------------------------------------------------------------------------------------")

	totalCoins := 0.00
	totalCost := 0.00

	for _, record := range records {
		trans := MyCoins{
			Date:   record[0],
			Coin:   record[1],
			Cost:   record[2],
			Fee:    record[3],
			Total:  record[4],
			Amount: record[5],
			Wallet: record[6],
		}

		fmt.Print("Date: ")
		fmt.Print(string(colorYellow), trans.Date, string(colorReset))
		fmt.Print(" | ")
		fmt.Print("Coin: ")
		fmt.Print(string(colorYellow), trans.Coin, string(colorReset))
		fmt.Print(" | ")
		fmt.Print("Fee: ")
		fmt.Print(string(colorYellow), trans.Fee, string(colorReset))
		fmt.Print(" | ")
		fmt.Print("Total: ")
		fmt.Print(string(colorYellow), trans.Total, string(colorReset))
		fmt.Print(" | ")
		fmt.Print("Amount: ")
		fmt.Print(string(colorYellow), trans.Amount, string(colorReset))
		fmt.Print(" | ")
		fmt.Print("Wallet: ")
		fmt.Print(string(colorYellow), trans.Wallet+"\n", string(colorReset))

		newCoin, err := strconv.ParseFloat(trans.Amount, 64)
		if err != nil {
			fmt.Println(err)
		}
		totalCoins = totalCoins + newCoin

		newCost, err := strconv.ParseFloat(trans.Total, 64)
		if err != nil {
			fmt.Println(err)
		}
		totalCost = totalCost + newCost

	}

	fmt.Println("------------------------------------------------------------------------------------------------\n")

	value := current * totalCoins
	change := value - totalCost
	change_percent := change / totalCost * 100

	c_color := string(colorGreen)
	if change < 0 {
		c_color = string(colorRed)
	}
	cp_color := string(colorGreen)
	if change_percent < 0 {
		cp_color = string(colorRed)
	}

	fmt.Print("Total " + symbol + ": ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%f", totalCoins)+"\n", string(colorReset))
	fmt.Print("Total Cost ($): ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%.2f", totalCost)+"\n", string(colorReset))
	fmt.Print("Current Value ($): ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%.2f", value)+"\n", string(colorReset))
	fmt.Print("Change ($): ")
	fmt.Print(c_color, fmt.Sprintf("%.2f", change)+"\n", string(colorReset))
	fmt.Print("Change (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", change_percent)+"%\n", string(colorReset))

}

// Read CSV
func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
