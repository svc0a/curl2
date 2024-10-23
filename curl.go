package curl2

import (
	"github.com/imroc/req/v3"
	"net/http"
)

func Get[IN, OUT any](url string, body *IN, headers map[string]string) (*OUT, error) {
	return call[IN, OUT](url, body, headers, http.MethodGet, nil)
}

func Post[IN, OUT any](url string, body *IN, headers map[string]string) (*OUT, error) {
	return call[IN, OUT](url, body, headers, http.MethodPost, nil)
}

func PostJson[IN, OUT any](url string, body *IN, headers map[string]string) (*OUT, error) {
	t := "application/json"
	return call[IN, OUT](url, body, headers, http.MethodPost, &t)
}

func PostXWwwFormUrlencoded[IN, OUT any](url string, body *IN, headers map[string]string) (*OUT, error) {
	t := "application/x-www-form-urlencoded"
	return call[IN, OUT](url, body, headers, http.MethodPost, &t)
}

func call[IN, OUT any](url string, body *IN, headers map[string]string, method string, ContentType *string) (*OUT, error) {
	var out OUT
	var r *req.Request
	if method == http.MethodGet {
		r = req.C().EnableInsecureSkipVerify().Get(url)
	} else {
		r = req.C().EnableInsecureSkipVerify().Post(url)
	}
	if ContentType != nil {
		r.SetContentType(*ContentType)
	}
	if headers != nil {
		r.SetHeaders(headers)
	}
	if body != nil {
		r.SetBody(*body)
	}
	err := r.Do().Into(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
