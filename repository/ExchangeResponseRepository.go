package repository

import (
	"server/entity"

	"gorm.io/gorm"
)

type ExchangeResponseRepository interface {
	StoreExchangeResponse(ecr *entity.ExchangeResponse)
}

type ExchangeResponseRepositoryImpl struct {
	db *gorm.DB
}

func NewExchangeResponseRepositoryImpl(db *gorm.DB) *ExchangeResponseRepositoryImpl {
	return &ExchangeResponseRepositoryImpl{db}
}

func (repo *ExchangeResponseRepositoryImpl) StoreExchangeResponse(ecr *entity.ExchangeResponse) {
	repo.db.Create(ecr)
}
