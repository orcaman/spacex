package spacex

import "time"

type ObjectInfo struct {
	Name            string    `json:"name"`
	LaunchDateUtc   time.Time `json:"launch_date_utc"`
	LaunchDateUnix  int       `json:"launch_date_unix"`
	LaunchMassKg    float64   `json:"launch_mass_kg"`
	LaunchMassLbs   float64   `json:"launch_mass_lbs"`
	NoradID         float64   `json:"norad_id"`
	EpochJd         string    `json:"epoch_jd"`
	OrbitType       string    `json:"orbit_type"`
	ApoapsisAu      float64   `json:"apoapsis_au"`
	PeriapsisAu     float64   `json:"periapsis_au"`
	SemiMajorAxisAu float64   `json:"semi_major_axis_au"`
	Eccentricity    float64   `json:"eccentricity"`
	Inclination     float64   `json:"inclination"`
	Longitude       float64   `json:"longitude"`
	PeriapsisArg    float64   `json:"periapsis_arg"`
	PeriodDays      float64   `json:"period_days"`
	SpeedKph        float64   `json:"speed_kph"`
	SpeedMph        float64   `json:"speed_mph"`
	EarthDistanceKm float64   `json:"earth_distance_km"`
	EarthDistanceMi float64   `json:"earth_distance_mi"`
	MarsDistanceKm  float64   `json:"mars_distance_km"`
	MarsDistanceMi  float64   `json:"mars_distance_mi"`
	Wikipedia       string    `json:"wikipedia"`
	Details         string    `json:"details"`
}
