package server

import (
	"context"
	"mime/multipart"

	"github.com/sdkopen/sdkopen-go-webserver/http"
)

type WebContext interface {
	Context() context.Context
	RequestHeader(key string) []string
	RequestHeaders() map[string][]string
	PathParam(key string) string
	RawQuery() string
	QueryParam(key string) string
	QueryArrayParam(key string) []string
	DecodeQueryParams(object any) error
	DecodeBody(object any) error
	DecodeFormData(object any) error
	StringBody() (string, error)
	Path() string
	Method() string
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	AddHeader(key string, value string)
	AddHeaders(headers map[string]string)
	Redirect(url string, statusCode http.StatusCode)
	ServeFile(path string)
	JsonResponse(statusCode http.StatusCode, body any)
	ErrorResponse(statusCode http.StatusCode, err error)
	EmptyResponse(statusCode http.StatusCode)
}
