package main

type Framework struct {
	Name       string
	ImportUrl  string
	Url        string
	Middleware []Middleware
}
