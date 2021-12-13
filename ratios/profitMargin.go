package ratios

// ProfitMargin calculates how much income is generated
// for every dollar in sales
func ProfitMargin(netIncome, sales float32) (float32, error) {
	if sales == 0.0 {
		return 0.0, ErrDivideByZero
	}
	return netIncome / sales, nil
}
