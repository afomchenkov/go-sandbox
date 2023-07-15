package main

type contactInfo struct {
	email   string
	zipCode int
}

type human struct {
	firstName string
	lastName  string
	contactInfo
}
