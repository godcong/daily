package daily

import (
	"net/http"
)

/*Client Client */
type Client struct {
	client *http.Client
}

func (c *Client) HTTPGet(url string, header, query map[string]interface{}) map[string]interface{} {
	return nil
}

func (c *Client) HTTPPost(url string, header, query, param map[string]interface{}) map[string]interface{} {
	return nil
}
