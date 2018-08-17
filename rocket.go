package spacex

const rocketsEndpoint = "https://api.spacexdata.com/v2/rockets"

type Rocket struct {
	Rocketid       float64 `json:"rocketid"`
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Active         bool    `json:"active"`
	Stages         float64 `json:"stages"`
	Boosters       float64 `json:"boosters"`
	CostPerLaunch  float64 `json:"cost_per_launch"`
	SuccessRatePct float64 `json:"success_rate_pct"`
	FirstFlight    string  `json:"first_flight"`
	Country        string  `json:"country"`
	Company        string  `json:"company"`
	Height         struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"height"`
	Diameter struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"diameter"`
	Mass struct {
		Kg float64 `json:"kg"`
		Lb float64 `json:"lb"`
	} `json:"mass"`
	PayloadWeights []struct {
		ID   string  `json:"id"`
		Name string  `json:"name"`
		Kg   float64 `json:"kg"`
		Lb   float64 `json:"lb"`
	} `json:"payload_weights"`
	FirstStage struct {
		Reusable       bool    `json:"reusable"`
		Engines        float64 `json:"engines"`
		FuelAmountTons float64 `json:"fuel_amount_tons"`
		BurnTimeSec    float64 `json:"burn_time_sec"`
		ThrustSeaLevel struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust_sea_level"`
		ThrustVacuum struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust_vacuum"`
	} `json:"first_stage"`
	SecondStage struct {
		Engines        float64 `json:"engines"`
		FuelAmountTons float64 `json:"fuel_amount_tons"`
		BurnTimeSec    float64 `json:"burn_time_sec"`
		Thrust         struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust"`
		Payloads struct {
			Option1          string `json:"option_1"`
			Option2          string `json:"option_2"`
			CompositeFairing struct {
				Height struct {
					Meters float64 `json:"meters"`
					Feet   float64 `json:"feet"`
				} `json:"height"`
				Diameter struct {
					Meters float64 `json:"meters"`
					Feet   float64 `json:"feet"`
				} `json:"diameter"`
			} `json:"composite_fairing"`
		} `json:"payloads"`
	} `json:"second_stage"`
	Engines struct {
		Number         float64 `json:"number"`
		Type           string  `json:"type"`
		Version        string  `json:"version"`
		Layout         string  `json:"layout"`
		EngineLossMax  float64 `json:"engine_loss_max"`
		Propellant1    string  `json:"propellant_1"`
		Propellant2    string  `json:"propellant_2"`
		ThrustSeaLevel struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust_sea_level"`
		ThrustVacuum struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust_vacuum"`
		ThrustToWeight float64 `json:"thrust_to_weight"`
	} `json:"engines"`
	LandingLegs struct {
		Number   float64 `json:"number"`
		Material string  `json:"material"`
	} `json:"landing_legs"`
	Wikipedia   string `json:"wikipedia"`
	Description string `json:"description"`
}
