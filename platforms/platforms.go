package platforms

import (
	"fmt"
	"math"
	"net/http"
	"zenitria-bot/config"
)

func GetXNO() (Platform, string) {
	url := config.GET_XNO_API_URL + "/public/stats"

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return Platform{}, ""
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	platform := parseJson(res.Body)

	change := ((float64(platform.Claims) - float64(platform.YesterdayClaims)) / float64(platform.YesterdayClaims)) * 100
	var percentage string

	if change > 0 {
		percentage = fmt.Sprintf("+%.2f", change)
	} else {
		percentage = fmt.Sprintf("%.2f", change)
	}

	if math.IsInf(change, 1) {
		percentage = "+∞"
	} else if math.IsInf(change, -1) {
		percentage = "-∞"
	}

	percentage += "%"

	return platform, percentage
}

func GetBAN() (Platform, string) {
	url := config.GET_BAN_API_URL + "/public/stats"

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return Platform{}, ""
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	platform := parseJson(res.Body)

	change := ((float64(platform.Claims) - float64(platform.YesterdayClaims)) / float64(platform.YesterdayClaims)) * 100
	var percentage string

	if change > 0 {
		percentage = fmt.Sprintf("+%.2f", change)
	} else {
		percentage = fmt.Sprintf("%.2f", change)
	}

	if math.IsInf(change, 1) {
		percentage = "+∞"
	} else if math.IsInf(change, -1) {
		percentage = "-∞"
	}

	percentage += "%"

	return platform, percentage
}
