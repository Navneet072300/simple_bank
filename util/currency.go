package util

const (
	USD = "USD"
	IND = "IND"
	EUR = "EUR"
)

func IsSupportCurrency(currency string) bool{
	switch currency{
		case USD, IND, EUR:
			return true
	}
	return false
}