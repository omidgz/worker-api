package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	ID int
}

func GetWorker(id int) Worker {
	return Worker{
		ID: id,
	}
}

func (w Worker) Do(ch chan Job) {
	for job := range ch {
		d := rand.Intn(10)
		time.Sleep(time.Duration(d) * time.Second)
		fmt.Printf("Job %s scheduled at %s done at %s by %d.\n", job.ID, job.At, time.Now().Format(time.UnixDate), w.ID)
	}
}
