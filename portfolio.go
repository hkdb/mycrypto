package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// Process and Display Portfolio Data
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
