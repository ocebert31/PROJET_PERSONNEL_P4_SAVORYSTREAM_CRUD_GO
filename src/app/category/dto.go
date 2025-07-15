package category

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}
