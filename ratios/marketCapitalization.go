package ratios

func MarketCapitalization(pricePerShare, sharesOutstanding float32) float32 {
	return pricePerShare * sharesOutstanding
}
