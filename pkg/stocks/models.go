package stocks

import "github.com/martinsirbe/go-graham-number/internal/alphavantage"

type StockDetails struct {
	GrahamNumber  float64           `json:"graham_number"`
	IsUndervalued bool              `json:"is_undervalued"`
	Price         float64           `json:"price"`
	EPS           float64           `json:"eps"`
	BVPS          float64           `json:"bvps"`
	AlphaVantage  *AlphaVantageData `json:"alpha_vantage"`
}

type AlphaVantageData struct {
	GlobalQuote *GlobalQuote `json:"global_quote"`
	Overview    *Overview    `json:"overview"`
}

type GlobalQuote struct {
	Symbol           string `json:"symbol"`
	Open             string `json:"open"`
	High             string `json:"high"`
	Low              string `json:"low"`
	Price            string `json:"price"`
	Volume           string `json:"volume"`
	LatestTradingDay string `json:"latest_trading_day"`
	PreviousClose    string `json:"previous_close"`
	Change           string `json:"change"`
	ChangePercent    string `json:"change_percent"`
}

func mapGlobalQuote(gq *alphavantage.GlobalQuote) *GlobalQuote {
	return &GlobalQuote{
		Symbol:           gq.Symbol,
		Open:             gq.Open,
		High:             gq.High,
		Low:              gq.Low,
		Price:            gq.Price,
		Volume:           gq.Volume,
		LatestTradingDay: gq.LatestTradingDay,
		PreviousClose:    gq.PreviousClose,
		Change:           gq.Change,
		ChangePercent:    gq.ChangePercent,
	}
}

type Overview struct {
	Symbol                     string `json:"symbol"`
	AssetType                  string `json:"asset_type"`
	Name                       string `json:"name"`
	Description                string `json:"description"`
	CIK                        string `json:"cik"`
	Exchange                   string `json:"exchange"`
	Currency                   string `json:"currency"`
	Country                    string `json:"country"`
	Sector                     string `json:"sector"`
	Industry                   string `json:"industry"`
	Address                    string `json:"address"`
	FiscalYearEnd              string `json:"fiscal_year_end"`
	LatestQuarter              string `json:"latest_quarter"`
	MarketCapitalization       string `json:"market_capitalization"`
	EBITDA                     string `json:"ebitda"`
	PERatio                    string `json:"pe_ratio"`
	PEGRatio                   string `json:"peg_ratio"`
	BookValue                  string `json:"book_value"`
	DividendPerShare           string `json:"dividend_per_share"`
	DividendYield              string `json:"dividend_yield"`
	EPS                        string `json:"eps"`
	RevenuePerShareTTM         string `json:"revenue_per_share_ttm"`
	ProfitMargin               string `json:"profit_margin"`
	OperatingMarginTTM         string `json:"operating_margin_ttm"`
	ReturnOnAssetsTTM          string `json:"return_on_assets_ttm"`
	ReturnOnEquityTTM          string `json:"return_on_equity_ttm"`
	RevenueTTM                 string `json:"revenue_ttm"`
	GrossProfitTTM             string `json:"gross_profit_ttm"`
	DilutedEPSTTM              string `json:"diluted_eps_ttm"`
	QuarterlyEarningsGrowthYOY string `json:"quarterly_earnings_growth_yoy"`
	QuarterlyRevenueGrowthYOY  string `json:"quarterly_revenue_growth_yoy"`
	AnalystTargetPrice         string `json:"analyst_target_price"`
	TrailingPE                 string `json:"trailing_pe"`
	ForwardPE                  string `json:"forward_pe"`
	PriceToSalesRatioTTM       string `json:"price_to_sales_ratio_ttm"`
	PriceToBookRatio           string `json:"price_to_book_ratio"`
	EVToRevenue                string `json:"ev_to_revenue"`
	EVToEBITDA                 string `json:"ev_to_ebitda"`
	Beta                       string `json:"beta"`
	WeekHigh                   string `json:"52_week_high"`
	WeekLow                    string `json:"52_week_low"`
	DayMovingAverage           string `json:"50_day_moving_average"`
	DayMovingAverage1          string `json:"200_day_moving_average"`
	SharesOutstanding          string `json:"shares_outstanding"`
	DividendDate               string `json:"dividend_date"`
	ExDividendDate             string `json:"ex_dividend_date"`
}

func mapOverview(o *alphavantage.Overview) *Overview {
	return &Overview{
		Symbol:                     o.Symbol,
		AssetType:                  o.AssetType,
		Name:                       o.Name,
		Description:                o.Description,
		CIK:                        o.CIK,
		Exchange:                   o.Exchange,
		Currency:                   o.Currency,
		Country:                    o.Country,
		Sector:                     o.Sector,
		Industry:                   o.Industry,
		Address:                    o.Address,
		FiscalYearEnd:              o.FiscalYearEnd,
		LatestQuarter:              o.LatestQuarter,
		MarketCapitalization:       o.MarketCapitalization,
		EBITDA:                     o.EBITDA,
		PERatio:                    o.PERatio,
		PEGRatio:                   o.PEGRatio,
		BookValue:                  o.BookValue,
		DividendPerShare:           o.DividendPerShare,
		DividendYield:              o.DividendYield,
		EPS:                        o.EPS,
		RevenuePerShareTTM:         o.RevenuePerShareTTM,
		ProfitMargin:               o.ProfitMargin,
		OperatingMarginTTM:         o.OperatingMarginTTM,
		ReturnOnAssetsTTM:          o.ReturnOnAssetsTTM,
		ReturnOnEquityTTM:          o.ReturnOnEquityTTM,
		RevenueTTM:                 o.RevenueTTM,
		GrossProfitTTM:             o.GrossProfitTTM,
		DilutedEPSTTM:              o.DilutedEPSTTM,
		QuarterlyEarningsGrowthYOY: o.QuarterlyEarningsGrowthYOY,
		QuarterlyRevenueGrowthYOY:  o.QuarterlyRevenueGrowthYOY,
		AnalystTargetPrice:         o.AnalystTargetPrice,
		TrailingPE:                 o.TrailingPE,
		ForwardPE:                  o.ForwardPE,
		PriceToSalesRatioTTM:       o.PriceToSalesRatioTTM,
		PriceToBookRatio:           o.PriceToBookRatio,
		EVToRevenue:                o.EVToRevenue,
		EVToEBITDA:                 o.EVToEBITDA,
		Beta:                       o.Beta,
		WeekHigh:                   o.WeekHigh,
		WeekLow:                    o.WeekLow,
		DayMovingAverage:           o.DayMovingAverage,
		DayMovingAverage1:          o.DayMovingAverage1,
		SharesOutstanding:          o.SharesOutstanding,
		DividendDate:               o.DividendDate,
		ExDividendDate:             o.ExDividendDate,
	}
}
