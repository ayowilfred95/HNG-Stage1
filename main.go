package main

import (
	"time"

	"github.com/gin-gonic/gin"
)


// create a struct to define our custom data 

type Parameter struct {
	SlackName      string    `json:"slack_name"`
	CurrentDay     string    `json:"current_day"`
	UTCTime        string    `json:"utc_time"`
	Track          string    `json:"track"`
	GitHubFileURL  string    `json:"github_file_url"`
	GitHubRepoURL  string    `json:"github_repo_url"`
	StatusCode     int       `json:"status_code"`

}


func main() {
	// Gin Web Framework was used to speed the server just like
	// Express
	router := gin.Default()

	// Endpoint
	router.GET("/api", func(c *gin.Context) {
		slackName := c.Query("slack_name")
		track := c.Query("track")

		currentDay := time.Now().Weekday().String()
		utcTime := time.Now().UTC()

		// UTC time within +/-2 minutes
		timeWindow := 2 * time.Minute
		currentUTC := time.Now().UTC()
		timeUTC := currentUTC.Sub(utcTime)

		if timeUTC > timeWindow || timeUTC < -timeWindow {
			c.JSON(400, gin.H{"error" : "UTC time is not within +/-2 minutes"})
			return
		}

		result := Parameter{
			SlackName: slackName,
			CurrentDay: currentDay,
			UTCTime: utcTime.Format(time.RFC3339),
			Track: track,
			GitHubFileURL: "https://github.com/ayowilfred95/hng-stage1",
			GitHubRepoURL: "https://github.com/ayowilfred95/hng-stage1",
			StatusCode: 200,
		}

		c.JSON(200,result)

	})

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
	router.Run() 
}
