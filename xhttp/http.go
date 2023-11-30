package xhttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/fyf2173/ysdk-go/xlog"
)

type DefaultClient struct {
	http.Client
}

type ClientOption func(client *DefaultClient)

func ClientCertOption(cert, key []byte) ClientOption {
	return func(client *DefaultClient) {
		transport := http.Transport{MaxIdleConnsPerHost: 100}
		tmpCert, err := tls.X509KeyPair(cert, key)
		if err != nil {
			panic(err)
		}

		clientCertPool := x509.NewCertPool()
		if !clientCertPool.AppendCertsFromPEM(cert) {
			panic("failed to parse root certificate")
		}
		transport.TLSClientConfig = &tls.Config{
			RootCAs:            clientCertPool,
			Certificates:       []tls.Certificate{tmpCert},
			InsecureSkipVerify: true,
		}
		client.Transport = &transport
	}
}

func NewClientDefault(ops ...ClientOption) *DefaultClient {
	var client = &DefaultClient{Client: http.Client{}}
	for _, v := range ops {
		v(client)
	}
	return client
}

func (dc *DefaultClient) Request(ctx context.Context, method, link string, params interface{}, resp IResponse, ops ...Option) error {
	xlog.Info(ctx, fmt.Sprintf(">>> 开始请求【[%s]link=%s】", method, link), slog.Any("params", params))
	req, err := http.NewRequest(method, link, nil)
	if err != nil {
		return err
	}
	ops = append(ops, SetTraceId(ctx))
	if params != nil {
		ops = append(ops, JsonBody(params))
	}
	for _, op := range ops {
		op(req)
	}
	response, err := dc.Client.Do(req)
	if err != nil {
		xlog.Error(ctx, err, slog.String("method", method), slog.String("link", link))
		return err
	}

	if response.StatusCode != 200 {
		xlog.Info(ctx, fmt.Sprintf("[%s]%s:%d", method, link, response.StatusCode))
		return fmt.Errorf("errorstatus:%d", response.StatusCode)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		xlog.Error(ctx, err)
		return err
	}
	if response.ContentLength <= DefaultRespSize && response.ContentLength > 0 {
		xlog.Info(ctx, "trace response", slog.String("response", string(bodyBytes)))
	}
	xlog.Info(ctx, fmt.Sprintf(">>> 结束请求[%s]%s", method, link), slog.Int64("content_length", response.ContentLength))
	if resp == nil {
		return nil
	}
	return resp.Unmarshal(bodyBytes, resp)
}
