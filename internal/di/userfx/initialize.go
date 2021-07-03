package userfx

import (
	"gorm.io/gorm"

	"go-practice/pkg/cache"
	"go-practice/pkg/user"
)

func provideUserDBRepository(db *gorm.DB) user.DBRepository {
	return user.NewDBRepository(db)
}

func provideUserRepository(dbRepo user.DBRepository, cache cache.Cache) user.Repository {
	return user.NewRepository(dbRepo, cache)
}
