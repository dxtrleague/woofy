package woocomm

import "github.com/dxtrleague/woofy"

type woocommerceAPI struct {
	c         *woofy.Client
	prdSvc    *ProductService
	prdCtgSvc *ProductCategoryService
	prdRvw    *ProductReviewService
}

// BuildAPI is for building WooCommerce Rest API client
// https://woocommerce.github.io/woocommerce-rest-api-docs/
func BuildAPI(client *woofy.Client) *woocommerceAPI {
	return &woocommerceAPI{
		c:         client,
		prdSvc:    NewProductService(client),
		prdCtgSvc: NewProductCategoryService(client),
		prdRvw:    NewProductReviewService(client),
	}
}

func (w *woocommerceAPI) Products() Products {
	return w.prdSvc
}

func (w *woocommerceAPI) ProductCategories() ProductCategories {
	return w.prdCtgSvc
}

func (w *woocommerceAPI) ProductReviews() ProductReviews {
	return NewProductReviewService(w.c)
}
