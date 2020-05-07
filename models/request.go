package models

type UserDTO struct {
	Name string `json:"name"binding:"required"`
	Age  int    `json:"age"binding:"required"`
}

type GetUpVoteRequest struct {
	PageId string `json:"pageid"binding:"required"`
	SiteId string `json:"siteid"binding:"required"`
}

type GetUpVoteResponse struct {
	Votes int `json:"votes"`
}
