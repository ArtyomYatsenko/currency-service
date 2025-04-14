package currency

import (
	"encoding/json"
	"github.com/ArtyomYatsenko/currency/internal/config"
	"github.com/ArtyomYatsenko/currency/internal/dto"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const currencyUrl = "https://latest.currency-api.pages.dev/v1/currencies/rub.json"

type Currency struct {
	baseURL    *url.URL
	httpClient *http.Client
	logger     *zap.Logger
}

func NewHttpClient(configHttp config.HttpClient, logger *zap.Logger) (*Currency, error) {
	parseURL, err := url.Parse(currencyUrl)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: configHttp.Timeout * time.Second,
	}
	return &Currency{
		baseURL:    parseURL,
		httpClient: client,
		logger:     logger,
	}, nil
}

func (c *Currency) FetchData() (dto.Currency, error) {

	data := dto.Currency{}

	urlStr := c.baseURL.String()

	c.logger.Info("making request", zap.String("url", urlStr))

	resp, err := c.httpClient.Get(urlStr)

	if err != nil {
		return data, err
	}

	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			c.logger.Error("resp body close", zap.Error(errClose))
		}
	}()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		log.Printf("json unmarshal: %s", err)
		return data, err
	}

	return data, nil

}
