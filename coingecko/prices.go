package coingecko

var (
	Prices Coins
)

func Init() {
	go worker()
}
