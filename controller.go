package main

import (
	"github.com/sdkopen/sdkopen-go-webserver/http"
	"github.com/sdkopen/sdkopen-go-webserver/server"
)

type Controller struct {
}

type Response struct {
	Status string `json:"status"`
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Routes() []server.Route {
	return []server.Route{
		{
			URI:      "v1/status",
			Method:   http.MethodGet,
			Function: c.FindById,
			Prefix:   server.PublicApi,
		},
	}
}

func (c *Controller) FindById(ctx server.WebContext) {
	ctx.JsonResponse(http.StatusOK, Response{
		Status: "OK",
	})
}
