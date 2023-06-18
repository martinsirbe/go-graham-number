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
	gqResp, err := c.httpClient.Get(fmt.Sprintf(c.url, "GLOBAL_QUOTE", stockSymbol, c.apiKey))
	if err != nil {
		return nil, err
	}
	defer gqResp.Body.Close()

	if gqResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got %d", gqResp.StatusCode)
	}

	var globalQuoteResponse struct {
		GlobalQuote *GlobalQuote `json:"Global Quote"`
	}
	if err := json.NewDecoder(gqResp.Body).Decode(&globalQuoteResponse); err != nil {
		return nil, fmt.Errorf("failed to decpde alpha vantage global quote: %w", err)
	}

	price, err := strconv.ParseFloat(globalQuoteResponse.GlobalQuote.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price from alpha vantage global quote: %w", err)
	}

	oResp, err := c.httpClient.Get(fmt.Sprintf(c.url, "OVERVIEW", stockSymbol, c.apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to get alpha vantage overview: %w", err)
	}
	defer oResp.Body.Close()

	if oResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got %d", oResp.StatusCode)
	}

	var overview *Overview
	if err := json.NewDecoder(oResp.Body).Decode(&overview); err != nil {
		return nil, err
	}

	eps, err := strconv.ParseFloat(overview.EPS, 64)
	if err != nil {
		return nil, err
	}

	bvps, err := strconv.ParseFloat(overview.BookValue, 64)
	if err != nil {
		return nil, err
	}

	return &StockDetails{
		Price:       price,
		EPS:         eps,
		BVPS:        bvps,
		GlobalQuote: globalQuoteResponse.GlobalQuote,
		Overview:    overview,
	}, err
}
