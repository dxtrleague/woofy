package wordpress

import "github.com/dxtrleague/woofy"

type wordpressAPI struct {
	c       *woofy.Client
	ctgSvc  *CategoryService
	cmntSvc *CommentService
	postSvc *PostService
}

// BuildAPI is for building WordPress Rest API client
// https://developer.wordpress.org/rest-api/
func BuildAPI(client *woofy.Client) *wordpressAPI {
	return &wordpressAPI{
		c:       client,
		ctgSvc:  NewCategoryService(client),
		cmntSvc: NewCommentService(client),
		postSvc: NewPostService(client),
	}
}

func (w *wordpressAPI) Categories() Categories {
	return w.ctgSvc
}

func (w *wordpressAPI) Comments() Comments {
	return w.cmntSvc
}

func (w *wordpressAPI) Posts() Posts {
	return w.postSvc
}
