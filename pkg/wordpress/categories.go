package wordpress

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dxtrleague/woofy"
	"github.com/rs/zerolog/log"
)

// CategoryService ...
type CategoryService struct {
	client   *woofy.Client
	endpoint string
}

// NewCategoryService ...
func NewCategoryService(client *woofy.Client) *CategoryService {
	return &CategoryService{
		client:   client,
		endpoint: client.Opt.BaseApiURL + "/categories",
	}
}

// List ...
func (s *CategoryService) List(c context.Context) (*[]Category, error) {

	log.Debug().Msg("List()")
	var categories []Category
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&categories)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	return &categories, nil

}

// Category ...
type Category struct {
	ID          int    `json:"id,omitempty"` // Unique identifier for the term.
	Parent      int    `json:"parent"`       // The parent term ID.
	Name        string `json:"name"`         // HTML title for the term. Required: 1
	Slug        string `json:"slug"`         // An alphanumeric identifier for the term unique to its type.
	Description string `json:"description"`  // HTML description of the term.
}
