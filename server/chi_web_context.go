package server

import (
	"context"
	"encoding/json"
	"github.com/sdkopen/sdkopen-go-webserver/validator"
	"io"
	"mime/multipart"
	netHttp "net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/sdkopen/sdkopen-go-webserver/http"
	"github.com/sdkopen/sdkopen-go-webserver/logging"
)

type chiWebContext struct {
	writer  netHttp.ResponseWriter
	request *netHttp.Request
}

func (ctx *chiWebContext) Context() context.Context {
	return ctx.request.Context()
}

func (ctx *chiWebContext) RequestHeader(key string) []string {
	return ctx.request.Header[key]
}

func (ctx *chiWebContext) RequestHeaders() map[string][]string {
	return ctx.request.Header
}

func (ctx *chiWebContext) PathParam(key string) string {
	return chi.URLParam(ctx.request, key)
}

func (ctx *chiWebContext) RawQuery() string {
	return ctx.request.URL.RawQuery
}

func (ctx *chiWebContext) QueryParam(key string) string {
	return ctx.request.URL.Query().Get(key)
}

func (ctx *chiWebContext) QueryArrayParam(key string) []string {
	result := []string{}
	for _, value := range ctx.request.URL.Query()[key] {
		result = append(result, strings.Split(value, ",")...)
	}

	return result
}

func (ctx *chiWebContext) DecodeQueryParams(value any) error {
	if err := validator.FormDecode(value, ctx.request.URL.Query()); err != nil {
		return err
	}

	return validator.Struct(value)
}

func (ctx *chiWebContext) DecodeBody(value any) error {
	body, err := io.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, value); err != nil {
		return err
	}

	return validator.Struct(value)
}

func (ctx *chiWebContext) DecodeFormData(value any) error {
	if err := ctx.request.ParseForm(); err != nil {
		return err
	}

	if err := validator.FormDecode(value, ctx.request.PostForm); err != nil {
		return err
	}

	return validator.Struct(value)
}

func (ctx *chiWebContext) StringBody() (string, error) {
	body, err := io.ReadAll(ctx.request.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (ctx *chiWebContext) Path() string {
	return ctx.request.URL.Path
}

func (ctx *chiWebContext) Method() string {
	return ctx.request.Method
}

func (ctx *chiWebContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.request.FormFile(key)
}

func (ctx *chiWebContext) AddHeader(key string, value string) {
	ctx.writer.Header().Add(key, value)
}

func (ctx *chiWebContext) AddHeaders(headers map[string]string) {
	for key, value := range headers {
		ctx.AddHeader(key, value)
	}
}

func (ctx *chiWebContext) Redirect(url string, statusCode http.StatusCode) {
	netHttp.Redirect(ctx.writer, ctx.request, url, statusCode.Int())
}

func (ctx *chiWebContext) ServeFile(path string) {
	netHttp.ServeFile(ctx.writer, ctx.request, path)
}

func (ctx *chiWebContext) JsonResponse(statusCode http.StatusCode, body any) {
	ctx.writer.Header().Add(http.ContentTypeHeader, http.ContentTypeJSON.String())
	ctx.writer.WriteHeader(statusCode.Int())

	bytesBody, err := http.JsonEncoder(body)
	if err != nil {
		ctx.ErrorResponse(http.StatusInternalServerError, err)
	}

	ctx.writer.Write(bytesBody)
}

func (ctx *chiWebContext) ErrorResponse(statusCode http.StatusCode, err error) {
	logging.Error("[%s] %s (%d): %v", ctx.request.Method, ctx.request.RequestURI, statusCode, err)
	ctx.JsonResponse(statusCode, &http.Error{Error: err.Error()})
}

func (ctx *chiWebContext) EmptyResponse(statusCode http.StatusCode) {
	ctx.writer.WriteHeader(statusCode.Int())
}
