package forecasting

func ExternalFinancingNeeded(totalSales,
	totalAssets,
	currentLiabilities,
	profitMargin,
	projectedSales,
	divPayoutRatio float64) (float64, error) {
	if totalSales == 0.0 {
		return 0.0, ErrDivideByZero // consolodate this and ratio error
	}
	as := totalAssets / totalSales
	ls := currentLiabilities / totalSales
	changeInSales := projectedSales - totalSales

	return as*changeInSales - ls*changeInSales - profitMargin*projectedSales*(1-divPayoutRatio), nil
}
