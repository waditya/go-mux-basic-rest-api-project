package store

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}
