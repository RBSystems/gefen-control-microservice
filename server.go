package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/gefen-control-microservice/handlers"
)

func main() {
	log.SetLevel("debug")
	port := ":8024"
	router := common.NewRouter()

	// Functionality Endpoints
	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	write.GET("/:address/output/:output/input/:input", handlers.SwitchInput)

	// Status/Hardware Info Endpoints
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))
	read.GET("/:address/output/:port/input", handlers.ShowOutput)
	read.GET("/:address/hardware", handlers.HardwareInfo)

	// log level endpoints
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
