package daily

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

/*Client Client */
type Client struct {
	client *http.Client
}

func DefaultClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

func (c *Client) HTTPGet(url string, header, query map[string]interface{}) map[string]interface{} {
	request, err := http.NewRequest("get", url, strings.NewReader(""))
	if err != nil {
		return nil
	}
	ctx := context.Background()
	response, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	rData, err := ioutil.ReadAll(io.LimitReader(response.Body, 1<<20))
	result := make(map[string]interface{})
	err = json.Unmarshal(rData, &result)
	if err != nil {
		return nil
	}
	return result
}

func (c *Client) HTTPPost(url string, header, query, param map[string]interface{}) map[string]interface{} {
	return nil
}
