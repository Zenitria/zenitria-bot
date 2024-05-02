package prices

type Price struct {
	Price  float32 `json:"price"`
	Change float32 `json:"change24h"`
}

type Coins struct {
	BTC Price `json:"BTC"`
	BAN Price `json:"BAN"`
	XMR Price `json:"XMR"`
	XNO Price `json:"XNO"`
}
