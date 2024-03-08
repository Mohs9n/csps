package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	now := time.Now()

	// fmt.Println(ttbc.Data[0].Timings.Fajr)
	// fmt.Printf("%v",ttbc)
	ttbc := fetchTimings()

	today_prayers := ttbc.Data[now.Day()].Timings
	const layout = "15:04 (EET)"
	parsedTime, err := time.Parse(layout, today_prayers.Fajr)
	if err != nil {
		log.Fatalln(err)
	}

	defaultDate := time.Date(now.Year(), now.Month(), 8, 0, 0, 0, 0, time.UTC)

	// Combine the parsed time with the default date
	combinedTime := time.Date(defaultDate.Year(), defaultDate.Month(), defaultDate.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, time.UTC)

	fmt.Println("Parsed time with default date:", combinedTime)

	// location, err := time.LoadLocation("Africa/Cairo") // Adjust based on actual EET location
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	times_string := make([]string, 5)
	times_string[0] = today_prayers.Fajr
	times_string[1] = today_prayers.Dhuhr
	times_string[2] = today_prayers.Asr
	times_string[3] = today_prayers.Maghrib
	times_string[4] = today_prayers.Isha


	times := make([]time.Time, 5)

	for _, e := range times_string {
		ptime, err := time.Parse(layout, e)
		if err != nil {
			log.Fatalln(err)
		}
		pptime := time.Date(now.Year(), now.Month(), now.Day(), ptime.Hour(), ptime.Minute(), ptime.Second(), ptime.Nanosecond(), now.Location())
		times = append(times, pptime)
	}

	cmd := exec.Command("notify-send", "-i", "/home/mohsen/Downloads/dotnet.svg", "-e", "-u", "critical", "hello")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	outstr := string(out)
	log.Println(outstr, "Notif")

	before := time.Now()
	time.Sleep(1 * time.Second)
	after := time.Now()

	time.Sleep(after.Sub(before))
	afterafter := time.Now()
	fmt.Printf("%v\n", afterafter.Sub(before))
	fmt.Printf("%v\n", combinedTime.Sub(now))

  parseTimings()
}
