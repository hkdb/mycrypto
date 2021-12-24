package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Process and Display Portfolio Data
func getCoins(show string, records [][]string, name string, symbol string, val float64, mcap float64) {
	if show == "t" {
		fmt.Println(name + "'s " + symbol + " Transaction History:")
		fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	}

	totalCoins := 0.00
	t := 0 // Number of transaction entries per coin
	totalPrice := 0.00
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

		if show == "t" {
			fmt.Print("Date: ")
			fmt.Print(string(colorYellow), trans.Date, string(colorReset))
			fmt.Print(" | ")
			fmt.Print("Cost: ")
			fmt.Print(string(colorYellow), trans.Cost, string(colorReset))
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
			fmt.Print("Price: ")
		}

		amount, aerr := strconv.ParseFloat(trans.Amount, 64)
		if aerr != nil {
			fmt.Println(aerr)
		}
		cost, cerr := strconv.ParseFloat(trans.Cost, 64)
		if cerr != nil {
			fmt.Println(cerr)
		}
		price := 1.00 / amount * cost

		if *s == "t" {
			priceString := strconv.FormatFloat(price, 'f', 2, 64)

			fmt.Print(string(colorYellow), priceString, string(colorReset))
			fmt.Print(" | ")
			fmt.Print("Wallet: ")
			fmt.Print(string(colorYellow), trans.Wallet+"\n", string(colorReset))
		}

		if cost != 0 {
			totalPrice = totalPrice + price
			t++
		}
		totalCoins = totalCoins + amount

		newCost, err := strconv.ParseFloat(trans.Total, 64)
		if err != nil {
			fmt.Println(err)
		}
		totalCost = totalCost + newCost
	}

	if *s == "t" {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------\n")
	} else {
		fmt.Println("----------------------------")
	}

	value := val * totalCoins
	ch := value - totalCost
	cp := 0.000000
	if totalCoins != 0.000000 {
		cp = ch / math.Abs(totalCost) * 100
	}

	c_color := string(colorGreen)
	if ch < 0 {
		c_color = string(colorRed)
	}
	cp_color := string(colorGreen)
	if cp < 0 {
		cp_color = string(colorRed)
	}

	fmt.Print("Total " + symbol + ": ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%f", totalCoins)+"\n", string(colorReset))

	avgp := totalPrice / float64(t)
	fmt.Print("Avg Price ($): ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%f", avgp)+"\n", string(colorReset))

	fmt.Print("Total Cost ($): ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%.2f", totalCost)+"\n", string(colorReset))
	fmt.Print("Current Value ($): ")
	fmt.Print(string(colorGreen), fmt.Sprintf("%.2f", value)+"\n", string(colorReset))
	fmt.Print("Change ($): ")
	fmt.Print(c_color, fmt.Sprintf("%.2f", ch)+"\n", string(colorReset))
	fmt.Print("Change (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", cp)+"%\n", string(colorReset))

	cost = cost + totalCost
	current = current + value
	change = change + ch

	if symbol == "BTC" {
		btc = btc + totalCost
		btcr = btcr + value
	} else if mcap >= 10000000000 {
		bcap = bcap + totalCost
		bcapr = bcapr + value
	} else {
		scap = scap + totalCost
		scapr = scapr + value
	}
}

func showPortfolioTotal(name string) {

	change_percent := (change / math.Abs(cost)) * 100

	btc_percent := btc / cost * 100
	bcap_percent := bcap / cost * 100
	scap_percent := scap / cost * 100

	btc_realized := btcr / current * 100
	bcap_realized := bcapr / current * 100
	scap_realized := scapr / current * 100

	// Coloring
	c_color := string(colorGreen)
	if change < 0 {
		c_color = string(colorRed)
	}
	cp_color := string(colorGreen)
	if change_percent < 0 {
		cp_color = string(colorRed)
	}

	fmt.Println(name + "'s Portfolio Position:")
	fmt.Println("#####################################################################################################################")
	fmt.Print("Cost ($): ")
	fmt.Print(c_color, fmt.Sprintf("%.2f", cost)+"\n", string(colorReset))
	fmt.Print("Current ($): ")
	fmt.Print(c_color, fmt.Sprintf("%.2f", current)+"\n", string(colorReset))
	fmt.Print("Change ($): ")
	fmt.Print(c_color, fmt.Sprintf("%.2f", change)+"\n", string(colorReset))
	fmt.Print("Change (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", change_percent)+"%\n", string(colorReset))
	fmt.Print("BTC Holdings (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", btc_percent)+"%\n", string(colorReset))
	fmt.Print("Big Cap Holdings (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", bcap_percent)+"%\n", string(colorReset))
	fmt.Print("Small Cap Holdings (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", scap_percent)+"%\n", string(colorReset))
	fmt.Print("BTC Holdings Realized (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", btc_realized)+"%\n", string(colorReset))
	fmt.Print("Big Cap Holdings Realized (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", bcap_realized)+"%\n", string(colorReset))
	fmt.Print("Small Cap Holdings Realized (%): ")
	fmt.Print(cp_color, fmt.Sprintf("%.2f", scap_realized)+"%\n", string(colorReset))
	fmt.Println("#####################################################################################################################")
	fmt.Println("\n")

}

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
