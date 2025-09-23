package resource

// PaginationQueryParam is a pagination query param
type PaginationQueryParam struct {
	Query    string `form:"query" json:"query"`
	Sort     string `form:"sort" json:"sort"`
	Order    string `form:"order" json:"order"`
	Limit    int    `form:"limit,default=10" json:"limit"`
	Offset   int    `form:"offset,default=0" json:"offset"`
	GameCode string `form:"game_code" json:"game_code"`
	Slug     string `form:"slug" json:"slug"`
	Lang     string `form:"lang" json:"lang"`
	Status   string `form:"status" json:"status"`
}

// Meta is a meta response
type Meta struct {
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}
