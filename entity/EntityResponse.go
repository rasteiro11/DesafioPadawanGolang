package entity

import "gorm.io/gorm"

type ExchangeResponse struct {
	gorm.Model
	ValorConvertido string
	SimboloMoeda    string
}
