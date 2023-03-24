package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"

	"github.com/martinsirbe/go-graham-number/internal/alphavantage"
)

type config struct {
	AlphaVantageAPIKey string `env:"ALPHA_VANTAGE_API_KEY"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if len(os.Args) != 2 {
		log.WithField("args", os.Args).Fatal("incorrect usage: \"gh SYMBOL\"")
	}

	stockSymbol := strings.ToUpper(os.Args[1])

	client := alphavantage.NewClient(cfg.AlphaVantageAPIKey)
	stockDetails, err := client.GetStockDetails(stockSymbol)
	if err != nil {
		log.Fatalf("failed to get stock details: %v", err)
	}

	value := 22.5 * stockDetails.EPS * stockDetails.BVPS
	grahamNumber := float64(0)
	if value > 0 {
		grahamNumber = math.Sqrt(value)
	}

	status := "overvalued"
	if grahamNumber > stockDetails.Price {
		status = "undervalued"
	}

	fmt.Printf("%s %s %.2f [P: %.2f, EPS: %.2f, BVPS: %.2f]\n", stockSymbol, status, grahamNumber,
		stockDetails.Price, stockDetails.EPS, stockDetails.BVPS)
}
