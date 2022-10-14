package binance

import (
	"context"
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

	_, err = io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return nil
}
