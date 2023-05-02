package product

type ProductCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
