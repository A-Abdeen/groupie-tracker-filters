package main

import (
	handlers "gt/webapp/handlers"
	server "gt/webapp/server"
)

func main() {
	// Initialize html & tmpl global array for all handlers
	if handlers.HtmlTmpl == nil {
		handlers.Init()
	}
	server.Server()
}
