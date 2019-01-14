package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

func GetOutput(address, output string) (string, *nerr.E) {
	conn, err := getConnection(address, true)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return "", nerr.Translate(err).Add("Telnet connection failed")
	}
	conn.Write([]byte("n 1" + "\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}
	response := strings.Split(string(b), "\r\n")

	//TODO: make it return just the number not "A#"
	log.L.Info(len(response))
	input := string(response[1])
	input = input[1:]

	log.L.Info(input)

	return input, nil
}
