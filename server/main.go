package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"
)

var (
	CACertFilePath = "../certificates/minica.pem"
	CertFilePath   = "../certificates/127.0.0.1/cert.pem"
	KeyFilePath    = "../certificates/127.0.0.1/key.pem"
)

func main() {
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	certPool := x509.NewCertPool()
	if caCertPEM, err := os.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		Certificates: []tls.Certificate{serverTLSCert},
	}

	server := http.Server{
		Addr:      ":3333",
		Handler:   http.HandlerFunc(handlers.BuyCandy),
		TLSConfig: tlsConfig,
	}
	defer server.Close()
	fmt.Println("server is listening ...")
	log.Fatal(server.ListenAndServeTLS("", ""))
}

