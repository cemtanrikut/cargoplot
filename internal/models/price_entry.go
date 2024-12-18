package models

type PriceEntry struct {
	Company int    `json:"Company"`
	Price   int    `json:"Price"`
	Origin  string `json:"Origin"`
	Date    string `json:"Date"`
}
