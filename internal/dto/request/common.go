package request

type Pagination struct {
	Page     *int `query:"page" json:"page"`
	PageSize *int `query:"page_size" json:"page_size"`
}

type SearchGetRequest struct {
	Pagination
	Search   string   `query:"search"`
	AscField []string `query:"asc_field"`
	DscField []string `query:"dsc_field"`
}
