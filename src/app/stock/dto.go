package stock

type CreateStockInput struct {
	SauceID  string `json:"sauce_id" binding:"required,uuid"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

type UpdateStockInput struct {
	Quantity int    `json:"quantity" binding:"required,min=1"`
}