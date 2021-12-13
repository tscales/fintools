package ratios

func EbitdaMargin(ebit, depreciation, amortization, sales float32) (float32, error) {
	if sales == 0.0 {
		return 0.0, ErrDivideByZero
	}
	return (ebit + depreciation + amortization) / sales, nil
}
