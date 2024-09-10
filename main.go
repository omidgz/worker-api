package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

var jobs chan Job

func main() {
	router := gin.Default()
	// router.GET("/job", getJobs)
	router.POST("/job", postJob)

	jobs = make(chan Job)

	for t := range 10 {
		go GetWorker(t).Do(jobs)
	}

	// wait
	router.Run("localhost:8080")
}

func postJob(c *gin.Context) {
	var job Job

	if err := c.BindJSON(&job); err != nil {
		// write err to pl
		return
	}

	job.At = time.Now().Format(time.UnixDate)
	jobs <- job
}
