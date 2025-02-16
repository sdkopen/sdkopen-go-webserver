package server

import "github.com/sdkopen/sdkopen-go-webserver/http"

type RoutePrefix string

const (
	PublicApi  RoutePrefix = "/public/"
	PrivateApi RoutePrefix = "/private/"
	AdminApi   RoutePrefix = "/admin/"
	Api        RoutePrefix = "/api/"
)

type Route struct {
	URI      string
	Method   http.Method
	Prefix   RoutePrefix
	Function func(ctx WebContext)
}
