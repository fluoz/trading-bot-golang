package exchanges

import (
	"context"
	"errors"
	binance "trading-bot/internal/exchanges/binance/futures"
)

type marketFutures struct {
	binance binance.Futures
}

func NewMarketFutures(binance binance.Futures) Futures {
	return &marketFutures{binance: binance}
}

func (e *marketFutures) Connection(ctx context.Context, exchange string) error {
	switch exchange {
	case exchangeBinance:
		return e.binance.Connection(ctx)
	}
	return nil
}

func (e *marketFutures) ServerTime(ctx context.Context, exchange string) (serverTimeResponse, error) {

	switch exchange {
	case exchangeBinance:
		res, err := e.binance.ServerTime(ctx)
		if err != nil {
			return serverTimeResponse{}, err
		}
		return serverTimeResponse{res.ServerTime}, nil

	default:
		return serverTimeResponse{}, errors.New("wrong exchange type")
	}
}
