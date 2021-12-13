package ratios

// ReturnOnAssets indicates how much profit a company is generating
// in relation to its total assets.
// a higher number indicates a company is using its assetts efficiently.
func ReturnOnAssets(netIncome, totalAssets float32) (float32, error) {
	if totalAssets == 0.0 {
		return 0.0, ErrDivideByZero
	}
	return netIncome / totalAssets, nil
}
