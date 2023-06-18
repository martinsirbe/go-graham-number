package stocks

import (
	"fmt"
	"math"

	"github.com/martinsirbe/go-graham-number/internal/alphavantage"
)

type Client struct {
	alphaVantageClient *alphavantage.Client
}

func NewClient(alphaVantageAPIKey string) *Client {
	return &Client{
		alphaVantageClient: alphavantage.NewClient(alphaVantageAPIKey),
	}
}

func (c *Client) GetStockDetails(stockSymbol string) (*StockDetails, error) {
	stockDetails, err := c.alphaVantageClient.GetStockDetails(stockSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get stock details: %w", err)
	}

	value := 22.5 * stockDetails.EPS * stockDetails.BVPS
	grahamNumber := float64(0)
	if value > 0 {
		grahamNumber = math.Sqrt(value)
	}

	var isUndervalued bool
	if grahamNumber > stockDetails.Price {
		isUndervalued = true
	}

	return &StockDetails{
		GrahamNumber:  grahamNumber,
		IsUndervalued: isUndervalued,
		Price:         stockDetails.Price,
		EPS:           stockDetails.EPS,
		BVPS:          stockDetails.BVPS,
		AlphaVantage: &AlphaVantageData{
			GlobalQuote: mapGlobalQuote(stockDetails.GlobalQuote),
			Overview:    mapOverview(stockDetails.Overview),
		},
	}, nil
}
