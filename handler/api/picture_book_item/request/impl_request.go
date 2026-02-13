package request

type ItemCreateRequest struct {
	BookId   string `json:"book_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Pic      string `json:"pic" binding:"required"`
	BPic     string `json:"b_pic"`
	Audio    string `json:"audio" binding:"required"`
	Content  string `json:"content"`
	Status   string `json:"status" binding:"required"`
	Position int    `json:"position"`
}

type ItemUpdateRequest struct {
	Id       int    `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Pic      string `json:"pic" binding:"required"`
	BPic     string `json:"b_pic"`
	Audio    string `json:"audio" binding:"required"`
	Content  string `json:"content"`
	Status   string `json:"status" binding:"required"`
	Position int    `json:"position"`
}

type ItemDeleteRequest struct {
	Id int `json:"id" binding:"required"`
}

type ItemStatusRequest struct {
	Id     int    `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type ItemListUri struct {
	BookId string `uri:"book_id" binding:"required"`
}

type ItemListRequest struct {
	Offset int64 `form:"offset"`
	Limit  int64 `form:"limit"`
}
