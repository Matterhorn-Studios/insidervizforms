package iv_models

type DB_SentimentDay struct {
	Buy  float32 `bson:"buy"`
	Sell float32 `bson:"sell"`
	Date string  `bson:"date"`
}
