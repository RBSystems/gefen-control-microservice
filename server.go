package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
)

func main() {
	port := ":8024"
	router := common.NewRouter()

	// Functionality Endpoints
	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))

	// Status/Hardware Info Endpoints
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))

	// log level endpoints
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
