package helpers

import (
	"fmt"
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
	//close connection when returned
	defer conn.Close()
	//execute telnet command to switch input
	log.L.Info(len(input))
	conn.Write([]byte("s " + input + "\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}

	if strings.Contains(string(b), "OUT OF RANGE") {
		return "", nerr.Create("Input is test of range", "Error")
	}

	response := strings.Split(fmt.Sprintf("%s", b), " ")
	log.L.Infof("response: '%s'", response)

	return fmt.Sprintf("%s", input), nil
}
