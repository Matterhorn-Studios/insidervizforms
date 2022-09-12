package insidervizforms

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"

	"github.com/Matterhorn-Studios/insidervizforms/iv_models"
)

func XmlForm4ToRawForm4(path string, accNum string, url string) (iv_models.RawForm4, error) {
	var form4 iv_models.RawForm4
	xmlFile, err := os.Open(path)
	if err != nil {
		return form4, err
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &form4)

	form4.AccessionNumber = accNum
	form4.Url = url

	return form4, nil
}

func XmlUrlForm4ToRawForm4(url string, accNum string, userAgent string) (iv_models.RawForm4, error) {
	var form4 iv_models.RawForm4

	// create the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return form4, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Host = "www.sec.gov"

	// create the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return form4, err
	}

	// get data from res body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return form4, err
	}

	// parse the xml
	xml.Unmarshal(data, &form4)

	form4.AccessionNumber = accNum
	form4.Url = url

	return form4, nil
}
