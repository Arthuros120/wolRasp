package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"wolRasp/config"
    "github.com/mlgd/gpio"
)

var count = 0

func handleConnection(c net.Conn) {

        log.Println("New Connection")

        for {

            netData, err := bufio.NewReader(c).ReadString('\n')

            if err != nil{

                if err == io.EOF{

					log.Println("Connection completed")
					
				}

			log.Println(err)

                break

            }

            temp := strings.TrimSpace(string(netData))

			if temp == config.General.Password {

				log.Println("Request Accepted")

				c.Write([]byte("$01\n"))

				if startComputer() == nil{

					c.Write([]byte("$01\n"))

				}else{

					c.Write([]byte("$02\n"))

				}

            }else{

				log.Println("Request Not Accepted")

				c.Write([]byte("$02\n")) 

			}
        }

        c.Close()

}

func startComputer() error{
    // Ouverture du port 23 en mode OUT
    pin, err := gpio.OpenPin(gpio.GPIO24, gpio.ModeOutput)

    if err != nil {

        log.Printf("Error opening pin! %s\n", err)
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

func main() {

    config.Get("/home/arks/code/wolRasp/serverConfig.json")

    PORT := ":" + config.General.Port

    log.Println("Start server to *" + PORT)

    l, err := net.Listen("tcp4", PORT)

    if err != nil {

            fmt.Println(err)
            return

    }

    defer l.Close()

    for{

        c, err := l.Accept()

        if err != nil {

            log.Println(err)
            return

        }

        go handleConnection(c)

        count++

    }
}