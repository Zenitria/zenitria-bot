package coingecko

type Price struct {
	Price  float64 `json:"usd"`
	Change float64 `json:"usd_24h_change"`
}

type Coins struct {
	Bitcoin Price `json:"bitcoin"`
	Monero  Price `json:"monero"`
	Nano    Price `json:"nano"`
}
