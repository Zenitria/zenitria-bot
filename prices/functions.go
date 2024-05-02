package prices

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"zenitria-bot/config"
)

func getPrices() Coins {
	url := config.ZENITRIA_API_URL + "/price/ban,btc,xmr,xno"

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return Coins{}
	}

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
