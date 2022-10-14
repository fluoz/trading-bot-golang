package binance

import "context"

type Futures interface {
	Connection(ctx context.Context) error
}
