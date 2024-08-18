package main

import (
        "log"
        "fmt"
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
        "golang.org/x/crypto/ssh"
)

type serialnumber struct {
        SN string
}


func main(){
        ztp := gin.Default()
        ztp.SetTrustedProxies([]string{"192.168.1.0/24"})

        ztp.POST("/ztp", func (c *gin.Context){
                var json serialnumber

      if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

                ip := c.ClientIP()
                c.JSON(http.StatusOK, gin.H{
                        "ton serial number": json.SN,
                        "ton IP": ip,
                })

                path := "/tmp/" + json.SN

        data, err := os.ReadFile(path)
        if  err != nil {
                panic(err)
        }
        fmt.Print(string(data))


        confssh := &ssh.ClientConfig{
        User: "admin",
        Auth: []ssh.AuthMethod{
                                ssh.Password("admin"),
        },

        HostKeyCallback: ssh.InsecureIgnoreHostKey(),

        }

        client, err := ssh.Dial("tcp", ip, confssh)
        if err != nil {
                log.Fatal("connexion failed bg", err)
        }
        defer client.Close()

        session, err := client.NewSession()
        if err != nil {
                log.Fatalf("session failed bg", err)
        }
        defer session.Close()


        output, err := session.CombinedOutput(string(data))
        if err != nil {
                log.Fatalf("command failed bg", err)
        }

    fmt.Printf("Output:\n", output)

        })


        ztp.Run()
}
