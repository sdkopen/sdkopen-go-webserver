package sdkopen_webserver

import (
	sdkopenWebBase "github.com/sdkopen/sdkopen-go-webbase"
	"github.com/sdkopen/sdkopen-go-webbase/server"
	"github.com/sdkopen/sdkopen-go-webserver/webserver"
)

func init() {
	sdkopenWebBase.Initialize()
}

func Initialize() {
	server.ListenAndServe(webserver.CreateChiServer)
}
