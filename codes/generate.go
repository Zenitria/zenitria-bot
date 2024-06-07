package codes

import "time"

func GenerateCode(amt int, hrs int, uses int) (string, time.Time) {
	code := createCode()
	expires := createExpires(hrs)

	if checkCode(code) {
		addCode(code, amt, expires, uses)
		return code, expires
	}

	return GenerateCode(amt, hrs, uses)
}
