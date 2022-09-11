package insidervizforms

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Matterhorn-Studios/insiderviz-forms/iv_models"
)

// returns a map with the form uris, the issuer doc, the reporter doc, the type of entity, and an error
func GetFormsFromJSON(path string, startDate string) (map[string][]string, iv_models.DB_Issuer_Doc, iv_models.DB_Reporter_Doc, string, error) {

	// create the reporter and issuer
	issuer := iv_models.DB_Issuer_Doc{}
	reporter := iv_models.DB_Reporter_Doc{}
	typeOfEntity := "Both"

	// read the json file
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, issuer, reporter, typeOfEntity, err
	}
	defer jsonFile.Close()

	// decode the json file
	var apiResponse iv_models.SecApiResponse
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return nil, issuer, reporter, typeOfEntity, err
	}

	// setup the cik
	cik := apiResponse.Cik
	for len(cik) < 10 {
		cik = "0" + cik
	}

	// check the type of entity
	if apiResponse.IsIssuer == 1 && apiResponse.IsReporter == 0 {
		typeOfEntity = "Issuer"
	} else if apiResponse.IsIssuer == 0 && apiResponse.IsReporter == 1 {
		typeOfEntity = "Reporter"
	}

	// parse the objects
	if typeOfEntity == "Issuer" || typeOfEntity == "Both" {
		// parse the issuer
		issuer.Name = apiResponse.Name
		issuer.Cik = cik
		issuer.Sic = apiResponse.Sic
		issuer.SicDescription = apiResponse.SicDescription
		issuer.Ein = apiResponse.Ein
		issuer.Tickers = apiResponse.Tickers
		issuer.Exchanges = apiResponse.Exchanges
		issuer.FiscalYearEnd = apiResponse.FiscalYearEnd
		issuer.StateOfIncorporation = apiResponse.StateOfIncorporation
		issuer.Phone = apiResponse.Phone
	}
	if typeOfEntity == "Reporter" || typeOfEntity == "Both" {
		// parse the reporter
		reporter.Name = apiResponse.Name
		reporter.Cik = cik
		reporter.IsCongressman = false
		reporter.Party = "Unknown"
	}

	// create the map
	m := make(map[string][]string)
	m["form13Fs"] = []string{}
	m["form4s"] = []string{}
	m["form4s_accNums"] = []string{}

	// loop through the filings
	recent := apiResponse.Filings.Recent
	for i := range recent.AccNums {
		// make sure the date is large enough
		if recent.ReportDates[i] > startDate {
			// check the form type
			if recent.Forms[i] == "4" {
				// get the acc number for the url
				accNumUrl := strings.ReplaceAll(apiResponse.Filings.Recent.AccNums[i], "-", "")

				// get the primary document for the url
				primaryDoc := strings.SplitAfter(recent.PrimaryDocs[i], "/")

				// build the url
				uri := "https://www.sec.gov/Archives/edgar/data/" + apiResponse.Cik + "/" + accNumUrl + "/" + primaryDoc[len(primaryDoc)-1]

				// add to the form4s map
				formSlice, exists := m["form4s"]
				if exists {
					formSlice = append(formSlice, uri)
					m["form4s"] = formSlice
				} else {
					m["form4s"] = []string{uri}
				}

				// add the accNum to map
				accSlice, exists := m["form4s_accNums"]
				if exists {
					accSlice = append(accSlice, apiResponse.Filings.Recent.AccNums[i])
					m["form4s_accNums"] = accSlice
				} else {
					m["form4s_accNums"] = []string{apiResponse.Filings.Recent.AccNums[i]}
				}
			} else if recent.Forms[i] == "4/A" {
				// get the acc number for the url
				accNumUrl := strings.ReplaceAll(apiResponse.Filings.Recent.AccNums[i], "-", "")

				// get the primary document for the url
				primaryDoc := strings.SplitAfter(recent.PrimaryDocs[i], "/")

				// build the url
				uri := "https://www.sec.gov/Archives/edgar/data/" + apiResponse.Cik + "/" + accNumUrl + "/" + primaryDoc[len(primaryDoc)-1]

				// add to the form4s map
				formSlice, exists := m["form4s"]
				if exists {
					formSlice = append(formSlice, uri)
					m["form4-amendments"] = formSlice
				} else {
					m["form4-amendments"] = []string{uri}
				}

				// add the accNum to map
				accSlice, exists := m["form4s_accNums"]
				if exists {
					accSlice = append(accSlice, apiResponse.Filings.Recent.AccNums[i])
					m["form4-amendments_accNums"] = accSlice
				} else {
					m["form4-amendments_accNums"] = []string{apiResponse.Filings.Recent.AccNums[i]}
				}

			} else if recent.Forms[i] == "13F-HR" {
				// get the acc number for the url
				accNumUrl := strings.ReplaceAll(apiResponse.Filings.Recent.AccNums[i], "-", "")

				// build the url
				uri := "https://www.sec.gov/Archives/edgar/data/" + apiResponse.Cik + "/" + accNumUrl

				// add to the form13Fs map
				formSlice, exists := m["form13Fs"]
				if exists {
					formSlice = append(formSlice, uri)
					m["form13Fs"] = formSlice
				} else {
					m["form13Fs"] = []string{uri}
				}

				// add the acc num
				accSlice, exists := m["form13Fs_accNums"]
				if exists {
					accSlice = append(accSlice, apiResponse.Filings.Recent.AccNums[i])
					m["form13Fs_accNums"] = accSlice
				} else {
					m["form13Fs_accNums"] = []string{apiResponse.Filings.Recent.AccNums[i]}
				}
			}
		}
	}

	return m, issuer, reporter, typeOfEntity, nil

}
