package alphavantage

func (c *Client) SetURL(url string) {
	c.url = url + "/query?function=%s&symbol=%s&apikey=%s"
}
