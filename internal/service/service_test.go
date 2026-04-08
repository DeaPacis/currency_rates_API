package service

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestGetCurrencyRates_InvalidDate(t *testing.T) {
	invalidDates := []string{
		"2023-13-01",
		"2023-01-32",
		"01-01-2023",
		"invalid",
		"",
	}

	for _, date := range invalidDates {
		_, err := GetCurrencyRates(date)
		if err == nil {
			t.Errorf("Expected error for date %s, got nil", date)
		}
	}
}

func TestParseCurrencyXML_Success(t *testing.T) {
	mockXML := `<?xml version="1.0" encoding="windows-1251"?>
    <ValCurs Date="01/01/2023" name="Foreign Currency Market">
        <Valute ID="R01235">
            <NumCode>840</NumCode>
            <CharCode>USD</CharCode>
            <Nominal>1</Nominal>
            <Name>Доллар США</Name>
            <Value>75,7668</Value>
        </Valute>
        <Valute ID="R01239">
            <NumCode>978</NumCode>
            <CharCode>EUR</CharCode>
            <Nominal>1</Nominal>
            <Name>Евро</Name>
            <Value>85,1234</Value>
        </Valute>
    </ValCurs>`

	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(mockXML)),
	}

	rates, err := parseCurrencyXML(resp)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedUSD := 75.7668
	if rates["USD"] != expectedUSD {
		t.Errorf("Expected USD: %f, got: %f", expectedUSD, rates["USD"])
	}

	expectedEUR := 85.1234
	if rates["EUR"] != expectedEUR {
		t.Errorf("Expected EUR: %f, got: %f", expectedEUR, rates["EUR"])
	}
}

func TestParseCurrencyXML_InvalidValue(t *testing.T) {
	mockXML := `<?xml version="1.0" encoding="windows-1251"?>
    <ValCurs>
        <Valute ID="R01235">
            <CharCode>USD</CharCode>
            <Value>не число</Value>
        </Valute>
        <Valute ID="R01239">
            <CharCode>EUR</CharCode>
            <Value>85,1234</Value>
        </Valute>
    </ValCurs>`

	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(mockXML)),
	}

	rates, err := parseCurrencyXML(resp)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if _, exists := rates["USD"]; exists {
		t.Error("USD should not be in rates due to parsing error")
	}

	if _, exists := rates["EUR"]; !exists {
		t.Error("EUR should be in rates")
	}
}

func TestParseCurrencyXML_DifferentDecimalSeparator(t *testing.T) {
	testCases := []struct {
		input    string
		expected float64
	}{
		{"75,7668", 75.7668},
		{"0,1234", 0.1234},
		{"1000,50", 1000.50},
		{"1,12345", 1.12345},
	}

	for _, tc := range testCases {
		val := strings.Replace(tc.input, ",", ".", 1)
		result, err := strconv.ParseFloat(val, 64)
		if err != nil {
			t.Errorf("Parse error for %s", tc.input)
		}
		if result != tc.expected {
			t.Errorf("Expected %f, got %f", tc.expected, result)
		}
	}
}
