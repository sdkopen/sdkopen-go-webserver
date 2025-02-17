package sdkopen

import (
	"github.com/sdkopen/sdkopen-go-webserver/observer"
	"github.com/sdkopen/sdkopen-go-webserver/server"
	"github.com/sdkopen/sdkopen-go-webserver/validator"
)

func init() {
	validator.Initialize()
	observer.Initialize()
}

func InitializeSdkOpenServer() {
	server.ListenAndServe(server.CreateChiServer)
}
