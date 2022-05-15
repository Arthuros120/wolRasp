package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
		"wolRasp/config"
)

func main() {

		config.Get("clientConfig.json")
	
        CONNECT := "92.128.78.135:31604"

        c, err := net.Dial("tcp", CONNECT)

        if err != nil {
                fmt.Println(err)
                return
        }

		fmt.Println("connect to server : " + CONNECT)

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                fmt.Fprintf(c, text+"\n")

                message, _ := bufio.NewReader(c).ReadString('\n')
                fmt.Print("->: " + message)
                if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
}