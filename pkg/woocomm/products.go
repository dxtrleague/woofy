package woocomm

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dxtrleague/woofy"
)

// ProductService API implementation of Products
type ProductService struct {
	client   *woofy.Client
	endpoint string
}

// NewProductService ...
func NewProductService(client *woofy.Client) *ProductService {
	return &ProductService{
		client:   client,
		endpoint: fmt.Sprintf("%s/products", client.Opt.BaseApiURL),
	}
}

// Delete This API helps you delete a product.
func (s *ProductService) Delete(c context.Context, id int) (*Product, error) {

	var p Product
	resp, _, errs := s.client.Req.Clone().Delete(fmt.Sprintf("%s/%v?force=true", s.endpoint, id)).EndStruct(&p)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to delete, httpStatus: %v", resp.StatusCode)
	}

	return &p, nil
}

// Get This API lets you retrieve a product by ID.
func (s *ProductService) Get(c context.Context, id int) (*Product, error) {

	var p Product
	resp, _, errs := s.client.Req.Clone().Get(fmt.Sprintf("%s/%v", s.endpoint, id)).EndStruct(&p)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to get, httpStatus: %v", resp.StatusCode)
	}

	return &p, nil
}

// List This API lets you retrieve all products.
func (s *ProductService) List(c context.Context) (*[]Product, error) {

	var ps []Product
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&ps)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	return &ps, nil
}

// Product ...
type Product struct {
	ID          int        `json:"id,omitempty"`
	SKU         string     `json:"sku"`
	Name        string     `json:"name"`
	Price       float64    `json:"price,string"`
	Description string     `json:"description"`
	DateCreated *time.Time `json:"date_created"`
}
