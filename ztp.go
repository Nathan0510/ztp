package main

import (
        "fmt"
//      "log"
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
        "golang.org/x/crypto/ssh"
        "strings"
        "io/ioutil"
)

type serialnumber struct {
        SN string
}

type DATACUST struct {
        LOGINPPP string
        PASSPPP string
        SN      string
}

func main(){

//      gin.SetMode(gin.ReleaseMode)
//      ztp := gin.New()
        ztp := gin.Default()
        ztp.SetTrustedProxies([]string{"192.168.1.0/24"})

        ztp.GET("/narutoooooooo", func (c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                        "message":"sasukeeeeeeee",
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

                address := ip + ":22"
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

                client, err := ssh.Dial("tcp", address, confssh)
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


        ztp.GET("/config", func (c *gin.Context){

                var json serialnumber

                if err := c.ShouldBindJSON(&json); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }

                path := "/tmp/" + json.SN

                data, err := os.ReadFile(path)
                if  err != nil {
                        fmt.Println("le fichier existe pas bg", err)
                }


                c.JSON(http.StatusOK, gin.H{
                        "le fichier": string(data),
                })


        })

        ztp.POST("/postconfig", func (c *gin.Context){

                var json DATACUST

                if err := c.ShouldBindJSON(&json); err != nil {
                        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                        return
                }

                tpl := "/tmp/tpl"
                path := "/tmp/" + json.SN

                data, err := os.ReadFile(tpl)
                if err != nil {
                        fmt.Println("le fichier existe pas bg", err)
                }

                config := strings.ReplaceAll(string(data), "LOGINPPP", json.LOGINPPP)
                config = strings.ReplaceAll(config, "PASSPPP", json.PASSPPP)

                if err := ioutil.WriteFile(path, []byte(config), 0644); err != nil {
                        fmt.Println("erreur pdt l ecriture du fichier bg", err)
                }


                c.JSON(http.StatusOK, gin.H{
                        "status": string(config),
                })

        })


        ztp.Run()
}
