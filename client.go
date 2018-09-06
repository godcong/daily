package daily

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func HTTPGet(url string, header http.Header, query url.Values) interface{} {
	return DefaultClient().HTTPGet(url, header, query)
}

func (c *Client) HTTPGet(uri string, header http.Header, query url.Values) interface{} {

	if query != nil {
		uri = uri + "?" + query.Encode()
	}

	request, err := http.NewRequest("GET", uri, nil)
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
	//result := make(map[string]interface{})
	var result interface{}
	err = json.Unmarshal(rData, &result)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return result
}

func (c *Client) HTTPPost(url string, header, query, param map[string]interface{}) map[string]interface{} {
	return nil
}
