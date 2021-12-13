package ratios

// DupontIdentity breaks the Return On Equity ratio in to 3 measurable items
// 1: Operating efficiency (profit Margin)
// 2: Use of Assets efficiency (total asset turnover)
// 3: Financial leverage (equity multiplier)
type DupontIdentity struct {
	ROE                float32
	ProfitMargin       float32
	TotalAssetTurnOver float32
	EquityMultiplier   float32
}

func GetDupontIdentity(netIncome, sales, totalAssets, totalEquity float32) DupontIdentity {
	pm, err := ProfitMargin(netIncome, sales)
	if err != nil {
		pm = 0
	}
	tat, err := TotalAssetTurnover(sales, totalAssets)
	if err != nil {
		tat = 0
	}
	em, err := EquityMultiplier(totalAssets, totalEquity)
	if err != nil {
		em = 0
	}
	roe, err := ReturnOnEquity(netIncome, totalAssets, 0)
	if err != nil {
		roe = 0
	}

	return DupontIdentity{
		ROE:                roe,
		ProfitMargin:       pm,
		TotalAssetTurnOver: tat,
		EquityMultiplier:   em,
	}

}
