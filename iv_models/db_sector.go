package iv_models

type DB_Sector struct {
	Name           string                 `json:"name" bson:"name"`
	HistoricalData []DB_Sector_Historical `json:"historicalData" bson:"historicalData"`
}

type DB_Sector_Historical struct {
	Date        string  `json:"date" bson:"date"`
	Total       float32 `json:"total" bson:"total"`
	TotalBought float32 `json:"totalBought" bson:"totalBought"`
	TotalSold   float32 `json:"totalSold" bson:"totalSold"`
	BuyOrSell   float32 `json:"buyOrSell" bson:"buyOrSell"`
}
