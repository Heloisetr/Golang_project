package order

import "context"

type Fetcher interface {
	DeleteAll(ctx context.Context, userID string) error
}
