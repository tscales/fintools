package ratios

func ReturnOnEquity(netIncome, totalAssets, totalLiabilities float32) (float32, error) {
	equity := totalAssets - totalLiabilities
	if equity == 0.0 {
		return 0.0, ErrDivideByZero
	}
	return netIncome / equity, nil
}
