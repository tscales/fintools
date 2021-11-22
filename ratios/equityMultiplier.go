package ratios

func EquityMultiplier(assets, equity float32) (float32, error) {
	if equity == 0 {
		return 0.0, ErrDivideByZero
	}
	return assets / equity, nil
}
