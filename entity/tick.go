package entity

type Tick struct {
	Id        int     `json:"-"`
	TimeStamp int64   `json:"timestamp"`
	Symbol    string  `json:"symbol"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
}
