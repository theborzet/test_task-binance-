package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BinanceClient struct {
	APIURL string
}

type result struct {
	Price string `json:"price"`
}

func NewBinanceClient(apiURL string) *BinanceClient {
	return &BinanceClient{
		APIURL: apiURL,
	}
}

func (c *BinanceClient) GetTickerPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("%s%sUSDT", c.APIURL, ticker)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
