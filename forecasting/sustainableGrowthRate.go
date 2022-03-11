package forecasting

func SustainableGrowthRate(returnOnEquity, retention float64) (float64, error) {
	rb := returnOnEquity * retention
	if 1-rb == 0 {
		return 0.0, ErrDivideByZero
	}
	return rb/1 - rb, nil
}
