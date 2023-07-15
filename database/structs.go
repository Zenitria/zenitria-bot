package database

type User struct {
	ID          string  `bson:"_id"`
	Level       int     `bson:"level"`
	XP          int     `bson:"xp"`
	NextLevelXP int     `bson:"nextLevelXP"`
	Warnings    int     `bson:"warnings"`
	Diamonds    float32 `bson:"diamonds"`
}

type Code struct {
	Code    string   `bson:"code"`
	Amount  int      `bson:"amount"`
	Expires string   `bson:"expires"`
	Uses    int      `bson:"uses"`
	Used    int      `bson:"used"`
	Users   []string `bson:"users"`
	IPs     []string `bson:"ips"`
}

type Channel struct {
	ID string `bson:"_id"`
}
