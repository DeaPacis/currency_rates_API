package model

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}
