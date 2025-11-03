package api

type CreateProductDto struct {
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}
