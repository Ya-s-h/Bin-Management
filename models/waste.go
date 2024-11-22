package model

type Waste struct {
	ID          uint   `json:"waste_id"`
	CollectedAt string `json:"timestamp"`
	BinID       int    `json:"bin_id"`
	Weight      int    `json:"weight_in_kgs"`
	Earnings    int    `json:"earnings"`
}
