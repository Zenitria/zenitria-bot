package platforms

type Platform struct {
	Days            int     `json:"days"`
	Users           int     `json:"users"`
	Withdrawn       float64 `json:"withdrawn"`
	YesterdayClaims int     `json:"yesterdayClaims"`
	Claims          int     `json:"claims"`
}
