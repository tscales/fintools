package ratios

// CurrentRatio divides a company's current assets buy their current liaibilites.
//
// A current ratio > 1 means that a company is capable of paying off its short term debts.
//
// Current assets: cash or inventory a company wishes to turn in to cash
// in the next 12 months.
//
// Current liabilities: debts that must be repaid within 12 months.
// this can include loans, rent, payroll, and taxes.

// You can find these numbers on a company's balance sheet
func CurrentRatio(assets float32, liabilities float32) (float32, error) {
	if liabilities == 0 {
		return 0.0, ErrDivideByZero
	}
	return assets / liabilities, nil
}
