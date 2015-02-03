package main

type Framework struct {
	Name        string
	ImportUrl   string
	Middlewares []Middleware
}
