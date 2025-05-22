package dto

type JsonAccountBalance struct {
	BtcFree    float64 `json:"btc_free"`
	BtcLocked  float64 `json:"btc_locked"`
	EthFree    float64 `json:"eth_free"`
	EthLocked  float64 `json:"eth_locked"`
	UsdtFree   float64 `json:"usdt_free"`
	UsdtLocked float64 `json:"usdt_locked"`
}
