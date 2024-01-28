package deepread

import (
	"go-deepread/init/server"
	"net/http"
)

// DefaultDRAPIHost 默认DeepRead API Host
const DefaultDRAPIHost = "http://127.0.0.1:8080"

type options struct {
	DRAPIHost string
	HTTP      *http.Client
}

// CtorOption 客户端对象构造参数
type CtorOption interface {
	applyTo(*options)
}

// impl Default for options
func defaultOptions() options {
	return options{
		DRAPIHost: server.ServerConf.ServerUrl,
		HTTP:      &http.Client{},
	}
}

//
//
//

type withDRAPIHost struct {
	x string
}

func WithDRAPIHost(host string) CtorOption {
	return &withDRAPIHost{x: host}
}

var _ CtorOption = (*withDRAPIHost)(nil)

func (x *withDRAPIHost) applyTo(y *options) {
	y.DRAPIHost = x.x
}

//
//
//

type withHTTPClient struct {
	x *http.Client
}

// WithHTTPClient 使用给定的 http.Client 作为 HTTP 客户端
func WithHTTPClient(client *http.Client) CtorOption {
	return &withHTTPClient{x: client}
}

var _ CtorOption = (*withHTTPClient)(nil)

func (x *withHTTPClient) applyTo(y *options) {
	y.HTTP = x.x
}
