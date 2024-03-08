package main

// import (
// 	"fmt"
//   "log"
// 	"time"
// )
//
// func main() {
//   layout := "02-01-2006 15:04 (MST)"
//   str := "12-03-2024 04:42 (EET)"
// 	pt, err := time.Parse(layout, str)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(pt)


	// Define the layout based on the time format
	// // Parse the time string
	// layout := "15:04 (EET)" // Define the layout based on the time format
	// str := "12:04 (EET)"
	// parsedTime, err := time.Parse(layout, str)
	// fmt.Println(parsedTime)
	// if err != nil {
	//     fmt.Println("Error parsing time:", err)
	//     return
	// }
	//
	// // Set a default date (e.g., today's date)
	// now := time.Now()
	// defaultDate := time.Date(now.Year(), 7, now.Day(), 0, 0, 0, 0, time.UTC)
	//
	// // Combine the parsed time with the default date
	// combinedTime := time.Date(defaultDate.Year(), defaultDate.Month(), defaultDate.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, time.UTC)
	//
	// fmt.Println("Parsed time with default date:", combinedTime)
// }
