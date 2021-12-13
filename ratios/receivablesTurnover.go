package ratios

// ReceivablesTurnover measures how fast a company collects money on their sales
func ReceivablesTurnover(sales, accountsReceivable float32) (float32, error) {
	if accountsReceivable == 0 {
		return 0.0, ErrDivideByZero
	}
	return sales / accountsReceivable, nil

}
