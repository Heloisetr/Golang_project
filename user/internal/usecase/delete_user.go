package usecase

import (
	"context"
	"epitech/deliveats/user/internal/infrastructure/order"
	"epitech/deliveats/user/internal/infrastructure/user"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DeleteUserCmd func(ctx context.Context, userID string) error

func DeleteUser(store user.Store, order order.Fetcher) DeleteUserCmd {
	return func(ctx context.Context, userID string) error {
		err := order.DeleteAll(ctx, userID)
		if err != nil {
			logrus.Error(err)
			return errors.Wrap(err, "error while deleting order of the user")
		}
		return store.Delete(ctx, userID)
	}
}
