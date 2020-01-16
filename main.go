package main

import (
	"fmt"
	"log"
	"flag"
	"net"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	var listenPort = 18746

	var otddHost string
	var otddPort int
	var username string
	var tag string
	flag.StringVar(&otddHost, "h", "", "the otdd server host")
	flag.IntVar(&otddPort, "p", 8746, "the otdd server port")
	flag.StringVar(&username, "u", "", "the username")
	flag.StringVar(&tag, "t", "", "the tag")
	flag.Parse()

	if otddHost == "" {
		panic(fmt.Sprintf("please specify the otdd server host."))
	}
	if username == "" {
		panic(fmt.Sprintf("please specify the username."))
	}
	if tag == "" {
		panic(fmt.Sprintf("please specify the tag."))
	}

	macAddr := getMacAddr()
	testRunner := NewTestRunner(username,tag,macAddr,listenPort,otddHost,otddPort)

	testRunner.Start()
}

func getMacAddr() string{
	interfaces,err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for _,inter := range interfaces {
		if inter.HardwareAddr.String() != "" {
			return inter.HardwareAddr.String()
		}
	}
	return ""
}
