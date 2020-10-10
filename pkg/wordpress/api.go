package wordpress

import (
	"context"
)

type (

	// API ...
	API interface {
		Categories() Categories
		Comments() Comments
		Posts() Posts
	}

	// Categories ...
	Categories interface {
		List(c context.Context) (*[]Category, error)
	}

	// Comments ...
	Comments interface {
		Get(c context.Context, id int) (*Comment, error) // Get comment by ID
		List(c context.Context) (*[]Comment, error)
	}

	// Posts ...
	Posts interface {
		Get(c context.Context, id int) (*Post, error)
		List(c context.Context) (*[]Post, error)
	}
)
