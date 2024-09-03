user@routeur:~/go/ztp$ cat last_ztp.go
package main

import (
        "fmt"
//      "log"
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
        "golang.org/x/crypto/ssh"
)

type serialnumber struct {
        SN string
}


func main(){

//      gin.SetMode(gin.ReleaseMode)
//      ztp := gin.New()
        ztp := gin.Default()
        ztp.SetTrustedProxies([]string{"192.168.1.0/24"})

        ztp.GET("/ping", func (c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                        "message":"pong",
                })
        })

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
                fmt.Println("le fichier existe pas bg", err)
        }


        confssh := &ssh.ClientConfig{
                User: "admin",
                Auth: []ssh.AuthMethod{
                        ssh.Password("admin"),
                },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),

        }

        client, err := ssh.Dial("tcp", "192.168.1.2:22", confssh)
        if err != nil {
                fmt.Println("connexion failed bg", err)
        }
        defer client.Close()

        session, err := client.NewSession()
        if err != nil {
                fmt.Println("session failed bg", err)
        }
        defer session.Close()

    if err := session.Run(string(data)); err != nil {
         fmt.Println("command failed bg", err)
    }



        })

        ztp.Run()
}
