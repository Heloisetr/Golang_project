package user

import (
	"context"
	"epitech/deliveats/user/domain"
)

type InMemory struct {
	Data map[string]domain.User
}

func NewInMemory() *InMemory {
	data := make(map[string]domain.User)
	return &InMemory{Data: data}
}

func (m *InMemory) Create(_ context.Context, user domain.User) error {
	_, ok := m.Data[user.UserID]
	if ok {
		return domain.ErrUserAlreadyExist
	}
	m.Data[user.UserID] = user
	return nil
}

func (m *InMemory) Get(_ context.Context, userID string) (domain.User, error) {
	user, ok := m.Data[userID]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}

func (m *InMemory) Delete(_ context.Context, userID string) error {
	_, ok := m.Data[userID]
	if !ok {
		return domain.ErrUserNotFound
	}
	delete(m.Data, userID)
	return nil
}
