package daily

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
	request, err := http.NewRequest("GET", url, nil)
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
	log.Println(string(rData))
	err = json.Unmarshal(rData, &result)
	if err != nil {
		return nil
	}
	return result
}

func (c *Client) HTTPPost(url string, header, query, param map[string]interface{}) map[string]interface{} {
	return nil
}
