package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/structs"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/gefen-control-microservice/helpers"
	"github.com/labstack/echo"
)

func SwitchInput(context echo.Context) error {
	output := context.Param("output")

	outport, _ := strconv.Atoi(output)
	outport = outport + 1

	input := context.Param("input")

	inport, _ := strconv.Atoi(input)
	inport = inport + 1

	address := context.Param("address")

	resp, err := helpers.SwitchInput(address, string(outport), string(inport))
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return context.JSON(http.StatusInternalServerError, err)
	}

	//decrement response by 1
	response, _ := strconv.Atoi(resp)
	response = response - 1
	return context.JSON(http.StatusOK, status.Input{Input: string(response)})
}

func ShowOutput(context echo.Context) error {
	output := context.Param("output")
	address := context.Param("address")

	//increment output by 1
	temp, _ := strconv.Atoi(output)
	port := temp + 1
	log.L.Info("The port number is %v ", port)
	resp, err := helpers.GetOutput(address, string(port))
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return context.JSON(http.StatusInternalServerError, err)
	}
	input, _ := strconv.Atoi(resp)
	input = input - 1
	return context.JSON(http.StatusOK, status.Input{Input: string(input)})
}

func HardwareInfo(context echo.Context) error {
	address := context.Param("address")
	ipaddr, macaddr, verdata, err := helpers.GetHardware(address)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, structs.HardwareInfo{
		NetworkInfo: structs.NetworkInfo{
			IPAddress:  ipaddr,
			MACAddress: macaddr,
		},
		FirmwareVersion: verdata,
	})
}
