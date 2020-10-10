package wordpress

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dxtrleague/woofy"
	"github.com/rs/zerolog/log"
)

// PostService ...
type PostService struct {
	client   *woofy.Client
	endpoint string
}

// NewPostService ...
func NewPostService(client *woofy.Client) *PostService {
	return &PostService{
		client:   client,
		endpoint: client.Opt.BaseApiURL + "/posts",
	}
}

// Get ...
func (s *PostService) Get(c context.Context, id int) (*Post, error) {

	log.Debug().Msgf("Get(%d)", id)
	var pc PostComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint + fmt.Sprintf("/%v", id)).EndStruct(&pc)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to get, httpStatus: %v", resp.StatusCode)
	}

	post := toPost(pc)

	return post, nil
}

// List ...
func (s *PostService) List(c context.Context) (*[]Post, error) {

	log.Debug().Msg("List()")
	var pcs []PostComplex
	resp, _, errs := s.client.Req.Clone().Get(s.endpoint).EndStruct(&pcs)
	if errs != nil && len(errs) > 0 {
		return nil, errs[len(errs)-1]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to list, httpStatus: %v", resp.StatusCode)
	}

	var posts []Post
	for _, p := range pcs {
		post := toPost(p)
		posts = append(posts, *post)
	}

	return &posts, nil

}

func toPost(pc PostComplex) *Post {
	post := &Post{
		ID:      pc.ID,
		Title:   pc.Title.Rendered,
		Content: pc.Content.Rendered,
	}
	return post
}

type (

	// Post ...
	Post struct {
		ID      int    `json:"id,omitempty"` // Unique identifier for the term.
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// PostComplex is to hold JSON response from WP-REST API
	PostComplex struct {
		ID      int       `json:"id,omitempty"`
		Title   PostTitle `json:"title"`
		Content Content   `json:"content"`
	}
	// PostTitle ...
	PostTitle struct {
		Rendered string `json:"rendered"`
	}
	// PostContent ...
	PostContent struct {
		Rendered string `json:"rendered"`
	}
)
