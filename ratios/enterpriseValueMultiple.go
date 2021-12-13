package ratios

func EnterpriseValueMultiple(ebitda,
	pricePerShare,
	sharesOutstanding,
	interestBearingDebt,
	cash float32) (float32, error) {
	if ebitda == 0 {
		return 0, ErrDivideByZero
	}
	ev := EnterpriseValue(pricePerShare, sharesOutstanding, interestBearingDebt, cash)
	return ev / ebitda, nil

}
