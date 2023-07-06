package coingecko

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func getPrices() Coins {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,monero,nano&vs_currencies=usd&include_24hr_change=true"

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, _ := io.ReadAll(res.Body)

	var prices Coins

	json.Unmarshal(body, &prices)

	return prices
}

func worker() {
	for {
		Prices = getPrices()

		time.Sleep(1 * time.Minute)
	}
}
