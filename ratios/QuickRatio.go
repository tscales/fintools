package ratios

// QuickRatio is similar to the CurrentRatio except inventory is subtracted.
//
// Inventory is one of the least liquid assets a company owns. Subtracting it
// from other assets gives a better picture of just how quickly a company can pay debts.
func QuickRatio(assets float32, inventory float32, liabilities float32) float32 {
	return (assets - inventory) / liabilities
}
