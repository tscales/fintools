package ratios

func CapitalIntensity(totalAssets, sales float64) (float64, error) {
	if sales == 0 {
		return 0.0, ErrDivideByZero
	}
	return totalAssets / sales, nil
}
