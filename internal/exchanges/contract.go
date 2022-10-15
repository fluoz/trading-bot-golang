package exchanges

import "context"

type Futures interface {
	Connection(ctx context.Context, exchange string) error
	ServerTime(ctx context.Context, exchange string) (serverTimeResponse, error)
}

const (
	exchangeBinance = `binance`
)
