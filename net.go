package daily

import (
	"net/http"
	"net/url"
)

func defaultHeader() http.Header {
	header := make(http.Header)
	header.Add("X-RateLimit-Limit", "60")
	header.Add("X-RateLimit-Remaining", "56")
	header.Add("X-RateLimit-Reset", "1372700873")
	return header
}

func defaultQuery() url.Values {
	values := make(url.Values)
	values.Add("per_page", "100")
	return values
}
