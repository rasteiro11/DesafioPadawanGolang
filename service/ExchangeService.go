package service

import (
	"server/entity"
	"server/errors"

	"github.com/gin-gonic/gin"
)

type ExchangeService interface {
	StoreExchangeResponse(c *gin.Context) (*entity.ExchangeResponse, *errors.ExchangeError)
}
