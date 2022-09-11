package iv_models

type SecApiResponse struct {
	Cik                             string   `json:"cik"`
	EntityType                      string   `json:"entityType"`
	Sic                             string   `json:"sic"`
	SicDescription                  string   `json:"sicDescription"`
	IsReporter                      int      `json:"insiderTransactionForOwnerExists"`
	IsIssuer                        int      `json:"insiderTransactionForIssuerExists"`
	Name                            string   `json:"name"`
	Ein                             string   `json:"ein"`
	Tickers                         []string `json:"tickers"`
	Exchanges                       []string `json:"exchanges"`
	Description                     string   `json:"description"`
	Website                         string   `json:"website"`
	InvestorWebsite                 string   `json:"InvestorWebsite"`
	Category                        string   `json:"category"`
	FiscalYearEnd                   string   `json:"fiscalYearEnd"`
	StateOfIncorporation            string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription string   `json:"stateOfIncorporationDescription"`
	Phone                           string   `json:"phone"`
	Flags                           string   `json:"flags"`
	Filings                         filings  `json:"filings"`
}

type filings struct {
	Recent recent `json:"recent"`
	Files  []file `json:"files"`
}

type recent struct {
	AccNums     []string `json:"accessionNumber"`
	Dates       []string `json:"filingDate"`
	ReportDates []string `json:"reportDate"`
	Forms       []string `json:"form"`
	PrimaryDocs []string `json:"primaryDocument"`
}

type file struct {
	Name        string `json:"name"`
	FilingCount int    `json:"filingCount"`
	FilingFrom  string `json:"filingFrom"`
	FilingTo    string `json:"filingTo"`
}
