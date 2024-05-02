package prices

var (
	Prices Coins
)

func Init() {
	go worker()
}
