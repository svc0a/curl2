package curl2

import (
	"github.com/imroc/req/v3"
	"net/http"
	"net/url"
)

type Service[IN any, OUT any] interface {
	SetUrl(string) Service[IN, OUT]
	SetBody(IN) Service[IN, OUT]
	SetContentType(ContentType) Service[IN, OUT]
	SetHeader(string, string) Service[IN, OUT]
	SetHeaders(map[string]string) Service[IN, OUT]
	SetProxy(string) Service[IN, OUT]
	Get() (*OUT, error)
	Post() (*OUT, error)
}

func Define[IN, OUT any]() Service[IN, OUT] {
	return &impl[IN, OUT]{
		headers: map[string]string{},
	}
}

type impl[IN any, OUT any] struct {
	url         string
	body        *IN
	contentType ContentType
	headers     map[string]string
	proxy       string
}

func (i *impl[IN, OUT]) SetUrl(s string) Service[IN, OUT] {
	i.url = s
	return i
}

func (i *impl[IN, OUT]) SetBody(in IN) Service[IN, OUT] {
	i.body = &in
	return i
}

func (i *impl[IN, OUT]) SetContentType(s ContentType) Service[IN, OUT] {
	i.contentType = s
	return i
}

func (i *impl[IN, OUT]) SetHeader(k string, v string) Service[IN, OUT] {
	i.headers[k] = v
	return i
}

func (i *impl[IN, OUT]) SetHeaders(m map[string]string) Service[IN, OUT] {
	for k, v := range m {
		i.headers[k] = v
	}
	return i
}

func (i *impl[IN, OUT]) SetProxy(s string) Service[IN, OUT] {
	i.proxy = s
	return i
}

func (i *impl[IN, OUT]) Get() (*OUT, error) {
	var out OUT
	cli := req.C().EnableInsecureSkipVerify()
	if i.proxy != "" {
		u, _ := url.Parse(i.proxy)
		proxy := http.ProxyURL(u)
		cli.SetProxy(proxy)
	}
	r := cli.Get()
	if i.contentType != "" {
		r.SetContentType(i.contentType.String())
	}
	if i.headers != nil {
		r.SetHeaders(i.headers)
	}
	if i.body != nil {
		r.SetBody(*i.body)
	}
	err := r.Do().Into(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (i *impl[IN, OUT]) Post() (*OUT, error) {
	var out OUT
	cli := req.C().EnableInsecureSkipVerify()
	if i.proxy != "" {
		u, _ := url.Parse(i.proxy)
		proxy := http.ProxyURL(u)
		cli.SetProxy(proxy)
	}
	r := cli.Post()
	if i.contentType != "" {
		r.SetContentType(i.contentType.String())
	}
	if i.headers != nil {
		r.SetHeaders(i.headers)
	}
	if i.body != nil {
		r.SetBody(*i.body)
	}
	err := r.Do().Into(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
