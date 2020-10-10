package woofy

import (
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// Client ...
type Client struct {
	Req *gorequest.SuperAgent
	Opt *Options
}

// Options ...
type Options struct {
	BaseApiURL string
	Token      string
}

// NewClient ...
func NewClient(options *Options) *Client {
	req := newHttpClient()
	req.Set("Authorization", fmt.Sprintf("Bearer %s", options.Token))
	return &Client{
		Req: req,
		Opt: options,
	}
}

func newHttpClient() *gorequest.SuperAgent {
	client := gorequest.New()
	client.Client = &http.Client{Jar: nil}
	client.Transport = &http.Transport{
		DisableKeepAlives: true,
	}
	return client
}
