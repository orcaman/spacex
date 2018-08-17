package spacex

import "time"

const (
	launchesEndpoint = "https://api.spacexdata.com/v2/launches"
)

type LaunchType int

const (
	LaunchTypePast     = LaunchType(3)
	LaunchTypeUpcoming = LaunchType(4)
	LaunchTypeAll      = LaunchType(5)
)

type Launch struct {
	FlightNumber    float64   `json:"flight_number"`
	MissionName     string    `json:"mission_name"`
	LaunchYear      string    `json:"launch_year"`
	LaunchDateUnix  float64   `json:"launch_date_unix"`
	LaunchDateUtc   time.Time `json:"launch_date_utc"`
	LaunchDateLocal string    `json:"launch_date_local"`
	Rocket          struct {
		RocketID   string `json:"rocket_id"`
		RocketName string `json:"rocket_name"`
		RocketType string `json:"rocket_type"`
		FirstStage struct {
			Cores []struct {
				CoreSerial     string  `json:"core_serial"`
				Flight         float64 `json:"flight"`
				Block          float64 `json:"block"`
				Reused         bool    `json:"reused"`
				LandSuccess    bool    `json:"land_success"`
				LandingType    string  `json:"landing_type"`
				LandingVehicle string  `json:"landing_vehicle"`
			} `json:"cores"`
		} `json:"first_stage"`
		SecondStage struct {
			Block    float64 `json:"block"`
			Payloads []struct {
				PayloadID      string    `json:"payload_id"`
				NoradID        []float64 `json:"norad_id"`
				Reused         bool      `json:"reused"`
				Customers      []string  `json:"customers"`
				Nationality    string    `json:"nationality"`
				Manufacturer   string    `json:"manufacturer"`
				PayloadType    string    `json:"payload_type"`
				PayloadMassKg  float64   `json:"payload_mass_kg"`
				PayloadMassLbs float64   `json:"payload_mass_lbs"`
				Orbit          string    `json:"orbit"`
				OrbitParams    struct {
					ReferenceSystem string    `json:"reference_system"`
					Regime          string    `json:"regime"`
					Longitude       float64   `json:"longitude"`
					SemiMajorAxisKm float64   `json:"semi_major_axis_km"`
					Eccentricity    float64   `json:"eccentricity"`
					PeriapsisKm     float64   `json:"periapsis_km"`
					ApoapsisKm      float64   `json:"apoapsis_km"`
					InclinationDeg  float64   `json:"inclination_deg"`
					PeriodMin       float64   `json:"period_min"`
					LifespanYears   float64   `json:"lifespan_years"`
					Epoch           time.Time `json:"epoch"`
					MeanMotion      float64   `json:"mean_motion"`
					Raan            float64   `json:"raan"`
				} `json:"orbit_params"`
			} `json:"payloads"`
		} `json:"second_stage"`
	} `json:"rocket"`
	Telemetry struct {
		FlightClub interface{} `json:"flight_club"`
	} `json:"telemetry"`
	Reuse struct {
		Core      bool `json:"core"`
		SideCore1 bool `json:"side_core1"`
		SideCore2 bool `json:"side_core2"`
		Fairings  bool `json:"fairings"`
		Capsule   bool `json:"capsule"`
	} `json:"reuse"`
	LaunchSite struct {
		SiteID       string `json:"site_id"`
		SiteName     string `json:"site_name"`
		SiteNameLong string `json:"site_name_long"`
	} `json:"launch_site"`
	LaunchSuccess bool `json:"launch_success"`
	Links         struct {
		MissionPatch      string      `json:"mission_patch"`
		MissionPatchSmall string      `json:"mission_patch_small"`
		RedditCampaign    string      `json:"reddit_campaign"`
		RedditLaunch      string      `json:"reddit_launch"`
		RedditRecovery    interface{} `json:"reddit_recovery"`
		RedditMedia       string      `json:"reddit_media"`
		Presskit          string      `json:"presskit"`
		ArticleLink       string      `json:"article_link"`
		Wikipedia         string      `json:"wikipedia"`
		VideoLink         string      `json:"video_link"`
	} `json:"links"`
	Details           string    `json:"details"`
	Upcoming          bool      `json:"upcoming"`
	StaticFireDateUtc time.Time `json:"static_fire_date_utc"`
}

func getLaunchesURL(launchType LaunchType) string {
	url := launchesEndpoint
	switch launchType {
	case LaunchTypeUpcoming:
		url = url + "/upcoming"
	case LaunchTypeAll:
		url = url + "/all"
	}
	return url
}
