package handlers

import (
	"doc-sync/entities"
	"doc-sync/sync_api"
	"doc-sync/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ParseHandler(c *gin.Context) {
	if ReqValidator(c) {
		var providerRequest entities.ProviderRequest
		specialty := c.Query("specialty")
		dateStr := c.Query("date")
		minScoreStr := c.Query("minScore")
		date, err := strconv.ParseInt(dateStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date parameter"})
			return
		}

		minScore, err := strconv.ParseFloat(minScoreStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minScore parameter"})
			return
		}

		// Create ProviderRequest struct
		providerRequest = entities.ProviderRequest{
			Specialty: specialty,
			Date:      date,
			MinScore:  minScore,
		}
		//filePath := "C:\\Users\\AhronRosenboim\\GolandProjects\\doc-sync\\providers.json"
		jsonData, err := ioutil.ReadFile("C:\\Users\\AhronRosenboim\\GolandProjects\\doc-sync\\providers.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		relevantProviders, err := sync_api.FilterProviders(jsonData, providerRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		relevantNames := ProviderRelevantNames(relevantProviders)

		// Respond with the result
		c.JSON(http.StatusOK, gin.H{"result": relevantNames})
	}
}
func ReqValidator(c *gin.Context) bool {
	specialty := c.Query("specialty")
	if specialty == "" || !utils.IsValidSpecialties(specialty) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong  parameter 'specialty'"})
		return false
	}

	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter 'date'"})
		return false
	}

	minScoreStr := c.Query("minScore")
	if minScoreStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter 'minScore'"})
		return false
	}

	minScore, err := strconv.ParseFloat(minScoreStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'minScore' format"})
		return false
	}

	if !utils.IsValidScore(minScore) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'minScore' must be between 0 and 10"})
		return false
	}

	return true
}
func ProviderRelevantNames(providers []entities.Provider) []string {
	var relevantNames []string
	for _, p := range providers {
		relevantNames = append(relevantNames, p.Name)
	}
	return relevantNames
}
