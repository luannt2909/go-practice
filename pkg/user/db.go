package user

import (
	"context"

	"gorm.io/gorm"
)

type DBRepository interface {
	GetByID(c context.Context, id uint64) (User, error)
	Create(c context.Context, user *User) error
}

type dbRepo struct {
	db *gorm.DB
}

func (r *dbRepo) Create(c context.Context, user *User) error {
	return r.db.Create(user).Error
}

func (r *dbRepo) GetByID(c context.Context, id uint64) (user User, err error) {
	err = r.db.First(&user, id).Error
	return
}

func NewDBRepository(db *gorm.DB) DBRepository {
	return &dbRepo{db: db}
}
