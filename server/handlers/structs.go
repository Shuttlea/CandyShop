package handlers

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type OrderReq struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type ErrorReq struct {
	Error string `json:"error"`
}