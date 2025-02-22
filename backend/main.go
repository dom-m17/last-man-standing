package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, getAllData())
	})
	r.GET("/teams", func(c *gin.Context) {
		c.JSON(http.StatusOK, getAllClubs())
	})

	return r
}

func main() {
	godotenv.Load()
	r := setupRouter()
	r.Run(":" + os.Getenv("PORT"))
}

func getAllData() interface{} {

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
	defer response.Body.Close()

	var teamInfo map[string]interface{}
	json.Unmarshal(responseData, &teamInfo)

	return teamInfo
}

func getAllClubs() interface{} {

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
	defer response.Body.Close()

	var teamsAsJSON map[string]interface{}
	json.Unmarshal(responseData, &teamsAsJSON)

	teams, ok := teamsAsJSON["teams"].([]interface{})
	if !ok {
		log.Fatal("Error parsing teams data")
	}

	var teamNames []string
	for _, team := range teams {
		teamMap, ok := team.(map[string]interface{})
		if ok {
			if name, exists := teamMap["name"].(string); exists {
				teamNames = append(teamNames, name)
			}
		}
	}

	return teamNames
}
