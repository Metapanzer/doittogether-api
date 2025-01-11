package model

type WebResponse[T any] struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    T               `json:"data"`
	Paging  *PagingMetadata `json:"paging,omitempty"`
}

type PagingMetadata struct {
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	Size      int `json:"size"`
}
