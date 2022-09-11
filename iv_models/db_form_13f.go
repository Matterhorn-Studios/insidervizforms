package iv_models

type DB_Form13F_Base struct {
	Cik             string               `json:"cik" bson:"cik"`
	ReporterName    string               `json:"reporterName" bson:"reporterName"`
	AccessionNumber string               `json:"accessionNumber" bson:"accessionNumber"`
	PeriodOfReport  string               `json:"periodOfReport" bson:"periodOfReport"`
	Url             string               `json:"url" bson:"url"`
	TotalHeld       float32              `json:"totalHeld" bson:"totalHeld"`
	Holdings        []DB_Form13F_Holding `json:"holdings" bson:"holdings"`
}

type DB_Form13F_ISSUER struct {
	Cik           string               `json:"cik" bson:"cik"`
	IssuerName    string               `json:"issuerName" bson:"issuerName"`
	IssuerTickers []string             `json:"issuerTickers" bson:"issuerTickers"`
	Holdings      []DB_Form13F_Holding `json:"holdings" bson:"holdings"`
}

// this can be an issuer or a reporter
type DB_Form13F_Holding struct {
	Name     string  `json:"name" bson:"name"`
	Shares   float32 `json:"shares" bson:"shares"`
	NetTotal float32 `json:"netTotal" bson:"netTotal"`
	Cik      string  `json:"cik" bson:"cik"`
}
