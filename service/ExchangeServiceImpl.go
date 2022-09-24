package service

import (
	"fmt"
	"server/currency"
	"server/entity"
	"server/errors"
	"server/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExchangeServiceImpl struct {
	repo repository.ExchangeResponseRepository
}

func NewExchangeResponseFromParams(c *gin.Context) (*entity.ExchangeResponse, *errors.ExchangeError) {
	a, amount_err := strconv.ParseFloat(c.Param("amount"), 64)
	if amount_err != nil {
		return nil, &errors.ExchangeError{Type: "PARSING AMOUNT ERROR"}
	}
	r, rate_err := strconv.ParseFloat(c.Param("rate"), 64)
	if rate_err != nil {
		return nil, &errors.ExchangeError{Type: "PARSING RATE ERROR"}
	}

	toCurr := currency.ParseCurrency(c.Param("to"))

	respSymbol, err := currency.GetCurrencySymbol(toCurr)
	if err != nil {
		return nil, err
	}

	return &entity.ExchangeResponse{ValorConvertido: fmt.Sprintf("%f", a*r), SimboloMoeda: respSymbol}, nil
}

func (esi *ExchangeServiceImpl) StoreExchangeResponse(c *gin.Context) (*entity.ExchangeResponse, *errors.ExchangeError) {
	exchangeResponse, err := NewExchangeResponseFromParams(c)
	if err != nil {
		return nil, err
	}
	esi.repo.StoreExchangeResponse(exchangeResponse)
	return exchangeResponse, nil
}

func NewExchangeServiceImpl(repo repository.ExchangeResponseRepository) *ExchangeServiceImpl {
	return &ExchangeServiceImpl{repo}
}
