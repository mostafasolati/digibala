package main

var lastUsedId = 0

type Product struct {
	Id              int
	Name            string
	Desc            string
	CategoryId      string
	Price           float32
	Rating          float32
	ImageURLs       []string
	StorageQuantity int
}
