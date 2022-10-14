package exchanges

import "context"

type Futures interface {
	Connection(ctx context.Context, exchange string) error
}

const (
	exchangeBinance = `binance`
)
