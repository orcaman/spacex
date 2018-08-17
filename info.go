package spacex

const infoEndpoint = "https://api.spacexdata.com/v2/info"

type Info struct {
	Name          string `json:"name"`
	Founder       string `json:"founder"`
	Founded       int    `json:"founded"`
	Employees     int    `json:"employees"`
	Vehicles      int    `json:"vehicles"`
	LaunchSites   int    `json:"launch_sites"`
	TestSites     int    `json:"test_sites"`
	Ceo           string `json:"ceo"`
	Cto           string `json:"cto"`
	Coo           string `json:"coo"`
	CtoPropulsion string `json:"cto_propulsion"`
	Valuation     int64  `json:"valuation"`
	Headquarters  struct {
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
	} `json:"headquarters"`
	Summary string `json:"summary"`
}
