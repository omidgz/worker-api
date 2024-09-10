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
		Log(fmt.Sprintf("Job %s done at %s by %d.\n", job.ID, time.Now().Format(time.UnixDate), w.ID))
	}
}
