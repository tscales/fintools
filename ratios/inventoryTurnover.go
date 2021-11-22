package ratios

// InventoryTurnover looks at cost of goods sold and inventory
// to see how many times over a company sold that amount of inventory
func InventoryTurnover(costOfGoodsSold, inventory float32) (float32, error) {
	if inventory == 0 {
		return 0.0, ErrDivideByZero
	}
	return costOfGoodsSold / inventory, nil
}
