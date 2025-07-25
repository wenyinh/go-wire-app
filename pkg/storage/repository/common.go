package repository

import (
	"strings"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

const AscOrder = "ASC"

type QueryModel interface {
	GetFieldByName(fieldName string) (field.OrderExpr, bool)
}

type PaginationParams struct {
	SortBy     string `json:"SortBy,omitempty" form:"SortBy" default:"gmt_create" example:"gmt_create"`                     // SortBy is the field which used to sort.
	Order      string `json:"Order,omitempty" binding:"oneof=Desc Asc DESC ASC" form:"Order" default:"DESC" example:"DESC"` // Order desc or asc
	PageNumber int    `json:"PageNumber,omitempty" binding:"numeric" form:"PageNumber" default:"1" example:"1"`             // PageNumber is the current page number when query.
	PageSize   int    `json:"PageSize,omitempty" binding:"numeric" form:"PageSize" default:"20" example:"20"`               // PageSize is the number of the page when query.
}

func WithPaginationQuery(paginationParams *PaginationParams, queryModel QueryModel) func(tx gen.Dao) gen.Dao {
	pageSize := paginationParams.PageSize
	pageNumber := paginationParams.PageNumber
	order := paginationParams.Order
	sortBy := paginationParams.SortBy

	return func(tx gen.Dao) gen.Dao {
		if pageSize > 0 && pageNumber > 0 {
			tx = tx.Offset((pageNumber - 1) * pageSize).Limit(pageSize)
		}

		if order != "" && sortBy != "" {
			sortField, ok := queryModel.GetFieldByName(sortBy)
			if ok {
				if strings.ToUpper(order) == AscOrder {
					tx = tx.Order(sortField)
				} else {
					tx = tx.Order(sortField.Desc())
				}
			}
		}
		return tx
	}
}

func DefaultPaginationParams() *PaginationParams {
	return &PaginationParams{
		PageNumber: 1,
		PageSize:   20,
		SortBy:     "gmt_create",
		Order:      "DESC",
	}
}

type SortParams struct {
	SortBy string `json:"SortBy,omitempty" form:"SortBy" default:"gmt_create" example:"gmt_create"`                     // SortBy is the field which used to sort.
	Order  string `json:"Order,omitempty" binding:"oneof=Desc Asc DESC ASC" form:"Order" default:"DESC" example:"DESC"` // Order desc or asc
}

func WithSortQuery(sortParams *SortParams, queryModel QueryModel) func(tx gen.Dao) gen.Dao {
	paginationParams := &PaginationParams{
		PageSize:   0,
		PageNumber: 0,
		Order:      sortParams.Order,
		SortBy:     sortParams.SortBy,
	}
	return WithPaginationQuery(paginationParams, queryModel)
}
