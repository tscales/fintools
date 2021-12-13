package ratios

// EnterpriseValue is similar to market capitalization. It calculates how much it would cost to buy all outstanding shares
// and pay off a companies debt.
func EnterpriseValue(pricePerShare, sharesOutstanding, interestbearingdebt, cash float32) float32 {
	return MarketCapitalization(pricePerShare, sharesOutstanding) + interestbearingdebt - cash
}
