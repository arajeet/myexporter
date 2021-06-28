package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/arajeet/myexporter/pkg/collector"
	"github.com/arajeet/myexporter/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	envUser = "COUCHBASE_USER"
	envPass = "COUCHBASE_PASS"
)

var (
	couchAddr      = flag.String("couchbase-address", "localhost", "The address where Couchbase Server is running")
	couchPort      = flag.String("couchbase-port", "", "The port where Couchbase Server is running.")
	prometheusPort = flag.String("server-port", "2112", "The port where Prometheus client  Server is running.")
	userFlag       = flag.String("couchbase-username", "Administrator", "Couchbase Server Username. Overridden by env-var COUCHBASE_USER if set.")
	passFlag       = flag.String("couchbase-password", "password", "Plaintext Couchbase Server Password. Recommended to pass value via env-ver COUCHBASE_PASS. Overridden by aforementioned env-var.")
	ca             = flag.String("ca", "", "PKI certificate authority file")
	clientCert     = flag.String("client-cert", "", "client certificate file to authenticate this client with couchbase-server")
	clientKey      = flag.String("client-key", "", "client private key file to authenticate this client with couchbase-server")
	refreshTime    = flag.String("per-node-refresh", "5", "How frequently to collect per_node_bucket_stats collector in seconds")
)

func createClient(username, password string) util.Client {
	// Default to nil
	var tlsClientConfig tls.Config

	// Default to insecure
	scheme := "http"
	port := "9102"

	// Update the TLS, scheme and port
	if len(*ca) != 0 && len(*clientCert) != 0 && len(*clientKey) != 0 {
		scheme = "https"
		port = "18091"

		tlsClientConfig = tls.Config{
			RootCAs: x509.NewCertPool(),
		}

		caContents, err := ioutil.ReadFile(*ca)
		if err != nil {
			fmt.Printf("could not read CA")
			os.Exit(1)
		}
		if ok := tlsClientConfig.RootCAs.AppendCertsFromPEM(caContents); !ok {
			fmt.Printf("failed to append CA")
			os.Exit(1)
		}

		certContents, err := ioutil.ReadFile(*clientCert)
		if err != nil {
			fmt.Printf("could not read client cert")
			os.Exit(1)
		}
		key, err := ioutil.ReadFile(*clientKey)
		if err != nil {
			fmt.Printf("could not read client key")
			os.Exit(1)
		}
		cert, err := tls.X509KeyPair(certContents, key)
		if err != nil {
			fmt.Printf("failed to create X509 KeyPair")
			os.Exit(1)
		}
		tlsClientConfig.Certificates = append(tlsClientConfig.Certificates, cert)
	} else {
		if len(*clientCert) != 0 || len(*clientKey) != 0 {
			fmt.Printf("please specify both clientCert and clientKey")
			os.Exit(1)
		}
	}

	if len(*couchPort) != 0 {
		port = *couchPort
	}

	log.Info("dial CB Server at: " + scheme + "://" + *couchAddr + ":" + port)

	return util.NewClient(scheme+"://"+*couchAddr+":"+port, username, password, &tlsClientConfig)
}

func main() {
	flag.Parse()
	username := *userFlag
	password := *passFlag
	//promPort := *prometheusPort
	/*if *couchAddr == "" || *couchPort == "" {
		panic("Please pass the address of the couchbase server and port details.")
	}*/
	s := "http://" + *couchAddr + ":" + *couchPort + "/api/v1/stats/"
	util.Url = s
	if os.Getenv(envUser) != "" {
		username = os.Getenv(envUser)
		util.Username = username
	} else {
		util.Username = username
	}
	if os.Getenv(envPass) != "" {
		password = os.Getenv(envPass)
		util.Password = password

	} else {
		util.Password = password
	}
	client := createClient(username, password)
	//i, _ := strconv.Atoi(*refreshTime)
	//util.Client = client
	promPort := ":" + *prometheusPort

	print("Promethus port =" + *prometheusPort)
	prometheus.MustRegister(collector.NewIndexCollector(client))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(promPort, nil))
	print("Prometheus started")

}
