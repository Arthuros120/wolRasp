package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"wolRasp/config"
    "log"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
    "io/ioutil"
)

// Cryptage
func RsaEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
    //Décrypterpem Format de la clé publique 
    block, _ := pem.Decode(publicKey)
    if block == nil {
        return nil, errors.New("public key error")
    }
    //  Résoudre la clé publique 
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    // Assertion de type
    pub := pubInterface.(*rsa.PublicKey)
    //Cryptage
    return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func read(path string) string {

    file, err := ioutil.ReadFile(path)

    if err != nil {
        log.Fatal(err)
    }
    
    return string(file)
    

}

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

        publicKey := []byte(read(config.General.PublicKey))

        msgEncode, err := RsaEncrypt(publicKey, []byte(config.General.Password))

        log.Println(msgEncode)

        msgEncodeStr := base64.StdEncoding.EncodeToString(msgEncode)

        log.Println(msgEncodeStr)

        fmt.Fprintf(c, msgEncodeStr + "\n")

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