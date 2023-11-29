package xhttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

type DefaultClient http.Client

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
	var client = &DefaultClient{}
	for _, v := range ops {
		v(client)
	}
	return client
}

func (dc *DefaultClient) Request(ctx context.Context, method, link string, params interface{}, resp IResponse, ops ...Option) error {
	return Request(ctx, method, link, params, resp, ops...)
}
