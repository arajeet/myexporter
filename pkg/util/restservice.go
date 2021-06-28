package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/arajeet/myexporter/pkg/objects"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Client is the couchbase client
type Client struct {
	domain string
	Client http.Client
}

// NewClient creates a new couchbase client
func NewClient(domain, user, password string, config *tls.Config) Client {
	var client = Client{
		domain: domain,
		Client: http.Client{
			Transport: &AuthTransport{
				Username: user,
				Password: password,
				config:   config,
			},
		},
	}
	print("retunring client")
	return client
}

// configTLS examines the configuration and creates a TLS configuration
func ConfigClientTLS(cacert, chain, key string) *tls.Config {
	tlsClientConfig := &tls.Config{
		RootCAs: x509.NewCertPool(),
	}

	caCert, err := ioutil.ReadFile(cacert)
	if err != nil {
		log.Fatal(err)
	}

	if ok := tlsClientConfig.RootCAs.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(fmt.Errorf("failed to append CA certificate"))
	}

	cert, err := tls.LoadX509KeyPair(chain, key)
	if err != nil {
		log.Fatal(err)
	}

	tlsClientConfig.Certificates = []tls.Certificate{cert}
	return tlsClientConfig
}

func (c Client) Url(path string) string {
	return c.domain + "/" + path
}

func (c Client) Get(path string, v interface{}) error {
	resp, err := c.Client.Get(c.Url(path))
	if err != nil {
		return errors.Wrapf(err, "failed to Get %s", path)
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response body from %s", path)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to Get %s metrics: %s %d", path, string(bts), resp.StatusCode)
	}

	if err := json.Unmarshal(bts, v); err != nil {
		return errors.Wrapf(err, "failed to unmarshall %s output: %s", path, string(bts))
	}
	return nil
}

// AuthTransport is a http.RoundTripper that does the authentication
type AuthTransport struct {
	Username string
	Password string
	config   *tls.Config

	Transport http.RoundTripper
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		TLSClientConfig:       t.config,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// RoundTrip implements the RoundTripper interface.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	req2.SetBasicAuth(t.Username, t.Password)
	return t.transport().RoundTrip(req2)
}

type IndexStats = objects.IndexStats

var Url string
var Username string
var Password string
var Port string
var client Client

func CallIndexstats(c Client) (m map[string]IndexStats) {

	//client := &http.Client{}

	//req, err := http.NewRequest("GET", "http://192.168.68.111:9102/api/v1/stats/", nil)
	//req, err := http.NewRequest("GET", Url, nil)
	//	req.SetBasicAuth("Administrator", "123456")
	//req.SetBasicAuth(Username, Password)
	mx := map[string]IndexStats{}
	err := c.Get("api/v1/stats/", &mx)
	if err != nil {
		log.Fatal(err)
	}
	delete(mx, "indexer")
	for k, v := range mx {
		fmt.Println(k)
		bucketnameIndex := strings.Split(k, ":")
		fmt.Println("bucketName ::  " + bucketnameIndex[0] + "\nIndex Name ::" + bucketnameIndex[1])
		fmt.Println(v.AvgDrainRate)
	}

	return mx
}
