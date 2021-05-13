package ad

import (
	"ads/domain"
	"context"
)

type Store interface {
	Create(ctx context.Context, ad domain.Ad) error
	Update(ctx context.Context, adID string, ad domain.Ad) (domain.Ad, error)
	Delete(ctx context.Context, adID string) error
	GetAll(ctx context.Context, userID string) ([]*domain.Ad, error)
	Get(ctx context.Context, adID string) (domain.Ad, error)
	GetByKeys(ctx context.Context, keyword string) ([]*domain.Ad, error)
}
