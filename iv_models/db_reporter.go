package iv_models

type DB_Reporter_Doc struct {
	Name          string  `json:"name" bson:"name"`
	Cik           string  `json:"cik" bson:"cik"`
	IsCongressman bool    `json:"isCongressman" bson:"isCongressman"`
	Party         string  `json:"party" bson:"party"`
	IsInst        bool    `json:"isInst" bson:"isInst" default:"false"`
	Last13FTotal  float64 `json:"last13FTotal" bson:"last13FTotal" default:"0"`
}
