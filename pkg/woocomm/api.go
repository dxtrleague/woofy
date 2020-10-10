package woocomm

import (
	"context"
)

type (

	// API ...
	API interface {
		Products() Products
		ProductCategories() ProductCategories
		Posts() ProductReviews
	}

	// Products The products API allows you to create, view, update,
	// and delete individual, or a batch, of products.
	Products interface {

		// Delete This API helps you delete a product.
		Delete(c context.Context, id int) (*Product, error)

		// Get This API lets you retrieve and view a specific product by ID.
		Get(c context.Context, id int) (*Product, error)

		// List This API lets you retrieve all products.
		List(c context.Context) (*[]Product, error)
	}

	// ProductCategories ...
	ProductCategories interface {

		// Create This API helps you to create a new product category.
		// minimal category's field are `Name` & `Image`
		// Image must have suffix extention sunch as `.jpg` or `.png`
		Create(c context.Context, prdCtg ProductCategory) (*ProductCategory, error)

		// Delete This API helps you delete a product category.
		// If category is used by any products then the category of them will be set to `Uncategorized`
		Delete(c context.Context, id int) (*ProductCategory, error)

		// Get This API lets you retrieve a product category by ID.
		Get(c context.Context, id int) (*ProductCategory, error)

		// List This API lets you retrieve all product categories.
		List(c context.Context) (*[]ProductCategory, error)
	}

	// ProductReviews The product reviews API allows you to create, view, update,
	// and delete individual, or a batch, of product reviews.
	ProductReviews interface {

		// Create This API helps you to create a new product review.
		Create(c context.Context, prdRvw ProductReview) (*ProductReview, error)

		// DeleteThis API helps you delete a product review.
		Delete(c context.Context, id int) (*ProductReview, error)

		// Get This API lets you retrieve a product review by ID.
		Get(c context.Context, id int) (*ProductReview, error)

		// List This API lets you retrieve all product review.
		List(c context.Context) (*[]ProductReview, error)
	}
)
