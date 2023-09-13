package xhttp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DefaultClient struct {
	http.Client
}

// NewClientWithCert 创建SSL链接客户端
func NewClientWithCert(pemCert, pemKey []byte) *http.Client {
	var client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}
	cert, err := tls.X509KeyPair(pemCert, pemKey)
	if err != nil {
		panic(err)
	}

	clientCertPool := x509.NewCertPool()
	if !clientCertPool.AppendCertsFromPEM(pemCert) {
		panic("failed to parse root certificate")
	}
	conf := &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	client.Transport = &http.Transport{
		MaxIdleConnsPerHost: 100,
		TLSClientConfig:     conf,
	}
	return client
}

// NewClientDefault 创建普通客户端
func NewClientDefault() *DefaultClient {
	return &DefaultClient{Client: *http.DefaultClient}
}

func (dc *DefaultClient) Request(method, link string, params interface{}, resp interface{}, ops ...Option) error {
	log.Printf(">>>> 开始请求【link=%s】，参数【%+v】", link, func() string {
		b, _ := json.Marshal(params)
		return string(b)
	}())
	req, err := http.NewRequest(method, link, nil)
	if err != nil {
		return err
	}
	if params != nil {
		ops = append(ops, JsonBody(params))
	}
	for _, op := range ops {
		op(req)
	}
	response, err := dc.Do(req)
	if err != nil {
		log.Printf("请求【link=%s】出错，err=%s", link, err)
		return err
	}

	if response.StatusCode != 200 {
		err := fmt.Errorf("接口【link=%+v】请求错误[status_code=%d]", link, response.StatusCode)
		log.Printf("%s", err)
		return err
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("读取数据出错，err=%+v", err)
		return err
	}
	if resp == nil {
		return nil
	}
	if response.ContentLength <= DefaultRespSize && response.ContentLength > 0 {
		log.Printf("获取响应数据，responseBody=%+v", string(bodyBytes))
	}
	log.Printf(">>>> 结束请求【%s】，响应数据【ContentLength=%d】 <<<<", link, response.ContentLength)

	return json.Unmarshal(bodyBytes, &resp)
}
