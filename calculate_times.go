package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)


func parseTimings() (Timings, error) {
  timeStrings := prepareTimings()

  var timings Timings
  timings.TimeTable = make(map[int][5]time.Time)

  layout := "02-01-2006 15:04 (MST)"
  for i := range timeStrings.TimeTable {
    var timesArray [5]time.Time
    for i, t := range timeStrings.TimeTable[i] {
      parsedTime, err := parseTime(layout, t)
      if err != nil {
        return timings, err
      }
      timesArray[i] = parsedTime
    }
    fmt.Println(timesArray)
  }


  return timings, nil
}

func parseTime(layout, timeString string) (time.Time, error) {
  parsedTime, err := time.Parse(layout, timeString)
  if err != nil {
    return time.Now(), err
  }
  return parsedTime, nil
}

func prepareTimings() TimeStrings {
	ttbc := fetchTimings()

	var timeStrings TimeStrings

	timeStrings.TimeTable = make(map[int][5]string)
	for i, v := range ttbc.Data {
		var timesArray [5]string

		timesArray[0] = v.Date.Gregorian.Date +" " + v.Timings.Fajr
		timesArray[1] = v.Date.Gregorian.Date +" " + v.Timings.Dhuhr
		timesArray[2] = v.Date.Gregorian.Date +" " + v.Timings.Asr
		timesArray[3] = v.Date.Gregorian.Date +" " + v.Timings.Maghrib
		timesArray[4] = v.Date.Gregorian.Date +" " + v.Timings.Isha

		timeStrings.TimeTable[i] = timesArray
	}
	return timeStrings
}

func fetchTimings() TTByCity {
	now := time.Now()
	url := fmt.Sprintf("http://api.aladhan.com/v1/calendarByCity/%v/%v?city=Giza&country=Egypt", now.Year(), int(now.Month()))
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// sb := string(body)
	//
	// log.Printf(sb)

	var ttbc TTByCity
	err = json.Unmarshal(body, &ttbc)
	if err != nil {
		log.Fatalln(err)
	}
	return ttbc
}
