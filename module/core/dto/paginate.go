package dto

import "golang-server/pkg/constants"

type Paginate struct {
	Limit  *int `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Offset *int `form:"offset" binding:"omitempty,gte=0"`
}

type PaginateResponse struct {
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Total  *int64 `json:"total,omitempty"`
}

func (p Paginate) InfoPaginate() (int, int) {
	limit := constants.DefaultLimit
	offset := constants.DefaultOffset

	if p.Limit != nil {
		if *p.Limit > constants.MaxLimit {
			limit = constants.MaxLimit
		} else {
			limit = *p.Limit
		}
	}
	if p.Offset != nil {
		offset = *p.Offset
	}
	return limit, offset
}
