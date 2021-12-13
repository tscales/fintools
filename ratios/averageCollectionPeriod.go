package ratios

// AverageCollectionPeriod indicates the average number of days it takes
// for a company to collect on their sales.
func AverageCollectionPeriod(sales, accountsReceivable float32) (float32, error) {
	turnover, err := ReceivablesTurnover(sales, accountsReceivable)

	if err != nil {
		return 0.0, err
	}
	return 365 / turnover, nil
}
