package model

type Currency struct {
	Coin string `json:"coin"`
}

type GetCurrencyPriceRequest struct {
	Coin      string `json:"coin"`
	TimeStamp int64  `json:"timestamp"`
}

type GetCurrencyPriceResponse struct {
	Price float64 `json:"price"`
}

type AD_REsponse struct {
	Status string `json:"status"`
	Coin   string `json:"coin"`
}
