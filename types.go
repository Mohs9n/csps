package main

import "time"

type TimeStrings struct {
  TimeTable map[int][5]string
}

type Timings struct {
  TimeTable map[int][5]time.Time
}

type TTByCity struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		Timings struct {
			Fajr       string `json:"Fajr"`
			Sunrise    string `json:"Sunrise"`
			Dhuhr      string `json:"Dhuhr"`
			Asr        string `json:"Asr"`
			Sunset     string `json:"Sunset"`
			Maghrib    string `json:"Maghrib"`
			Isha       string `json:"Isha"`
			Imsak      string `json:"Imsak"`
			Midnight   string `json:"Midnight"`
			Firstthird string `json:"Firstthird"`
			Lastthird  string `json:"Lastthird"`
		} `json:"timings"`
		Date struct {
			Readable  string `json:"readable"`
			Timestamp string `json:"timestamp"`
			Gregorian struct {
				Date    string `json:"date"`
				Format  string `json:"format"`
				Day     string `json:"day"`
				Weekday struct {
					En string `json:"en"`
				} `json:"weekday"`
				Month struct {
					Number int    `json:"number"`
					En     string `json:"en"`
				} `json:"month"`
				Year        string `json:"year"`
				Designation struct {
					Abbreviated string `json:"abbreviated"`
					Expanded    string `json:"expanded"`
				} `json:"designation"`
			} `json:"gregorian"`
			Hijri struct {
				Date    string `json:"date"`
				Format  string `json:"format"`
				Day     string `json:"day"`
				Weekday struct {
					En string `json:"en"`
					Ar string `json:"ar"`
				} `json:"weekday"`
				Month struct {
					Number int    `json:"number"`
					En     string `json:"en"`
					Ar     string `json:"ar"`
				} `json:"month"`
				Year        string `json:"year"`
				Designation struct {
					Abbreviated string `json:"abbreviated"`
					Expanded    string `json:"expanded"`
				} `json:"designation"`
				Holidays []any `json:"holidays"`
			} `json:"hijri"`
		} `json:"date"`
		Meta struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Timezone  string  `json:"timezone"`
			Method    struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Params struct {
					Fajr float64 `json:"Fajr"`
					Isha float64 `json:"Isha"`
				} `json:"params"`
				Location struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"location"`
			} `json:"method"`
			LatitudeAdjustmentMethod string `json:"latitudeAdjustmentMethod"`
			MidnightMode             string `json:"midnightMode"`
			School                   string `json:"school"`
			Offset                   struct {
				Imsak    int `json:"Imsak"`
				Fajr     int `json:"Fajr"`
				Sunrise  int `json:"Sunrise"`
				Dhuhr    int `json:"Dhuhr"`
				Asr      int `json:"Asr"`
				Maghrib  int `json:"Maghrib"`
				Sunset   int `json:"Sunset"`
				Isha     int `json:"Isha"`
				Midnight int `json:"Midnight"`
			} `json:"offset"`
		} `json:"meta"`
	} `json:"data"`
}
