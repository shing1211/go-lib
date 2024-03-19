package model

type Holding struct {
	InstruID        string  `bson:"instruid" json:"instruid,omitempty"`
	PLCode          string  `bson:"plcode" json:"plcode,omitempty"`
	StockCode       string  `bson:"stockcode" json:"stockcode,omitempty"`
	InstruShortName string  `bson:"instrushortname" json:"instrushortname,omitempty"`
	TotalQty        float64 `bson:"totalqty" json:"totalqty,omitempty"`
}
