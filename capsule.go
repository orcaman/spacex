package spacex

const capsulesEndpoint = "https://api.spacexdata.com/v2/capsules"

type Capsule struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Active           bool    `json:"active"`
	CrewCapacity     float64 `json:"crew_capacity"`
	SidewallAngleDeg float64 `json:"sidewall_angle_deg"`
	OrbitDurationYr  float64 `json:"orbit_duration_yr"`
	HeatShield       struct {
		Material    string  `json:"material"`
		SizeMeters  float64 `json:"size_meters"`
		TempDegrees float64 `json:"temp_degrees"`
		DevPartner  string  `json:"dev_partner"`
	} `json:"heat_shield"`
	Thrusters []struct {
		Type   string  `json:"type"`
		Amount float64 `json:"amount"`
		Pods   float64 `json:"pods"`
		Fuel1  string  `json:"fuel_1"`
		Fuel2  string  `json:"fuel_2"`
		Thrust struct {
			KN  float64 `json:"kN"`
			Lbf float64 `json:"lbf"`
		} `json:"thrust"`
	} `json:"thrusters"`
	LaunchPayloadMass struct {
		Kg float64 `json:"kg"`
		Lb float64 `json:"lb"`
	} `json:"launch_payload_mass"`
	LaunchPayloadVol struct {
		CubicMeters float64 `json:"cubic_meters"`
		CubicFeet   float64 `json:"cubic_feet"`
	} `json:"launch_payload_vol"`
	ReturnPayloadMass struct {
		Kg float64 `json:"kg"`
		Lb float64 `json:"lb"`
	} `json:"return_payload_mass"`
	ReturnPayloadVol struct {
		CubicMeters float64 `json:"cubic_meters"`
		CubicFeet   float64 `json:"cubic_feet"`
	} `json:"return_payload_vol"`
	PressurizedCapsule struct {
		PayloadVolume struct {
			CubicMeters float64 `json:"cubic_meters"`
			CubicFeet   float64 `json:"cubic_feet"`
		} `json:"payload_volume"`
	} `json:"pressurized_capsule"`
	Trunk struct {
		TrunkVolume struct {
			CubicMeters float64 `json:"cubic_meters"`
			CubicFeet   float64 `json:"cubic_feet"`
		} `json:"trunk_volume"`
		Cargo struct {
			SolarArray         float64 `json:"solar_array"`
			UnpressurizedCargo bool    `json:"unpressurized_cargo"`
		} `json:"cargo"`
	} `json:"trunk"`
	HeightWTrunk struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"height_w_trunk"`
	Diameter struct {
		Meters float64 `json:"meters"`
		Feet   float64 `json:"feet"`
	} `json:"diameter"`
	Wikipedia   string `json:"wikipedia"`
	Description string `json:"description"`
}
