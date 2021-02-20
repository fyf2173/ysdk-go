package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	http.Client
	l log.Logger
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

// doPost 提交请求
func (c *Client) DoRequest(method, link string, header map[string]string, data interface{}, i interface{}) error {
	b, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, link, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(b, i)
	c.l.Println(saveLogger(link, header, data, i))
	return err
}

func saveLogger(link string, header map[string]string, data interface{}, i interface{}) string {
	var wplog struct {
		Ts     int64       `json:"t"`
		Tm     time.Time   `json:"ft"`
		Link   string      `json:"link"`
		Header interface{} `json:"header"`
		Body   interface{} `json:"body"`
		Resp   interface{} `json:"resp"`
	}
	wplog.Ts = time.Now().Unix()
	wplog.Tm = time.Now()
	wplog.Link = link
	wplog.Header = header
	wplog.Body = data
	wplog.Resp = i
	b, _ := json.Marshal(wplog)
	return string(b)
}
