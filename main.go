package main

import (
	"fmt"
	"log"
	"time"

	"github.com/afa7789/satsukashii/pkg/bitcoin_price"
)

func main() {
	// print("Bitcoin Price Fetcher\n")
	// Load CSV file
	fetcher, err := bitcoin_price.NewBTCPricesCSV("assets/bitcoin_2016-01-01_2025-05-01.csv")
	if err != nil {
		log.Fatalf("Failed to load CSV: %v", err)
	}

	// Fetch price by date
	// date, _ := time.Parse("2006-01-02", "2024-01-02")
	// price, err := fetcher.FetchPriceByDate(date)
	// if err != nil {
	// 	log.Fatalf("Price not found for %s: %v", date.Format("2006-01-02"), err)
	// // }

	// fmt.Printf("Price on %s: Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f\n",
	// 	date.Format("2006-01-02"),
	// 	price.Open, price.High, price.Low, price.Close)

	// Fetch historical data after a date
	startYear := "2016"
	endYear := "2025"
	startDate, _ := time.Parse("2006-01-02", startYear+"-01-01")
	prices, err := fetcher.FetchHistoricalData(startDate)
	if err != nil {
		log.Fatalf("Failed to fetch historical data: %v", err)
	}

	print("\n")
	// Load CSV file
	fetcherXMR, err := bitcoin_price.NewBTCPricesCSV("assets/monero_2016-01-01_2025-05-01.csv")
	if err != nil {
		log.Fatalf("Failed to load CSV: %v", err)
	}

	// // Fetch price by date
	// // date, _ = time.Parse("2006-01-02", "2024-01-02")
	// priceXMR, err := fetcherXMR.FetchPriceByDate(date)
	// if err != nil {
	// 	log.Fatalf("Price not found for %s: %v", date.Format("2006-01-02"), err)
	// }

	// fmt.Printf("Price on %s: Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f\n",
	// 	date.Format("2006-01-02"),
	// 	priceXMR.Open, priceXMR.High, priceXMR.Low, priceXMR.Close)

	// Fetch historical data after a date
	// startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	pricesXMR, err := fetcherXMR.FetchHistoricalData(startDate)
	if err != nil {
		log.Fatalf("Failed to fetch historical data: %v", err)
	}

	// Calculate weekly investments on every Friday from 2016-01-01 to 2025-05-01
	// startDate, _ := time.Parse("2006-01-02", "2016-01-01")
	endDate, _ := time.Parse("2006-01-02", endYear+"-05-01")

	var totalInvested float64
	var bitcoinAccum float64
	var moneroAccum float64

	// Loop day-by-day over the period.
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Friday {
			// Use the price for the day if it exists in both datasets.
			btcPrice, hasBTC := prices[d]
			xmrPrice, hasXMR := pricesXMR[d]
			if hasBTC && hasXMR {
				// With $10, the coin amount purchased is 10 / open price.
				bitcoinAccum += 10 / btcPrice.Open
				moneroAccum += 10 / xmrPrice.Open
				totalInvested += 10
			}
		}
	}

	// Find the last available price on or before endDate for both coins.
	var lastBTCPrice = prices[endDate]
	{
		_, ok := prices[endDate]
		if !ok {
			for d := endDate; d.After(startDate); d = d.AddDate(0, 0, -1) {
				if p, exists := prices[d]; exists {
					lastBTCPrice = p
					break
				}
			}
		}
	}

	var lastXMRPrice = pricesXMR[endDate]
	{
		_, ok := pricesXMR[endDate]
		if !ok {
			for d := endDate; d.After(startDate); d = d.AddDate(0, 0, -1) {
				if p, exists := pricesXMR[d]; exists {
					lastXMRPrice = p
					break
				}
			}
		}
	}

	// Calculate current value by multiplying the accumulated coins with the last available open price.
	finalBitcoinValue := bitcoinAccum * lastBTCPrice.Open
	finalMoneroValue := moneroAccum * lastXMRPrice.Open

	// ANSI escape code for yellow color and reset code.
	yellow := "\033[33m"
	reset := "\033[0m"

	fmt.Printf("Investment period: %s%s%s to %s%s%s\n", yellow, startDate.Format("2006-01-02"), reset, yellow, endDate.Format("2006-01-02"), reset)
	fmt.Printf("Total invested weekly: %s$%.2f%s\n", yellow, totalInvested, reset)
	weeks := int(endDate.Sub(startDate).Hours() / 24 / 7)
	fmt.Printf("Number of weeks: %s%d%s\n", yellow, weeks, reset)
	fmt.Printf("Value per week: %s$%.2f%s\n", yellow, totalInvested/float64(weeks), reset)
	fmt.Printf("Bitcoin:\n\tTotal coins accumulated: %s%.6f%s,\n\tValue on %s: %s$%.2f%s\n",
		yellow, bitcoinAccum, reset,
		endDate.Format("2006-01-02"),
		yellow, finalBitcoinValue, reset)
	fmt.Printf("Monero:\n\tTotal coins accumulated: %s%.6f%s,\n\tValue on %s: %s$%.2f%s\n",
		yellow, moneroAccum, reset,
		endDate.Format("2006-01-02"),
		yellow, finalMoneroValue, reset)
	fmt.Printf("Amount got per week for:\n\tBitcoin: $%s%.2f%s,\n\tMonero: $%s%.2f%s\n",
		yellow, finalBitcoinValue/float64(weeks), reset,
		yellow, finalMoneroValue/float64(weeks), reset)
	fmt.Printf("Total value of:\n\tBitcoin: %s$%.2f%s,\n\tMonero: %s$%.2f%s\n",
		yellow, finalBitcoinValue, reset,
		yellow, finalMoneroValue, reset)
	fmt.Print("\n")
}
