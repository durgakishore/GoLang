package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	inputPath              string
	destIPAddress          string
	port                   string
	serverNetworkInterface string
	numOfPacketsPerSecond  int
	maxMessagesPerPacket   int
	symbol                 = "ALL"

	nPacketsSentIndex = 0
)

type packetData struct{}

func main() {

	/*[Server.exe -infile input.txt -multicast 225.1.1.7 65071 192.168.195.4 -packetcount 3 --mmp 5 -symbol IBM]
	  [    0 		  1       2          3         4       5          6             7     8    9  10    11   12]*/

	cmdArguments := os.Args

	if !processCmdArguments(cmdArguments) {
		return
	}

	txtFileData := readInputFileData(inputPath)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	startSendingUDPPacket(txtFileData, wg)
	wg.Wait()
}

func processCmdArguments(cmdArguments []string) bool {

	if len(cmdArguments) < 10 {
		fmt.Println("Please enter all the required parameters")
		return false
	}

	fmt.Println(cmdArguments)

	if cmdArguments[1] == "-infile" {
		inputPath = cmdArguments[2]
	} else {
		fmt.Println("Please enter the input file in a valid commandline format.")
		return false
	}

	fmt.Println(inputPath)
	if cmdArguments[3] == "-multicast" {
		destIPAddress = cmdArguments[4]
	} else {
		fmt.Println("Please enter the destination IPAddress in a valid commandline format.")
		return false
	}

	fmt.Println(destIPAddress)
	if len(cmdArguments[5]) > 0 {
		port = cmdArguments[5]
	} else {
		fmt.Println("Port should not be empty")
		return false
	}
	fmt.Println(port)
	if len(cmdArguments[6]) > 0 {
		serverNetworkInterface = cmdArguments[6]
	} else {
		fmt.Println("Server network interface should not be empty")
		return false
	}
	fmt.Println(serverNetworkInterface)
	if cmdArguments[7] == "-packetcount" {
		numOfPacketsPerSecond, _ = strconv.Atoi(cmdArguments[8])
	} else {
		fmt.Println("Please enter a valid packet count per second.")
		return false
	}
	fmt.Println(numOfPacketsPerSecond)
	if cmdArguments[9] == "--mmp" {
		maxMessagesPerPacket, _ = strconv.Atoi(cmdArguments[10])
	} else {
		fmt.Println("Please enter a valid max messages per packet.")
		return false
	}
	fmt.Println(maxMessagesPerPacket)
	if len(cmdArguments) > 11 {
		symbol = cmdArguments[12]
	}
	fmt.Println(symbol)
	return true
}

func readInputFileData(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var fileData []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		str = strings.Replace(str, "", "", -1) // to remove the ETX symbol from the line

		if symbol == "ALL" || strings.Contains(str, symbol) {
			fileData = append(fileData, str)
		}
	}
	return fileData
}

func startSendingUDPPacket(txtFileData []string, wg *sync.WaitGroup) {

	address := serverNetworkInterface + ":" + port
	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	_, addr, err := connection.ReadFromUDP(buffer)
	fmt.Printf("Client connected  %s\n", addr.String())
	var data []byte

	for {

		if len(txtFileData) <= nPacketsSentIndex {
			break
		}

		if (len(txtFileData) - nPacketsSentIndex) < maxMessagesPerPacket {
			slice := txtFileData[nPacketsSentIndex+1:]
			data = []byte(strings.Join(slice, ","))
		} else {
			slice := txtFileData[nPacketsSentIndex+1 : nPacketsSentIndex+maxMessagesPerPacket]
			data = []byte(strings.Join(slice, ","))
		}
		nPacketsSentIndex = nPacketsSentIndex + maxMessagesPerPacket

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}
	}

	fmt.Printf("Total  Messages in input files : %d\n", len(txtFileData))
	fmt.Printf("Max Messages per packet : %d\n", maxMessagesPerPacket)
	fmt.Printf("Total Packets sent: %d ", (len(txtFileData) / maxMessagesPerPacket))

	wg.Done()
}
