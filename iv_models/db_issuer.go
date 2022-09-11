package iv_models

type DB_Issuer_Doc struct {
	Name                 string      `json:"name" bson:"name"`
	Cik                  string      `json:"cik" bson:"cik"`
	Sic                  string      `json:"sic" bson:"sic"`
	SicDescription       string      `json:"secDescription" bson:"secDescription"`
	Ein                  string      `json:"ein" bson:"ein"`
	Tickers              []string    `json:"tickers" bson:"tickers"`
	Exchanges            []string    `json:"exchanges" bson:"exchanges"`
	FiscalYearEnd        string      `json:"fiscalYearEnd" bson:"fiscalYearEnd"`
	StateOfIncorporation string      `json:"stateOfIncorporation" bson:"stateOfIncorporation"`
	Phone                string      `json:"phone" bson:"phone"`
	StockData            []StockData `json:"stockData" bson:"stockData"`
	StockDataSplit       bool        `json:"stockDataSplit" bson:"stockDataSplit"`
	Splits               []Split     `json:"splits" bson:"splits"`
	Sector               string      `json:"sector" bson:"sector"`
	Industry             string      `json:"industry" bson:"industry"`
}

type StockData struct {
	Date   string  `json:"date" bson:"date"`
	Close  float64 `json:"close" bson:"close"`
	Volume int     `json:"volume" bson:"volume"`
}

type Split struct {
	Date  string `json:"date" bson:"date"`
	Split string `json:"split" bson:"split"`
}
