package main

import (
	_ "link-shortener/conn"
	"link-shortener/routing"
)

func main() {
	routing.StartHttp()
}