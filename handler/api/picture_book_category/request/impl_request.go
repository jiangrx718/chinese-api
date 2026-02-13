package request

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type" binding:"required"`
}

type CategoryUpdateRequest struct {
	CategoryId   string `json:"category_id" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Position     int    `json:"position"`
	CategoryType int    `json:"category_type" binding:"required"`
}

type CategoryDeleteRequest struct {
	CategoryId string `json:"category_id" binding:"required"`
}

type CategoryStatusRequest struct {
	CategoryId string `json:"category_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type CategoryListRequest struct {
	Offset       int64  `form:"offset"`
	Limit        int64  `form:"limit"`
	CategoryName string `form:"category_name"`
	CategoryType int    `form:"category_type"`
}
