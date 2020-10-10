package woocomm

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dxtrleague/woofy"
)

// ProductReviewService ...
type ProductReviewService struct {
	client   *woofy.Client
	endpoint string
}

// NewProductReviewService ...
func NewProductReviewService(client *woofy.Client) *ProductReviewService {
	return &ProductReviewService{
		client:   client,
		endpoint: fmt.Sprintf("%s/products/reviews", client.Opt.BaseApiURL),
	}
}

// Delete This API helps you delete a product review.
func (s *ProductReviewService) Delete(c context.Context, id int) (*ProductReview, error) {

	var dpr ProductReviewDeleted
	var pr ProductReview
	resp, _, errs := s.client.Req.Clone().Delete(fmt.Sprintf("%s/%v?force=true", s.endpoint, id)).EndStruct(&dpr)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to delete, httpStatus: %v", resp.StatusCode)
	}
	if dpr.Deleted {
		pr = dpr.Previous
	}

	return &pr, nil
}

// Get This API lets you retrieve a product review by ID.
func (s *ProductReviewService) Get(c context.Context, id int) (*ProductReview, error) {

	var pr ProductReview
	resp, _, errs := s.client.Req.Clone().Get(fmt.Sprintf("%s/%v", s.endpoint, id)).EndStruct(&pr)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to get, httpStatus: %v", resp.StatusCode)
	}

	return &pr, nil

}

// Create This API helps you to create a new product review.
func (s *ProductReviewService) Create(c context.Context, prdRvw ProductReview) (*ProductReview, error) {

	if prdRvw.ProductID <= 0 {
		return nil, errors.New("product id invalid")
	}
	if prdRvw.Review == "" {
		return nil, errors.New("review is empty")
	}
	if prdRvw.Reviewer == "" {
		return nil, errors.New("reviewer is empty")
	}
	if prdRvw.ReviewerEmail == "" {
		return nil, errors.New("reviewer email is empty")
	}
	if !validRating(prdRvw.Rating) {
		return nil, errors.New("rating invalid")
	}

	var pr ProductReview
	resp, _, errs := s.client.Req.Clone().Post(s.endpoint).Type("json").SendStruct(prdRvw).EndStruct(&pr)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("fail to create, httpStatus: %v", resp.StatusCode)
	}

	return &pr, nil
}

// List This API lets you retrieve all product review.
func (s *ProductReviewService) List(c context.Context) (*[]ProductReview, error) {

	var prs []ProductReview
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&prs)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	return &prs, nil
}

// ProductReview ...
type ProductReview struct {
	ID            int          `json:"id,omitempty"`   // Unique identifier for the resource.
	ProductID     int          `json:"product_id"`     // Unique identifier for the product that the review belongs to.
	Review        string       `json:"review"`         // The content of the review.
	Reviewer      string       `json:"reviewer"`       // Reviewer name.
	ReviewerEmail string       `json:"reviewer_email"` // Reviewer email.
	Rating        ReviewRating `json:"rating"`         // Review rating (0 to 5).
	DateCreated   string       `json:"date_created"`   // The date the review was created, in the site's timezone.
}

// ReviewRating Review rating (0 to 5).
type ReviewRating uint8

const (
	// ReviewRatingPoor 1 star rating
	ReviewRatingPoor ReviewRating = iota + 1
	// ReviewRatingFair 2 stars rating
	ReviewRatingFair
	// ReviewRatingAverage 3 stars rating
	ReviewRatingAverage
	// ReviewRatingGood 4 stars rating
	ReviewRatingGood
	// ReviewRatingExcellent 5 stars rating
	ReviewRatingExcellent
)

var ratings = []ReviewRating{
	ReviewRatingPoor,
	ReviewRatingFair,
	ReviewRatingAverage,
	ReviewRatingGood,
	ReviewRatingExcellent,
}

func validRating(rating ReviewRating) bool {
	for _, rate := range ratings {
		if rate == rating {
			return true
		}
	}
	return false
}

// ProductReviewDeleted ...
type ProductReviewDeleted struct {
	Deleted  bool          `json:"deleted"`
	Previous ProductReview `json:"previous"`
}

// ReviewStatus is the Status of the ProductReview
type ReviewStatus string

const (
	ReviewStatusApproved ReviewStatus = "approved"
	ReviewStatusHold                  = "hold"
	ReviewStatusSpam                  = "spam"
	ReviewStatusUnspam                = "unspam"
	ReviewStatusTrash                 = "trash"
	ReviewStatusUntrash               = "untrash"
)

// ProductReviewFilter ...
type ProductReviewFilter struct {
	Context         FilterContext `json:"context"`          // Scope under which the request is made; determines fields present in response. Options: view and edit. Default is view.
	Page            int           `json:"page"`             // Current page of the collection. Default is 1.
	PerPage         int           `json:"per_page"`         // Maximum number of items to be returned in result set. Default is 10.
	Search          string        `json:"search"`           // Limit results to those matching a string.
	After           string        `json:"after"`            // Limit response to resources published after a given ISO8601 compliant date.
	Before          string        `json:"before"`           // Limit response to resources published before a given ISO8601 compliant date.
	Exclude         []int         `json:"exclude"`          // Ensure result set excludes specific IDs.
	Include         []int         `json:"include"`          // Limit result set to specific ids.
	Offset          int           `json:"offset"`           // Offset the result set by a specific number of items.
	Order           Order         `json:"order"`            // Order sort attribute ascending or descending. Options: asc and desc. Default is desc.
	OrderBy         OrderBy       `json:"orderby"`          // Sort collection by resource attribute. Options: date, date_gmt, id, slug, include and product. Default is date_gmt.
	Reviewer        []int         `json:"reviewer"`         // Limit result set to reviews assigned to specific user IDs.
	ReviewerExclude []int         `json:"reviewer_exclude"` // Ensure result set excludes reviews assigned to specific user IDs.
	ReviewerEmail   []string      `json:"reviewer_email"`   // Limit result set to that from a specific author email.
	Product         []int         `json:"product"`          // Limit result set to reviews assigned to specific product IDs.
	Status          ReviewStatus  `json:"status"`           // Status of the review. Options: approved, hold, spam, unspam, trash and untrash. Defaults to approved.
}

// NewProductReviewFilter ...
func NewProductReviewFilter() *ProductReviewFilter {
	return &ProductReviewFilter{
		Context: FilterContextView,
		Page:    1,
		PerPage: 10,
		Order:   OrderDSC,
		OrderBy: OrderByDateGMT,
		Status:  ReviewStatusApproved,
	}
}
