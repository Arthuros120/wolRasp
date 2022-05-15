package main

import (
        "bufio"
        "fmt"
        "net"
        "strings"
		"wolRasp/config"
)

func main() {

		config.Get("serverConfig.json")

        PORT := ":" + config.General.Port

		fmt.Println("Start server to *" + PORT)

        l, err := net.Listen("tcp", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }

        defer l.Close()

        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                if strings.TrimSpace(string(netData)) == config.General.Password {
                        
					fmt.Println("Request Accepted")
					c.Write([]byte("$01"))    

                }else{

					fmt.Println("Request Not Accepted")
					c.Write([]byte("$02")) 

				}
        }
}