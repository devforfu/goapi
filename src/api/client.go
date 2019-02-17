package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Client struct {
    BaseURL string
}

func (c *Client) FetchPriceList() (prices Database, err error) {
    resp, err := c.get("/list")
    prices = make(Database)
    for k, v := range resp { prices[k] = Dollars(v.(float64)) }
    return prices, err
}

func (c *Client) get(endpoint string) (result Response, err error) {
    resp, err := http.Get(c.url(endpoint))
    if err != nil { return }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(&result)
    return result, err
}

func (c *Client) url(endpoint string) string {
    return fmt.Sprintf("%s%s", c.BaseURL, endpoint)
}
