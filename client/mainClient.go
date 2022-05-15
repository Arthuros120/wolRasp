package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"wolRasp/config"
)

func main() {

		config.Get("clientConfig.json")
	
        CONNECT := config.General.Address + ":" + config.General.Port

		fmt.Println("Attempt to connect to " + CONNECT)

        c, err := net.Dial("tcp", CONNECT)
        if err != nil {
                fmt.Println(err)
                return
        }

		fmt.Println("Connected to server : " + CONNECT + "\nSend signal...")

        fmt.Fprintf(c, config.General.Password + "\n")

        message, _ := bufio.NewReader(c).ReadString('\n')

        if strings.TrimSpace(string(message)) == "$01" {
                        
            fmt.Println("Request Accepted")
            
            message, _ = bufio.NewReader(c).ReadString('\n')

            if strings.TrimSpace(string(message)) == "$01" {

                fmt.Println("Successful computer boot sequence")

            }else{

                fmt.Println("Unsuccessful computer boot sequence")

                return

            }

        }else{

            fmt.Println("Request Not Accepted")
            return

        }
}