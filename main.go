package main

import "github.com/sdkopen/sdkopen-go-webserver/server"

func init() {
	InitializeSdkOpen()
}

func main() {
	server.AddRoutes(NewController().Routes())

	server.ListenAndServe()
}
