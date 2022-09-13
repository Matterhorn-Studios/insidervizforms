package iv_models

type DB_FeaturedIssuer struct {
	Index            int     `bson:"index"`
	Cik              string  `bson:"cik"`
	NetTotal         float32 `bson:"netTotal"`
	NumberOfInsiders int     `bson:"numberOfInsiders"`
	Name             string  `bson:"name"`
	Ticker           string  `bson:"ticker"`
	MostRecentDate   string  `bson:"mostRecentDate"`
	CompanyValue     float32 `bson:"companyValue"`
}
