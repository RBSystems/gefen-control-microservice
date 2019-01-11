package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

//SwitchInput takes the IP address, the output and the input from the user and
//switches the input to the one requested

func SwitchInput(address, ouput, input string) (string, *nerr.E) {
	//establish telnet connection to device
	//TODO:check input is valid before sending command (sending input 22 will change it to 2)

	conn, err := getConnection(address, true)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return "", nerr.Translate(err).Add("Telnet connection failed")
	}
	//execute telnet command to switch input
	conn.Write([]byte("s " + input + "\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}
	if strings.Contains(string(b), "OUT OF RANGE") {
		return "", nerr.Create("Input is out of range", "Error")
	}
	response := strings.Split(string(b), " ")
	log.L.Info(len(response))
	log.L.Info(response)
	return input, nil
}