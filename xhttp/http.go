package xhttp

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

type DefaultClient struct {
	http.Client
	l *log.Logger
}

type rqlog struct {
	Ts     int64       `json:"t"`
	Tm     time.Time   `json:"ft"`
	Dur    int         `json:"dur"`
	Link   string      `json:"link"`
	Header interface{} `json:"header"`
	Body   interface{} `json:"body"`
	Resp   interface{} `json:"resp"`
}

func (rl *rqlog) String() string {
	b, _ := json.Marshal(rl)
	return string(b)
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
func NewClientDefault(logger *log.Logger) *DefaultClient {
	return &DefaultClient{
		Client: *http.DefaultClient,
		l: logger,
	}
}

func (dc *DefaultClient) JsonRequest(link, method string, header map[string]string, data interface{}, i interface{}) error {
	var (
		rq rqlog
		ts = time.Now().Nanosecond()
		te = time.Now().Nanosecond()
	)
	defer func() {
		rq.Header = header
		rq.Body = data
		rq.Resp = i
		rq.Link = link
		rq.Ts = time.Now().Unix()
		rq.Tm = time.Now()
		rq.Dur = (te - ts) / 1000000
		if dc.l == nil {
			return
		}
		dc.l.Println(rq.String())
		return
	}()
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, link, bytes.NewBuffer(b))
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := dc.Do(req)
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	err = json.Unmarshal(b, i)
	te = time.Now().Nanosecond()
	return err
}

func (dc *DefaultClient) XmlRequest(link, method string, header map[string]string, data interface{}, i interface{}) error {
	var (
		rq rqlog
		ts = time.Now().Nanosecond()
		te = time.Now().Nanosecond()
	)
	defer func() {
		rq.Header = header
		rq.Body = data
		rq.Resp = i
		rq.Link = link
		rq.Ts = time.Now().Unix()
		rq.Tm = time.Now()
		rq.Dur = (te - ts) / 1000000
		dc.l.Println(rq.String())
	}()
	b, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, link, bytes.NewBuffer(b))
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	req.Header.Set("Content-Type", "application/xml")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := dc.Do(req)
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		te = time.Now().Nanosecond()
		return err
	}
	err = xml.Unmarshal(b, i)
	te = time.Now().Nanosecond()
	return err
}
