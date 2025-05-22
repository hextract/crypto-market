package dto

type JsonExecutionResult struct {
	Order  *JsonOrder `json:"order"`
	Price  float64    `json:"price"`
	Volume float64    `json:"volume"`
}
