package iv_models

type DB_DeltaForm struct {
	AccessionNumber       string        `json:"accessionNumber" bson:"accessionNumber"`
	FormClass             string        `json:"formClass" bson:"formClass"`
	PeriodOfReport        string        `json:"periodOfReport" bson:"periodOfReport"`
	AveragePricePerShare  float32       `json:"averagePricePerShare" bson:"averagePricePerShare"`
	NetTotal              float32       `json:"netTotal" bson:"netTotal"`
	SharesTraded          float32       `json:"sharesTraded" bson:"sharesTraded"`
	PostTransactionShares float32       `json:"postTransactionShares" bson:"postTransactionShares"`
	BuyOrSell             string        `json:"buyOrSell" bson:"buyOrSell"`
	Url                   string        `json:"url" bson:"url"`
	DateAdded             string        `json:"dateAdded" bson:"dateAdded"`
	Issuer                DB_Issuer     `json:"issuer" bson:"issuer"`
	Reporters             []DB_Reporter `json:"reporters" bson:"reporters"`
}

type DB_Issuer struct {
	IssuerCik      string `json:"issuerCik" bson:"issuerCik"`
	IssuerName     string `json:"issuerName" bson:"issuerName"`
	IssuerTicker   string `json:"issuerTicker" bson:"issuerTicker"`
	IssuerSector   string `json:"issuerSector" bson:"issuerSector"`
	IssuerIndustry string `json:"issuerIndustry" bson:"issuerIndustry"`
}

type DB_Reporter struct {
	ReporterCik     string `json:"reporterCik" bson:"reporterCik"`
	ReporterName    string `json:"reporterName" bson:"reporterName"`
	ReporterTitle   string `json:"reporterTitle" bson:"reporterTitle"`
	ReporterAddress string `json:"reporterAddress" bson:"reporterAddress"`
}
