package ratios

// PriceEarningRatio measures how much investors are willing to pay
// per dollar of a company's current earnings.
func PriceEarningsRatio(netIncome, sharesOutstanding, pricePerShare float32) (float32, error) {
	eps, err := EarningsPerShare(netIncome, sharesOutstanding)

	if err != nil {
		return 0.0, err
	}
	return pricePerShare / eps, nil
}
