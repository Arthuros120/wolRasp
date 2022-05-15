package main

import (
        "bufio"
        "fmt"
        "net"
        "strings"
		"wolRasp/config"
		"os"
		"os/signal"
		"syscall"
		"time"
		"github.com/mlgd/gpio"
		"strconv"
)

var count = 0

func handleConnection(c net.Conn) {

        fmt.Print("New Connection")

        for {
			
            netData, err := bufio.NewReader(c).ReadString('\n')
            if err != nil {
                    fmt.Println(err)
                    return
            }

            temp := strings.TrimSpace(string(netData))

			if temp == config.General.Password {

				fmt.Println("Request Accepted")
				c.Write([]byte("$01\n"))

				if startComputer() == nil{

					c.Write([]byte("$01\n"))

				}else{

					c.Write([]byte("$02\n"))

				}

            }else{

				fmt.Println("Request Not Accepted")
				c.Write([]byte("$02\n")) 

			}
        }

        c.Close()

}

func main() {
        
		config.Get("serverConfig.json")

        PORT := ":" + config.General.Port

		fmt.Println("Start server to *" + PORT)

        l, err := net.Listen("tcp4", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        for {

                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }

                go handleConnection(c)

                count++

        }
}

func startComputer() error{
    // Ouverture du port 23 en mode OUT
    pin, err := gpio.OpenPin(gpio.GPIO24, gpio.ModeOutput)

    if err != nil {

        fmt.Printf("Error opening pin! %s\n", err)
        return err

    }

    // Création d’une variable pour l’interception du signal de fin de programme
    c := make(chan os.Signal, 1)

    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    signal.Notify(c, syscall.SIGKILL)

    // Go routine (thread parallèle) d’attente de fin du programme
    // pour l’extinction de la LED et la fermeture du port
    go func() {
        <-c
        pin.Clear()
        pin.Close()
    }()
    
    pin.Set()
    // Attente d’une seconde
    time.Sleep(time.Second)
    // Extinction de la LED
    pin.Clear()

	return nil
}