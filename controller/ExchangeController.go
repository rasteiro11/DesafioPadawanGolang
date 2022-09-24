package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/service"
)

type ExchangeController struct {
	router          *gin.Engine
	exchangeService service.ExchangeService
}

func NewExchangeController(router *gin.Engine, exchangeService service.ExchangeService) *ExchangeController {
	return &ExchangeController{router, exchangeService}
}

func (ec *ExchangeController) Exchange(c *gin.Context) {
	exchangeResponse, err := ec.exchangeService.StoreExchangeResponse(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Type,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"simboloMoeda": exchangeResponse.SimboloMoeda, "valorConvertido": exchangeResponse.ValorConvertido})
	}
}

func (ec *ExchangeController) MountRoutes() {
	ec.router.GET("/exchange/:amount/:from/:to/:rate", ec.Exchange)
}
