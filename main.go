package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var jobs chan Job

func main() {
	router := gin.Default()
	router.GET("/log", getLogs)
	router.POST("/job", postJob)

	jobs = make(chan Job)

	for t := range 10 {
		go GetWorker(t).Do(jobs)
	}

	// wait
	router.Run("localhost:8080")
}

func getLogs(c *gin.Context) {
	c.JSON(200, JsonLog())
}

func postJob(c *gin.Context) {
	var job Job

	if err := c.BindJSON(&job); err != nil {
		c.JSON(301, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	job.At = time.Now().Format(time.UnixDate)
	Log(fmt.Sprintf("Job %s scheduled at %s by %s", job.ID, job.At, job.By))
	jobs <- job
}
