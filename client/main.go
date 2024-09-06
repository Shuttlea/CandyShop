package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"flag"
)

func httpsClient(url string, order Order) {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		return
	}
	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	if caCertPEM, err := os.ReadFile(CACertFilePath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}

	req,_ := json.Marshal(order)
	resp, err := client.Post(url,"application/json",bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		var r OrderResp
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err!=nil{
			fmt.Println("Can't unmarshall response : ", err.Error())
			return
		}
		fmt.Println(r.Thanks,"Your change is",r.Change)
	} else {
		var e ErrorResp
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err!=nil{
			fmt.Println("Can't unmarshall response : ", err.Error())
			return
		}
		fmt.Println(e.Error)
	}
}

func main() {
	money := flag.Int("m",0,"money")
	candyCount := flag.Int("c",0,"candy count")
	candyType := flag.String("k","AA","cnady type")
	flag.Parse()
	order := Order{*money, *candyType, *candyCount}
	httpsClient("https://127.0.0.1:3333", order)
}