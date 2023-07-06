package platforms

import (
	"fmt"
	"net/http"
)

func GetXNO() (Platform, string) {
	url := "https://api.get-xno.com/global/stats"

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

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

	percentage += "%"

	return platform, percentage
}
