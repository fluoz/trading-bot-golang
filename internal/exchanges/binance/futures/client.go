package binance

import (
	"context"
	"io"
	"net/http"
)

type binanceFutures struct {
}

func NewBinanceFutures() Futures {
	return &binanceFutures{}
}

func (b *binanceFutures) Connection(ctx context.Context) error {

	req, err := http.NewRequest(http.MethodGet, "/fapi/v1/ping", nil)

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
