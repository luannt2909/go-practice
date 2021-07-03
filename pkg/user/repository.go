package user

import (
	"context"

	"go-practice/pkg/cache"
)

type Repository interface {
	Create(c context.Context, user *User) error
	GetByID(c context.Context, id uint64) (User, error)
}

type repository struct {
	dbRepo DBRepository
	cache  cache.Cache
}

func (r *repository) Create(c context.Context, user *User) error {
	panic("implement me")
}

func (r *repository) GetByID(c context.Context, id uint64) (User, error) {
	panic("implement me")
}

func NewRepository(dbRepo DBRepository,
	cache cache.Cache) Repository {
	return &repository{
		dbRepo: dbRepo,
		cache:  cache,
	}
}
