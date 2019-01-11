package handlers

import (
	"net/http"

	"github.com/byuoitav/common/status"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/gefen-control-microservice/helpers"
	"github.com/labstack/echo"
)

func SwitchInput(context echo.Context) error {
	output := context.Param("output")
	input := context.Param("input")
	address := context.Param("address")

	resp, err := helpers.SwitchInput(address, output, input)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, status.Input{Input: resp})
}

func ShowOutput(context echo.Context) error {
	output := context.Param("output")
	address := context.Param("address")
	resp, err := helpers.GetOutput(address, output)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, status.Input{Input: resp})
}
