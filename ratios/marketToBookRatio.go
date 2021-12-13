package ratios

func MarketToBookRatio(pricePerShare, totalEquity, sharesOutstanding float32) (float32, error) {
	if sharesOutstanding == 0.0 {
		return 0.0, ErrDivideByZero
	}
	bookValue := totalEquity / sharesOutstanding

	return pricePerShare / bookValue, nil

}
