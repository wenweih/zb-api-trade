package main

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/jasonlvhit/gocron"
)

var (
	dbTrade respTrades
	dbDepth respDepth
)

//==================================================//
func writeTask(r *receiver) {
	r.Lock()
	depRes := depth("depth", "eos_usdt", "20")
	tradeRes := trades("trades", "eos_usdt")
	r.Unlock()
	db.Batch(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("TokenTrade"))
		if err != nil {
			Exit(err.Error())
		}

		// save depth data
		depthEncodeed, err := json.Marshal(depRes)
		if err != nil {
			Exit(err.Error())
		}
		b.Put([]byte("depth"), depthEncodeed)

		// save trade data
		TradeEncodeed, err := json.Marshal(tradeRes)
		if err != nil {
			Exit(err.Error())
		}
		b.Put([]byte("trade"), TradeEncodeed)

		return nil
	})
}

func writeTaskRun(r *receiver) {
	defer r.Done()
	gocron.Every(1).Seconds().Do(writeTask, r)
	<-gocron.Start()
}

//==================================================//
func readTask() {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("TokenTrade")).Cursor()

		_, tradeValue := c.Seek([]byte("trade"))
		_, depthValue := c.Seek([]byte("depth"))

		json.Unmarshal(tradeValue, &dbTrade)
		json.Unmarshal(depthValue, &dbDepth)

		return nil
	})
}

func readTaskRun(r *receiver) {
	defer r.Done()
	gocron.Every(1).Seconds().Do(readTask)
	<-gocron.Start()
}
