package account

import (
	"account/domain"
	"context"
)

type InMemory struct {
	Data map[string]domain.Account
}

func NewInMemory() *InMemory {
	data := make(map[string]domain.Account)
	return &InMemory{Data: data}
}

func (m *InMemory) Create(_ context.Context, account domain.Account) error {
	_, ok := m.Data[account.AccountID]
	if ok {
		return domain.ErrEmailAlreadyUsed
	}
	m.Data[account.AccountID] = account
	return nil
}

func (m *InMemory) Get(_ context.Context, accountID string) (domain.Account, error) {
	account, ok := m.Data[accountID]
	if !ok {
		return domain.Account{}, domain.ErrAccountNotFound
	}

	return account, nil
}
