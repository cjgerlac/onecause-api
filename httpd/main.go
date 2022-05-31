package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func main() {
	r := gin.Default()
	credUsername := "c137@onecause.com"
	credPassword := "#th@nH@rm#y#r!$100%D0p#"

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}

	r.Use(cors.New(config))

	r.POST("/loginOne", func(c *gin.Context) {
		timestamp := time.Now()
		hour := strconv.Itoa(timestamp.Hour())
		minute := strconv.Itoa(timestamp.Minute())
		newtime := hour + minute

		var loginCreds creds
		if err := c.BindJSON(&loginCreds); err != nil {
			return
		}
		if loginCreds.Username != credUsername || loginCreds.Password != credPassword || loginCreds.Token != newtime {
			c.Status(http.StatusUnauthorized)
			c.Error(fmt.Errorf("invalid credentials"))
		} else {
			c.Status(http.StatusOK)
		}
	})
	r.Run("0.0.0.0:4201")
}
