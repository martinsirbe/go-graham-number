# Graham Number
The [Graham Number][graham-number] is a command-line interface (CLI) that calculates the Graham Number for 
the given stock symbol.

The Graham Number is used to determine the intrinsic value of a stock. It is named after [Benjamin Graham][benjamin-graham], the father of 
value investing. It's calculated using the company's earnings per share ([EPS][eps]) and book value per share ([BVPS][bvps]). The Graham 
Number formula is `sqrt(22.5 * EPS * BVPS)`.  
The idea behind the Graham Number is that it provides a rough estimate of what a company's stock should be worth based 
on its earnings and assets. If the Graham Number exceeds the current stock price, the stock may be undervalued; otherwise, 
it might not be a good investment opportunity.  
It's important to note that the Graham Number is just one of many metrics used in fundamental analysis and should not be 
relied on solely to make investment decisions. Other factors, such as the company's financial health, industry trends, 
and macroeconomic factors, should also be considered when making investment decisions.

## Installation

Before installing the Graham Number CLI Tool, you must have Go installed on your machine. You can download and install 
the latest version of Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

```shell
go install github.com/martinsirbe/go-graham-number/cmd/gn@latest
```

## Usage
The Graham Number CLI Tool requires a stock symbol as an argument and [The Alpha Vantage][alpha-vantage] 
API key passed via the `ALPHA_VANTAGE_API_KEY` environment variable.

```bash
gh SYMBOL
```

For example, to calculate the Graham Number for Meta Platforms Inc. (META):
```bash
gh META
```

The output will display the BVPS, EPS, and calculated Graham Number for the stock, as well as a comparison to the current price of the stock:
```
META overvalued 97.97 [P: 204.28, EPS: 8.87, BVPS: 48.09]
```
In this example, the Graham Number is lower than the current price of the stock, indicating that the stock may be overvalued.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENCE) file for details.

---
[graham-number]: https://www.investopedia.com/terms/g/graham-number.asp
[alpha-vantage]: https://www.alphavantage.co/documentation
[benjamin-graham]: https://www.investopedia.com/terms/b/bengraham.asp
[eps]: https://www.investopedia.com/terms/e/eps.asp
[bvps]: https://www.investopedia.com/terms/b/bvps.asp
