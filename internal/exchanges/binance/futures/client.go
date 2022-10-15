package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	applications "trading-bot/internal/applications"
)

type clientFutures struct {
	config *applications.Config
}

func NewClientFutures(config *applications.Config) Futures {
	return &clientFutures{config: config}
}

func (b *clientFutures) Connection(ctx context.Context) error {

	url := fmt.Sprintf("%s%s", b.config.Exchanges.Binance.Futures.BaseUrl, b.config.Exchanges.Binance.Futures.PathConnectivity)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

func (b *clientFutures) ServerTime(ctx context.Context) (serverTimeResponse, error) {
	response := serverTimeResponse{}
	url := fmt.Sprintf("%s%s", b.config.Exchanges.Binance.Futures.BaseUrl, b.config.Exchanges.Binance.Futures.PathServerTime)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return response, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return response, err
	}

	if res.StatusCode != http.StatusOK {
		return response, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, err
}
