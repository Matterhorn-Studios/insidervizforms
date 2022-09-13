package iv_models

type DB_Sector struct {
	Id             DB_Sector_Id           `json:"_id" bson:"_id"`
	HistoricalData []DB_Sector_Historical `json:"historicalData" bson:"historicalData"`
}

type DB_Sector_Id struct {
	Name string `json:"name" bson:"name"`
}

type DB_Sector_Historical struct {
	Date        string  `json:"date" bson:"date"`
	Total       float64 `json:"total" bson:"total"`
	TotalBought float64 `json:"totalBought" bson:"totalBought"`
	TotalSold   float64 `json:"totalSold" bson:"totalSold"`
	BuyOrSell   string  `json:"buyOrSell" bson:"buyOrSell"`
}
