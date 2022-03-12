package valuation

import "math"

func futureValue(cash, rate, terms float64) float64 {
	return cash * math.Pow(1+rate, terms)
}
