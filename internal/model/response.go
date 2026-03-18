package model

type InfoResponse struct {
	Version string `json:"version"`
	Service string `json:"service"`
	Author  string `json:"author"`
}

type CurrencyResponse struct {
	Data    map[string]float64 `json:"data"`
	Service string             `json:"service"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
