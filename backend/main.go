package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	getAllClubs()

	r := gin.Default()
	r.GET("/clubs", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func getAllClubs() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://api.football-data.org/v4/competitions/PL/teams", nil)
	req.Header.Set("X-Auth-Token", os.Getenv("FOOTBALL_DATA_API_KEY"))
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
