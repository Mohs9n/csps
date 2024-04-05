package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	now := time.Now()

	db, err := bolt.Open("/home/mohsen/src/go/csps/csps.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var todayTimes [5]time.Time

	err = db.Update(func(tx *bolt.Tx) error {
		bktName := []byte("times")
		keyName := []byte(fmt.Sprintf("%v/%v", now.Year(), int(now.Month())))
		bkt := tx.Bucket(bktName)
		if bkt == nil {
			bkt, err = tx.CreateBucket(bktName)
			if err != nil {
				return fmt.Errorf("error creating bucket")
			}

		}
		// bkt = tx.Bucket(bktName)
		// if bkt == nil {
		// 	return fmt.Errorf("Bucket not there??")
		// }
		timingsJson := bkt.Get(keyName)
    if timingsJson == nil {
			timings, err := parseTimings()
			if err != nil {
				return err
			}
			tmjson, err := json.Marshal(timings)
			if err != nil {
				return err
			}

			err = bkt.Put(keyName, tmjson)
			if err != nil {
				return err
			}
      timingsJson = bkt.Get(keyName)
    }
		var timings Timings
		err = json.Unmarshal(timingsJson, &timings)
		if err != nil {
      log.Println("Error Unmarshaling json from bucket")
			return err
		}
		todayTimes = timings.TimeTable[now.Day()-1]

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	sleep(todayTimes)
}

func sleep(today_times [5]time.Time) {
	sleepDuration, err := getNextSleep(today_times)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sleeping for", sleepDuration-(10*time.Minute))
	time.Sleep(sleepDuration - (10 * time.Minute))

	cmd := exec.Command("notify-send", "-e", "-u", "critical", "10 Minutes to Next Prayer! Computer Sleeping in 5 Minutes!")
	_, err = cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Sleeping for: ", (5 * time.Minute))
	time.Sleep(5 * time.Minute)

	sleepCmd := exec.Command("systemctl", "suspend")
	_, err = sleepCmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
}
