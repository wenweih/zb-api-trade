package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/jasonlvhit/gocron"
)

func depthTask() {
	depRes := depth("depth", "btc_usdt", "20")
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("depth"))
		if err != nil {
			Exit(err.Error())
		}
		encodeed, err := json.Marshal(depRes)
		if err != nil {
			Exit(err.Error())
		}
		b.Put([]byte("depth"), encodeed)
		return nil
	})
}

func readDepthTask() {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("depth"))
		v := b.Get([]byte("depth"))
		fmt.Printf("%sn", v)
		return nil
	})
}
func readDepthTaskRun(r *receiver) {
	defer r.Done()
	gocron.Every(1).Seconds().Do(readDepthTask)
	<-gocron.Start()
}

func depthTaskRun(r *receiver) {
	defer r.Done()
	gocron.Every(1).Seconds().Do(depthTask)
	<-gocron.Start()
}
