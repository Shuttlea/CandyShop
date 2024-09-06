package main

var (
	CACertFilePath = "../certificates/minica.pem"
	CertFilePath   = "../certificates/client/cert.pem"
	KeyFilePath    = "../certificates/client/key.pem"
)

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type OrderResp struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type ErrorResp struct {
	Error string `json:"error"`
}