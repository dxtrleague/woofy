package woocomm

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dxtrleague/woofy"
	"github.com/rs/zerolog/log"
)

// ProductCategoryService ...
type ProductCategoryService struct {
	client   *woofy.Client
	endpoint string
}

// NewProductCategoryService ...
func NewProductCategoryService(client *woofy.Client) *ProductCategoryService {
	return &ProductCategoryService{
		client:   client,
		endpoint: fmt.Sprintf("%s/products/categories", client.Opt.BaseApiURL),
	}
}

// Get This API lets you retrieve a product Category by ID.
func (s *ProductCategoryService) Get(c context.Context, id int) (*ProductCategory, error) {

	var pcx ProductCategoryComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint + fmt.Sprintf("/%v", id)).EndStruct(&pcx)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	prdCtg := toProductCategory(pcx)

	return prdCtg, nil

}

// Delete This API helps you delete a product category.
func (s *ProductCategoryService) Delete(c context.Context, id int) (*ProductCategory, error) {

	var pcx ProductCategoryComplex
	resp, _, errs := s.client.Req.Clone().Delete(s.endpoint + fmt.Sprintf("/%v?force=true", id)).EndStruct(&pcx)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to delete, httpStatus: %v", resp.StatusCode)
	}

	prdCtg := toProductCategory(pcx)

	return prdCtg, nil

}

// Create This API helps you to create a new product Category.
func (s *ProductCategoryService) Create(c context.Context, prdCtg ProductCategory) (*ProductCategory, error) {

	log.Debug().Msgf("Create(n: %s, i: %s)", prdCtg.Name, prdCtg.Image)
	if prdCtg.Name == "" {
		return nil, errors.New("name is empty")
	}
	if prdCtg.Image == "" {
		return nil, errors.New("image is empty")
	}

	payload := toProductCategoryComplex(prdCtg)
	var prdCtgCpx ProductCategoryComplex
	resp, _, errs := s.client.Req.Clone().Post(s.endpoint).Type("json").SendStruct(payload).EndStruct(&prdCtgCpx)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("fail to create, httpStatus: %v", resp.StatusCode)
	}

	pCtg := toProductCategory(prdCtgCpx)

	return pCtg, nil

}

// List ...
func (s *ProductCategoryService) List(c context.Context) (*[]ProductCategory, error) {

	log.Debug().Msg("List()")
	var pcxs []ProductCategoryComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&pcxs)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	var prdCtgs []ProductCategory
	for _, p := range pcxs {
		pCtg := toProductCategory(p)
		prdCtgs = append(prdCtgs, *pCtg)
	}

	return &prdCtgs, nil

}

func toProductCategory(pcx ProductCategoryComplex) *ProductCategory {
	comment := &ProductCategory{
		ID:          pcx.ID,
		ParentID:    pcx.ParentID,
		Name:        pcx.Name,
		Description: pcx.Description,
		Image:       pcx.Image.Source,
	}
	return comment
}

func toProductCategoryComplex(pcx ProductCategory) *ProductCategoryComplex {
	comment := &ProductCategoryComplex{
		ID:          pcx.ID,
		ParentID:    pcx.ParentID,
		Name:        pcx.Name,
		Description: pcx.Description,
		Image: Image{
			Source: pcx.Image,
		},
	}
	return comment
}

type (

	// ProductCategory ...
	ProductCategory struct {
		ID          int    `json:"id,omitempty"` // Unique identifier for the resource.
		ParentID    int    `json:"parent"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}

	// ProductCategoryComplex ...
	ProductCategoryComplex struct {
		ID          int    `json:"id,omitempty"` // Unique identifier for the resource.
		ParentID    int    `json:"parent"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
	}
)

// CategoryStatus is the Status of the ProductCategory
// type CategoryStatus string

// CategoryRating ...
// type CategoryRating uint8

// const (
// 	CategoryStatusApproved CategoryStatus = "approved"
// 	CategoryStatusHold                    = "hold"
// 	CategoryStatusSpam                    = "spam"
// 	CategoryStatusUnspam                  = "unspam"
// 	CategoryStatusTrash                   = "trash"
// 	CategoryStatusUntrash                 = "untrash"
// )

// const (
// 	CategoryRatingPoor      CategoryRating = iota + 1 // 1 star rating
// 	CategoryRatingFair                                // 2 stars rating
// 	CategoryRatingAverage                             // 3 stars rating
// 	CategoryRatingGood                                // 4 stars rating
// 	CategoryRatingExcellent                           // 5 stars rating
// )

// ProductCategoryFilter ...
// type ProductCategoryFilter struct {
// 	Context           FilterContext  `json:"context"`            // Scope under which the request is made; determines fields present in response. Options: view and edit. Default is view.
// 	Page              int            `json:"page"`               // Current page of the collection. Default is 1.
// 	PerPage           int            `json:"per_page"`           // Maximum number of items to be returned in result set. Default is 10.
// 	Search            string         `json:"search"`             // Limit results to those matching a string.
// 	After             string         `json:"after"`              // Limit response to resources published after a given ISO8601 compliant date.
// 	Before            string         `json:"before"`             // Limit response to resources published before a given ISO8601 compliant date.
// 	Exclude           []int          `json:"exclude"`            // Ensure result set excludes specific IDs.
// 	Include           []int          `json:"include"`            // Limit result set to specific ids.
// 	Offset            int            `json:"offset"`             // Offset the result set by a specific number of items.
// 	Order             Order          `json:"order"`              // Order sort attribute ascending or descending. Options: asc and desc. Default is desc.
// 	OrderBy           OrderBy        `json:"orderby"`            // Sort collection by resource attribute. Options: date, date_gmt, id, slug, include and product. Default is date_gmt.
// 	Categoryer        []int          `json:"Categoryer"`         // Limit result set to Categorys assigned to specific user IDs.
// 	CategoryerExclude []int          `json:"Categoryer_exclude"` // Ensure result set excludes Categorys assigned to specific user IDs.
// 	CategoryerEmail   []string       `json:"Categoryer_email"`   // Limit result set to that from a specific author email.
// 	Product           []int          `json:"product"`            // Limit result set to Categorys assigned to specific product IDs.
// 	Status            CategoryStatus `json:"status"`             // Status of the Category. Options: approved, hold, spam, unspam, trash and untrash. Defaults to approved.
// }

// NewProductCategoryFilter ...
// func NewProductCategoryFilter() *ProductCategoryFilter {
// 	return &ProductCategoryFilter{
// 		Context: FilterContextView,
// 		Page:    1,
// 		PerPage: 10,
// 		Order:   OrderDSC,
// 		OrderBy: OrderByDateGMT,
// 		Status:  CategoryStatusApproved,
// 	}
// }
