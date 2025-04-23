package listing

import (
	"gorm.io/gorm"
	"math"
)

type Pageable struct {
	Page       int         `json:"page,omitempty;query:page"`
	Limit      int         `json:"limit,omitempty;query:limit"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalPages int         `json:"totalPages"`
	TotalRows  int64       `json:"totalRows"`
	Rows       interface{} `json:"rows"`
}

func getPage(p *Pageable) int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func getLimit(p *Pageable) int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func getOffset(p *Pageable) int {
	return (getPage(p) - 1) * getLimit(p)
}

func GetSort(p *Pageable) string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func paginate(value interface{}, paged *Pageable, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db.Model(value).Count(&totalRows)
	paged.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(paged.Limit)))
	paged.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(getOffset(paged)).
			Limit(getLimit(paged)).
			Order(GetSort(paged))
	}
}

func Paginate(value interface{}, pageable Pageable, db *gorm.DB) Pageable {
	db.Scopes(paginate(value, &pageable, db)).Find(&value)
	pageable.Rows = value
	return pageable
}

func PrePaginate(value interface{}, pageable *Pageable, db *gorm.DB) *gorm.DB {
	return db.Scopes(paginate(value, pageable, db))
}

func PostPaginate(value interface{}, pageable *Pageable) Pageable {
	pageable.Rows = value
	return *pageable
}
