package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Exchange struct {
	amount float64
	from   string
	to     string
	rate   float64
}

func NewExchange(amount float64, from string, to string, rate float64) *Exchange {
	return &Exchange{amount, from, to, rate}
}

type ExchangeError struct {
	Type string
}

func (e *ExchangeError) Error() string {
	return e.Type
}

func NewExchangeFromParams(c *gin.Context) (*Exchange, *ExchangeError) {
	a, amount_err := strconv.ParseFloat(c.Param("amount"), 64)
	if amount_err != nil {
		return nil, &ExchangeError{Type: "PARSING AMOUNT ERROR"}
	}
	r, rate_err := strconv.ParseFloat(c.Param("rate"), 64)
	if rate_err != nil {
		return nil, &ExchangeError{Type: "PARSING RATE ERROR"}
	}
	return NewExchange(a, c.Param("from"), c.Param("to"), r), nil
}

func exchange(c *gin.Context) {
	excg, err := NewExchangeFromParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Type,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"amout": excg.amount,
			"from":  excg.from,
			"to":    excg.to,
			"rate":  excg.rate,
		})

	}
}

func main() {
	r := gin.Default()

	r.GET("/exchange/:amount/:from/:to/:rate", exchange)

	r.Run()

}
