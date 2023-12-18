package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"encoding/json"
	//"bytes"
)

type RequestBody struct {
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
   // Crea un router con gin
   router := gin.Default()

   router.POST("/saludo", func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		responseString := fmt.Sprintf("Hello, %s %s!", requestBody.Nombre, requestBody.Apellido)

		c.String(http.StatusOK, responseString)
   })

   router.Run()
}