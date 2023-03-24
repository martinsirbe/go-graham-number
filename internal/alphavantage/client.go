package alphavantage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	httpClient *http.Client
	url        string
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		url:    "https://www.alphavantage.co/query?function=%s&symbol=%s&apikey=%s",
		apiKey: apiKey,
	}
}

func (c *Client) GetStockDetails(stockSymbol string) (*StockDetails, error) {
	price, err := c.getGlobalQuote(stockSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get alpha vantage global quote: %w", err)
	}

	eps, bvps, err := c.getOverview(stockSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get alpha vantage overview: %w", err)
	}

	return &StockDetails{
		Price: price,
		EPS:   eps,
		BVPS:  bvps,
	}, err
}

func (c *Client) getGlobalQuote(stockSymbol string) (float64, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(c.url, "GLOBAL_QUOTE", stockSymbol, c.apiKey))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var globalQuote *GlobalQuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&globalQuote); err != nil {
		return 0, err
	}

	return strconv.ParseFloat(globalQuote.GlobalQuote.CurrentPrice, 64)
}

func (c *Client) getOverview(stockSymbol string) (float64, float64, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(c.url, "OVERVIEW", stockSymbol, c.apiKey))
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var overview *OverviewResponse
	if err := json.NewDecoder(resp.Body).Decode(&overview); err != nil {
		return 0, 0, err
	}

	eps, err := strconv.ParseFloat(overview.EPS, 64)
	if err != nil {
		return 0, 0, err
	}

	bvps, err := strconv.ParseFloat(overview.BVPS, 64)
	if err != nil {
		return 0, 0, err
	}

	return eps, bvps, nil
}
