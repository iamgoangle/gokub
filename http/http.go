package http

import (
	"github.com/hashicorp/go-retryablehttp"
	"log"
	"net/http"
)

type Configs struct {
	HTTPClient *http.Client
}

type Client struct {
	http *retryablehttp.Client
}

// New instant new http client
func New(cfg *Configs) (*Client, error) {
	h := newRetryableHTTPClient()
	c := &Client{
		http: h,
	}

	log.Println("[New]: new http client successfully")

	return c, nil
}

// Get is a http wrapper to GET request
func (c *Client) Get(url string) (*http.Response, error) {
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post is a http wrapper to POST request
func (c *Client) Post(url string) (*http.Response, error) {
	return nil, nil
}

// Put is a http wrapper to Put request
func (c *Client) Put(url string) (*http.Response, error) {
	return nil, nil
}

// Patch is a http wrapper to PATCH request
func (c *Client) Patch(url string) (*http.Response, error) {
	return nil, nil
}

// Delete is a http wrapper to DELETE request
func (c *Client) Delete(url string) (*http.Response, error) {
	return nil, nil
}

// Do is a http wrapper to golang http.do request
func (c *Client) Do(url string) (*http.Response, error) {
	return nil, nil
}

// PostForm is a http wrapper to Post Form request
func (c *Client) PostForm(url string) (*http.Response, error) {
	return nil, nil
}

// Head is a http wrapper to HEAD request
func (c *Client) Head(url string) (*http.Response, error) {
	return nil, nil
}

// newRetryableHTTPClient is a vendor http retryable
func newRetryableHTTPClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()

	return retryClient
}
