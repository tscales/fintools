package ratios

// TotalDebtRatio calculates what percentage of a companies assets
// are provided by debt.
func TotalDebtRatio(assets, equity float32) (float32, error) {
	if assets == 0 {
		return 0.0, ErrDivideByZero
	}
	return (assets - equity) / assets, nil
}
