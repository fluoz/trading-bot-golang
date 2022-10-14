package exchanges

import "context"

type Futures interface {
	Connection(ctx context.Context, provider string) error
}

const (
	exchangeBinance = `binance`
)
