package ratios

// DaysSaleOfInventory tells us, given the current
// amount of inventory, how many days it takes to sell
func DaysSaleOfInventory(costOfGoodsSold, inventory float32) (float32, error) {
	turnover, err := InventoryTurnover(costOfGoodsSold, inventory)

	if err != nil {
		return 0.0, err
	}
	return 365 / turnover, nil
}
