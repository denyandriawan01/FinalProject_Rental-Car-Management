package models

type Paginate struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}