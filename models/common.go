package models

type PageSize struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type PaginateReq struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}
