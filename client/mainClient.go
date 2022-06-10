package main

import (

	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"wolRasp/config"

)

func main() {

		config.Get("clientConfig.json")
	
        CONNECT := config.General.Address + ":" + config.General.Port

		log.Println("Attempt to connect to " + CONNECT)

        c, err := net.Dial("tcp", CONNECT)

        if err != nil {

                log.Println(err)
                return

        }

		log.Println("Connected to server : " + CONNECT)

        log.Println("Send signal...")

        fmt.Fprintf(c, config.General.Password + "\n")

        message, _ := bufio.NewReader(c).ReadString('\n')

        if strings.TrimSpace(string(message)) == "$01" {
                        
            log.Println("Request Accepted")
            
            message, _ = bufio.NewReader(c).ReadString('\n')

            if strings.TrimSpace(string(message)) == "$01" {

                log.Println("Successful computer boot sequence")

            }else{

                log.Println("Unsuccessful computer boot sequence")

                return

            }

        }else{

            log.Println("Request Not Accepted")
            return

        }

    c.Close()
    
}