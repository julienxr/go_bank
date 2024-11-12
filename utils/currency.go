package utils

import "strings"

var money map[string]string

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	_, exists := money[strings.ToUpper(currency)]
	if exists {
		return true
	}

	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}

func AddSupportForCurrency(currency string) {
	CUR := strings.ToUpper(currency)

	if money == nil {
		money = make(map[string]string)
	}
	if _, ok := money[CUR]; ok {
		return
	}
	money[CUR] = CUR
}
