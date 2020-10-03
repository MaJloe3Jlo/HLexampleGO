package main

// Object that we operate in blockchain.
type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}
