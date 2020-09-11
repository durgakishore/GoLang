package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	ipAddress              string
	port                   string
	serverNetworkInterface string
	symbolIBMMap           = make(map[int]string)
	symbolBACMap           = make(map[int]string)
)

func main() {

	//-multicast 225.1.1.7 65071 192.168.195.126

	cmdArguments := os.Args

	if !processCmdArguments(cmdArguments) {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	startUDPClientProcessing(wg)
	wg.Wait()
}

func processCmdArguments(cmdArguments []string) bool {
	if len(cmdArguments) < 4 {
		fmt.Println("Please enter all the required parameters")
		return false
	}

	fmt.Println(cmdArguments)

	if cmdArguments[1] == "-multicast" {
		ipAddress = cmdArguments[2]
	} else {
		fmt.Println("Please enter the destination IPAddress in a valid commandline format.")
		return false
	}
	fmt.Println(ipAddress)
	if len(cmdArguments[3]) > 0 {
		port = cmdArguments[3]
	} else {
		fmt.Println("Port should not be empty")
		return false
	}
	fmt.Println(port)
	if len(cmdArguments[4]) > 0 {
		serverNetworkInterface = cmdArguments[4]
	} else {
		fmt.Println("Server network interface should not be empty")
		return false
	}
	fmt.Println(serverNetworkInterface)

	return true
}

func startUDPClientProcessing(wg *sync.WaitGroup) {
	address := serverNetworkInterface + ":" + port
	s, err := net.ResolveUDPAddr("udp4", address)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	defer c.Close()

	data := []byte("trying to send a sample request to server\n")
	_, err = c.Write(data)

	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	fmt.Printf("UDP server is connected and server address is %s\n", c.RemoteAddr().String())
	for {
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}
		//fmt.Printf("Received Data : %s\n", string(buffer[0:n]))

		/*
			3=1388|4=1388|5=IBM|8=157.58|9=197|1021=600|854=n|474=157.666449179|1584=143133|444=20160713|459=95838066963000|461=18393|55=95838067160000|22=154068|361=0.61|981=0|16=1468418318.083|18=1468418318.083|362=0.39,
			3=1388|4=1388|5=BAC|8=157.58|9=197|1021=600|854=n|474=157.666449179|1584=143133|444=20160713|459=95838066963000|461=18393|55=95838067160000|22=154068|361=0.61|981=0|16=1468418318.083|18=1468418318.083|362=0.39,
			3=1388|4=1388|5=IBM|8=157.58|9=100|1021=600|854=n|474=157.666425537|1584=143233|444=20160713|459=95838070277000|461=18395|55=95838070535000|22=154171|361=0.61|981=0|16=1468418318.0937|18=1468418318.0937|362=0.39,
			3=1388|4=1388|5=BAC|8=157.58|9=100|1021=600|854=n|474=157.666425537|1584=143233|444=20160713|459=95838070277000|461=18395|55=95838070535000|22=154171|361=0.61|981=0|16=1468418318.0937|18=1468418318.0937|362=0.39
		*/

		var resultIBM []string
		var resultBAC []string
		data := string(buffer[0:n])
		result := strings.Split(data, ",")
		for v := range result {
			if strings.Contains(result[v], "IBM") {
				resultIBM = strings.Split(result[v], "|")
			}
			if strings.Contains(result[v], "BAC") {
				resultBAC = strings.Split(result[v], "|")
			}
			for v := range resultIBM {
				resultIBM := strings.Split(resultIBM[v], "=")
				if !strings.Contains(resultIBM[1], "IBM") {
					token, _ := strconv.Atoi(resultIBM[0])
					value := resultIBM[1]
					symbolIBMMap[token] = value
				}
			}
			for v := range resultBAC {
				resultBAC := strings.Split(resultBAC[v], "=")
				if !strings.Contains(resultBAC[1], "BAC") {
					token, _ := strconv.Atoi(resultBAC[0])
					value := resultBAC[1]
					symbolBACMap[token] = value
				}
			}
		}

		if len(symbolIBMMap) > 0 {
			fmt.Println("\nIBM :  Token 5 and 8 values are")
			for k, v := range symbolIBMMap {
				fmt.Println(k, "  ", v)
			}

		}
		if len(symbolBACMap) > 0 {
			fmt.Println("\nBAC :  Token 5 and 8 values are")
			for k, v := range symbolBACMap {
				fmt.Println(k, "  ", v)
			}

		}
	}
	wg.Done()
}
