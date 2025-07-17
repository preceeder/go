package httpRequest

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type IPv4FallbackClient struct {
	ipv4Client    *resty.Client
	defaultClient *resty.Client
	tryTimeout    time.Duration
}
type R struct {
	ipv4Client    *resty.Client
	defaultClient *resty.Client
	tryTimeout    time.Duration
	Context       context.Context
	Body          any
	Headers       map[string]string
	QueryParam    map[string]string
	Result        any
	FormData      map[string]string
}

// 构造函数
func NewIPv4FallbackClient(timeout time.Duration) *IPv4FallbackClient {
	ipv4Transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: 30 * time.Second,
			Resolver: &net.Resolver{
				PreferGo: true,
				Dial: func(ctx context.Context, _, address string) (net.Conn, error) {
					return net.Dial("tcp4", address)
				},
			},
		}).DialContext,
	}

	ipv4Client := resty.New().SetHeader("Content-Type", "application/json").SetTransport(ipv4Transport)
	ipv4Client.SetTimeout(timeout)

	defaultClient := resty.New().SetHeader("Content-Type", "application/json").SetTimeout(timeout)

	return &IPv4FallbackClient{
		ipv4Client:    ipv4Client,
		defaultClient: defaultClient,
		tryTimeout:    timeout,
	}
}

func (c *IPv4FallbackClient) R() *R {
	return &R{
		ipv4Client:    c.ipv4Client,
		defaultClient: c.defaultClient,
		tryTimeout:    c.tryTimeout,
	}
}

func (c *R) SetBody(body any) *R {
	c.Body = body
	return c
}

func (c *R) SetHeaders(headers map[string]string) *R {
	c.Headers = headers
	return c
}
func (c *R) SetQueryParams(query map[string]string) *R {
	c.QueryParam = query
	return c
}
func (c *R) SetFormData(form map[string]string) *R {
	c.FormData = form
	return c
}

func (c *R) SetResult(result any) *R {
	c.Result = result
	return c
}

func (c *R) SetContext(context context.Context) *R {
	c.Context = context
	return c
}

func (c *R) Get(url string) (*resty.Response, error) {
	return c.do("GET", url)
}

func (c *R) Post(url string) (*resty.Response, error) {
	return c.do("POST", url)
}

func (c *R) Put(url string) (*resty.Response, error) {
	return c.do("PUT", url)
}

func (c *R) Delete(url string) (*resty.Response, error) {
	return c.do("DELETE", url)
}

func (c *R) do(method, url string) (*resty.Response, error) {
	resp, err := c.sendRequest(c.ipv4Client.R(), method, url)
	if err == nil || errors.Is(err, context.DeadlineExceeded) {
		return resp, err
	}
	slog.ErrorContext(c.Context, "ipv4 请求失败", "method", method, "url", url)
	return c.sendRequest(c.defaultClient.R(), method, url)
}
func (c *R) sendRequest(req *resty.Request, method string, url string) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	if c.Body != nil {
		req.SetBody(c.Body)
	}
	if c.Headers != nil {
		req.SetHeaders(c.Headers)
	}
	if c.Result != nil {
		req.SetResult(c.Result)
	}

	if c.FormData != nil {
		req.SetFormData(c.FormData)
	}

	if c.QueryParam != nil {
		req.SetQueryParams(c.QueryParam)
	}

	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.Post(url)
	case "PUT":
		resp, err = req.Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	}
	return resp, err
}
