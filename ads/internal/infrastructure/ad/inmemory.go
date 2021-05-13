package ad

import (
	"ads/domain"
	"context"
)

type InMemory struct {
	Data   map[string]domain.Ad
	UserAd map[string][]string
}

func NewInMemory() *InMemory {
	data := make(map[string]domain.Ad)
	//To get All ad or delete all add => return string tab of all adID
	userAd := make(map[string][]string)
	return &InMemory{Data: data, UserAd: userAd}
}

func (m *InMemory) Create(ctx context.Context, ad domain.Ad) error {
	m.Data[ad.AdID] = ad
	//append in the tab all ads corresponding to the userId corresponding
	m.UserAd[ad.UserID] = append(m.UserAd[ad.UserID], ad.AdID)
	return nil
}

func (m *InMemory) Get(ctx context.Context, adID string) (domain.Ad, error) {
	ad, ok := m.Data[adID]
	if !ok {
		return domain.Ad{}, domain.ErrAdNotFound
	}
	return ad, nil
}

//Update(ctx context.Context, adID string) (domain.Ad, error)

func (m *InMemory) Delete(ctx context.Context, adID string) error {
	return nil
}

//GetAll(ctx context.Context, userID string) ([]domain.Ad, error)
