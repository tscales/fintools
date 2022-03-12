package valuation

//NetPresentValue takes the cost
func NetPresentValue(cost, presentValue float64) float64 {
	return presentValue - cost
}
