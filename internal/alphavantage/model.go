package alphavantage

type StockDetails struct {
	Price float64
	EPS   float64
	BVPS  float64
}

type GlobalQuoteResponse struct {
	GlobalQuote struct {
		CurrentPrice string `json:"05. price"`
	} `json:"Global Quote"`
}

type OverviewResponse struct {
	EPS  string `json:"EPS"`
	BVPS string `json:"BookValue"`
}
