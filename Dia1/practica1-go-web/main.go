package main

import(
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

func main(){

	jsonFile := "products.json"
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	err = json.Unmarshal(content, &products)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var foundItem Product
		found := false
		for _, item := range products {
			if item.ID == id {
				foundItem = item
				found = true
				break
			}
		}
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, foundItem)
	})

	router.GET("/products/search", func(c *gin.Context) {
		priceGtParam := c.Query("priceGt")

		priceGt, err := strconv.ParseFloat(priceGtParam, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priceGt parameter"})
			return
		}

		var matchingItems []Product
		for _, item := range products {
			if item.Price > priceGt {
				matchingItems = append(matchingItems, item)
			}
		}

		c.JSON(http.StatusOK, matchingItems)
	})

	router.Run(":8080")
}