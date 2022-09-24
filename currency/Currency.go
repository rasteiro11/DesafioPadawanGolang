package currency

import "server/errors"

type Currency int

const (
	NONE Currency = iota
	USD
	EUR
	BRL
)

func ParseCurrency(strCurr string) Currency {
	var currType Currency
	switch strCurr {
	case "USD":
		currType = USD
	case "EUR":
		currType = EUR
	case "BRL":
		currType = BRL
	default:
		currType = NONE
		return currType
	}
	return currType
}

func GetCurrencySymbol(curr Currency) (string, *errors.ExchangeError) {
	var currType string
	switch curr {
	case USD:
		currType = "$"
	case EUR:
		currType = "€"
	case BRL:
		currType = "R$"
	default:
		return "", &errors.ExchangeError{Type: "COULD NOT GET CURRENCY SYMBOL"}
	}

	return currType, nil
}
