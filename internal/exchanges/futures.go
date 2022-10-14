package exchanges

import (
	"context"
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
		err := e.binance.Connection(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
