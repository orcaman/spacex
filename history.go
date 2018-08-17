package spacex

import "time"

const historyEndpoint = "https://api.spacexdata.com/v2/info/history"

type History struct {
	Title         string    `json:"title"`
	EventDateUtc  time.Time `json:"event_date_utc"`
	EventDateUnix int       `json:"event_date_unix"`
	FlightNumber  int       `json:"flight_number"`
	Details       string    `json:"details"`
	Links         struct {
		Reddit    string `json:"reddit"`
		Article   string `json:"article"`
		Wikipedia string `json:"wikipedia"`
	} `json:"links"`
}
