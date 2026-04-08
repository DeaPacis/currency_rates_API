package service

import (
	"app/internal/model"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
)

func GetCurrencyRates(date string) (map[string]float64, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	resp, err := fetchCurrencyXML(t)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("error closing response body:", err)
		}
	}()

	return parseCurrencyXML(resp)
}

func fetchCurrencyXML(t time.Time) (*http.Response, error) {
	url := fmt.Sprintf(
		"https://www.cbr.ru/scripts/XML_daily.asp?date_req=%s",
		t.Format("02/01/2006"),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseCurrencyXML(resp *http.Response) (map[string]float64, error) {
	var data model.ValCurs

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	rates := map[string]float64{}

	for _, v := range data.Valutes {
		val := strings.Replace(v.Value, ",", ".", 1)

		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Printf("Error parsing currency \"%s\" value", v.CharCode)
			continue
		}

		rates[v.CharCode] = floatVal
	}

	return rates, nil
}
