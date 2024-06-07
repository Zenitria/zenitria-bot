package prices

type Price struct {
	Price  float64 `json:"price"`
	Change float64 `json:"change24h"`
}

type Coins struct {
	BTC Price `json:"BTC"`
	BAN Price `json:"BAN"`
	XMR Price `json:"XMR"`
	XNO Price `json:"XNO"`
}
