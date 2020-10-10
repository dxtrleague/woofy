package woocomm

import "time"

// Image ...
type Image struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Source       string     `json:"src"`
	AltText      string     `json:"alt"`
	DateCreated  *time.Time `json:"date_created"`
	DateModified *time.Time `json:"date_modified"`
}

// FilterContext is scope under which the request is made; determines fields present in response.
type FilterContext string

const (
	FilterContextView FilterContext = "view"
	FilterContextEdit               = "edit"
)

// Order sort attribute ascending or descending.
type Order string

const (
	OrderASC Order = "asc"
	OrderDSC       = "desc"
)

// OrderBy Sort collection by object attribute.
type OrderBy string

const (
	OrderByDate    OrderBy = "date"
	OrderByDateGMT         = "date_gmt"
	OrderByID              = "id"
	OrderByInclude         = "include"
	OrderByTitle           = "title"
	OrderBySlug            = "slug"
	OrderByProduct         = "product"
)
