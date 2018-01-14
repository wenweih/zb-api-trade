package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func depthTask() {
	depRes := depth("depth", "btc_usdt", "20")
	fmt.Println(depRes)
}

func depthTaskRun(r *receiver) {
	defer r.Done()
	gocron.Every(1).Seconds().Do(depthTask)
	<-gocron.Start()
}
