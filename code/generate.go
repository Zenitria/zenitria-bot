package code

func GenerateCode(amt int, hrs int, uses int) string {
	code := createCode()
	expires := createExpires(hrs)

	if checkCode(code) {
		addCode(code, amt, expires, uses)
		return code
	}

	return GenerateCode(amt, hrs, uses)
}
