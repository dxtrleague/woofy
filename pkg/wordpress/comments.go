package wordpress

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dxtrleague/woofy"
	"github.com/rs/zerolog/log"
)

// CommentService ...
type CommentService struct {
	client   *woofy.Client
	endpoint string
}

// NewCommentService ...
func NewCommentService(client *woofy.Client) *CommentService {
	return &CommentService{
		client:   client,
		endpoint: client.Opt.BaseApiURL + "/comments",
	}
}

// Get ...
func (s *CommentService) Get(c context.Context, id int) (*Comment, error) {

	log.Debug().Msgf("Get(%d)", id)
	var pc CommentComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint + fmt.Sprintf("/%v", id)).EndStruct(&pc)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to get, httpStatus: %v", resp.StatusCode)
	}

	comment := toComment(pc)

	return comment, nil

}

// List ...
func (s *CommentService) List(c context.Context) (*[]Comment, error) {

	log.Debug().Msg("List()")
	var pcs []CommentComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&pcs)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	var comments []Comment
	for _, p := range pcs {
		comment := toComment(p)
		comments = append(comments, *comment)
	}

	return &comments, nil

}

func toComment(cc CommentComplex) *Comment {
	comment := &Comment{
		ID:       cc.ID,
		PostID:   cc.PostID,
		ParentID: cc.ParentID,
		Content:  cc.Content.Rendered,
		Date:     cc.Date,
	}
	return comment
}

type (

	// Comment ...
	Comment struct {
		ID       int        `json:"id,omitempty"` // Unique identifier for the term.
		PostID   int        `json:"post"`
		ParentID int        `json:"parent"` // Parent Comment ID
		Content  string     `json:"content"`
		Date     *time.Time `json:"date"`
	}

	// CommentComplex is to hold JSON response from WP-REST API
	CommentComplex struct {
		ID       int        `json:"id,omitempty"`
		PostID   int        `json:"post"`
		ParentID int        `json:"parent"` // Parent Comment ID
		Content  Content    `json:"content"`
		Date     *time.Time `json:"date"`
	}
	// CommentTitle ...
	CommentTitle struct {
		Rendered string `json:"rendered"`
	}
	// CommentContent ...
	CommentContent struct {
		Rendered string `json:"rendered"`
	}
)
