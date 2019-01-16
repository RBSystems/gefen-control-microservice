package helpers

import (
	"net"
	"strings"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

//This function returns the current input that is being shown as the output
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

	log.L.Info(len(response))
	input := string(response[1])
	input = input[1:]

	log.L.Info(input)

	return input, nil
}

//This function gets the IP Address (ipaddr), Software and hardware
//version (verdata), and mac address (macaddr) of the device
func GetHardware(address string) (string, string, string, *nerr.E) {
	conn, gerr := getConnection(address, true)
	if gerr != nil {
		log.L.Errorf("Failed to get connection with %s: %s", address, gerr.Error())
		return "", "", "", nerr.Translate(gerr).Add("Telnet connection failed")
	}
	ipaddr, err := GetIPAddress(address, conn)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return "", "", "", err.Add("Telnet connection failed")
	}

	verdata, err := GetVerData(address, conn)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return "", "", "", err.Add("Telnet connection failed")
	}

	macaddr, err := GetMacAddress(address, conn)
	if err != nil {
		log.L.Errorf("Failed to establish connection with %s : %s", address, err.Error())
		return "", "", "", err.Add("Telnet connection failed")
	}
	return ipaddr, macaddr, verdata, nil
}

func GetIPAddress(address string, conn *net.TCPConn) (string, *nerr.E) {
	conn.Write([]byte("#show_ip\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}
	ipaddr := ""
	response := strings.Split(string(b), " : ")
	if len(response) >= 1 {
		ipaddr = strings.Replace(response[1], "telnet->", "", -1)
		ipaddr = strings.TrimSpace(ipaddr)
	}
	log.L.Info(ipaddr)
	return ipaddr, nil
}

//gets software and hardware data
func GetVerData(address string, conn *net.TCPConn) (string, *nerr.E) {
	conn.Write([]byte("#show_ver_data\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}
	verdata := ""
	response := strings.Split(string(b), " : ")
	if len(response) >= 1 {
		verdata = strings.Replace(response[1], "telnet->", "", -1)
		verdata = strings.TrimSpace(verdata)
	}
	log.L.Info(verdata)
	return verdata, nil
}

//gets macaddress of device
func GetMacAddress(address string, conn *net.TCPConn) (string, *nerr.E) {
	conn.Write([]byte("#show_mac_addr\r\n"))
	b, err := readUntil(CARRIAGE_RETURN, conn, 3)
	if err != nil {
		return "", nerr.Translate(err).Add("failed to read from connection")
	}
	macaddr := ""
	response := strings.Split(string(b), " : ")
	if len(response) >= 1 {
		macaddr = strings.Replace(response[1], "telnet->", "", -1)
		macaddr = strings.TrimSpace(macaddr)
	}
	log.L.Info(macaddr)
	return macaddr, nil
}
