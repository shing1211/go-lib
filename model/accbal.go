package model

type AccBal struct {
	Coin   string  `json:"coin"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}
