package insidervizforms

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Matterhorn-Studios/insidervizforms/iv_models"
	"github.com/PuerkitoBio/goquery"
)

func GetThirteenFFromUrl(url string, accNum string, userAgent string) (iv_models.DB_Form13F_Base, error) {
	var thirteenF iv_models.DB_Form13F_Base

	// create the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return thirteenF, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Host = "www.sec.gov"

	// create the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return thirteenF, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return thirteenF, err
	}

	// find the links
	curIdx := 0
	links := make([]string, 2)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), ".xml") {
			links[curIdx] = "https://www.sec.gov" + s.AttrOr("href", "")
			curIdx++
		}
	})

	// try the first
	if len(links) < 2 {
		return thirteenF, err
	}
	var infoTable InfoTable
	var primaryDoc Form13_PrimaryDoc
	firstLinkPrimaryDoc, err := parsePrimaryDoc(links[0], userAgent)
	if err != nil {
		// first link is info table
		firstLinkInfoTable, err := parseInfoTable(links[0], userAgent)
		if err != nil {
			return thirteenF, err
		}
		secondLinkPrimaryDoc, err := parsePrimaryDoc(links[1], userAgent)
		if err != nil {
			return thirteenF, err
		}
		infoTable = firstLinkInfoTable
		primaryDoc = secondLinkPrimaryDoc

	} else {
		// first link is primary doc
		primaryDoc = firstLinkPrimaryDoc
		secondLinkInfoTable, err := parseInfoTable(links[1], userAgent)
		if err != nil {
			return thirteenF, err
		}
		infoTable = secondLinkInfoTable
	}

	// populate 13f
	thirteenF.AccessionNumber = accNum
	thirteenF.PeriodOfReport = primaryDoc.Header.FilerInfo.PeriodOfReport
	thirteenF.ReporterName = primaryDoc.FormData.CoverPage.FilingManager.Name
	thirteenF.Cik = primaryDoc.Header.FilerInfo.Filer.Credentials.Cik
	thirteenF.Url = url

	// go over entries
	total := 0.0
	for _, entry := range infoTable.Entries {
		var nHolding iv_models.DB_Form13F_Holding
		cik, err := convertCusipToCik(entry.Cusip)
		if err == nil {
			nHolding.Name = entry.NameOfIssuer
			nHolding.Cik = cik
			nHolding.Shares = float32(entry.SharesOrAmt.SshPrnamt)
			nHolding.NetTotal = float32(entry.Value) * 1000
			total += entry.Value * 1000
			thirteenF.Holdings = append(thirteenF.Holdings, nHolding)
		}
	}

	thirteenF.TotalHeld = float32(total)

	return thirteenF, nil
}

// CUSIP
type tableEntry struct {
	Cik        string
	CusipSix   string
	CusipEight string
}

func convertCusipToCik(cusip string) (string, error) {

	// read the csv file
	file, err := os.Open("data/mapping.csv")

	if err != nil {
		return "", err
	}

	r := csv.NewReader(file)

	// remove last character from cusip
	cusipCheck := cusip[:len(cusip)-1]

	for {

		record, err := r.Read()

		if err != nil {
			return "", err
		}

		var curTable tableEntry
		curTable.Cik = strings.TrimSuffix(record[0], ".0")
		curTable.CusipSix = record[1]
		curTable.CusipEight = record[2]

		for len(curTable.Cik) < 10 {
			curTable.Cik = "0" + curTable.Cik
		}

		if curTable.CusipEight == cusipCheck {
			return curTable.Cik, nil
		}
	}
}

// INFO TABLE
func parseInfoTable(uri string, userAgent string) (InfoTable, error) {
	var doc InfoTable
	// create the http request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return doc, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Host = "www.sec.gov"

	// create the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return doc, err
	}

	// get the data from res body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return doc, err
	}

	// parse the data
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

type InfoTable struct {
	XMLName xml.Name         `xml:"informationTable"`
	Entries []InfoTableEntry `xml:"infoTable"`
}

type InfoTableEntry struct {
	XMLName      xml.Name    `xml:"infoTable"`
	NameOfIssuer string      `xml:"nameOfIssuer"`
	Cusip        string      `xml:"cusip"`
	Value        float64     `xml:"value"`
	SharesOrAmt  SharesOrAmt `xml:"shrsOrPrnAmt"`
}

type SharesOrAmt struct {
	XMLName   xml.Name `xml:"shrsOrPrnAmt"`
	SshPrnamt float64  `xml:"sshPrnamt"`
}

// PRIMARY DOC
func parsePrimaryDoc(uri string, userAgent string) (Form13_PrimaryDoc, error) {
	var doc Form13_PrimaryDoc
	// create the http request
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return doc, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Host = "www.sec.gov"

	// create the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return doc, err
	}

	// get the data from res body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return doc, err
	}

	// parse the data
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		return doc, err
	}

	return doc, nil

}

type Form13_PrimaryDoc struct {
	XMLName  xml.Name   `xml:"edgarSubmission"`
	Header   HeaderData `xml:"headerData"`
	FormData FormData   `xml:"formData"`
}

type FormData struct {
	XMLName   xml.Name  `xml:"formData"`
	CoverPage CoverPage `xml:"coverPage"`
}

type CoverPage struct {
	XMLName       xml.Name      `xml:"coverPage"`
	FilingManager FilingManager `xml:"filingManager"`
}

type FilingManager struct {
	XMLName xml.Name `xml:"filingManager"`
	Name    string   `xml:"name"`
}

type HeaderData struct {
	XMLName   xml.Name  `xml:"headerData"`
	FilerInfo FilerInfo `xml:"filerInfo"`
}

type FilerInfo struct {
	XMLName        xml.Name `xml:"filerInfo"`
	PeriodOfReport string   `xml:"periodOfReport"`
	Filer          Filer    `xml:"filer"`
}

type Filer struct {
	XMLName     xml.Name    `xml:"filer"`
	Credentials Credentials `xml:"credentials"`
}

type Credentials struct {
	XMLName xml.Name `xml:"credentials"`
	Cik     string   `xml:"cik"`
}
