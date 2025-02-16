package sdkopen

import (
	"github.com/sdkopen/sdkopen-go-webserver/observer"
	"github.com/sdkopen/sdkopen-go-webserver/validator"
)

func InitializeSdkOpen() {
	validator.Initialize()
	observer.Initialize()
}
