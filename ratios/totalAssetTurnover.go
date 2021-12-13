package ratios

// TotalAssetTurnover calculates how much is generated in sales
// for every dollar of a companies assets.
func TotalAssetTurnover(sales, totalAssets float32) (float32, error) {
	if totalAssets == 0 {
		return 0.0, ErrDivideByZero
	}
	return sales / totalAssets, nil
}
