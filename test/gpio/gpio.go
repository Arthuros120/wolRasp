package main
 
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
 
    "github.com/mlgd/gpio"
)
 
func main() {
    // Ouverture du port 23 en mode OUT
    pin, err := gpio.OpenPin(gpio.GPIO24, gpio.ModeOutput)

    if err != nil {
        fmt.Printf("Error opening pin! %s\n", err)
        return
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
        os.Exit(0)
    }()
    
    pin.Set()
    // Attente d’une seconde
    time.Sleep(time.Second)
    // Extinction de la LED
    pin.Clear()
}