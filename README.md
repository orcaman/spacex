# spacex [![Build Status](https://travis-ci.org/orcaman/spacex.svg?branch=master)](https://travis-ci.org/orcaman/spacex) [![GoDoc](https://godoc.org/github.com/orcaman/spacex?status.svg)](https://godoc.org/github.com/orcaman/spacex)

Go client library for the SpaceX REST API developed at [r-spacex](https://github.com/r-spacex/SpaceX-API/blob/master/docs/home.md). Because Rockets are cool.

![Space Gophers Animated Gif](./images/space_gophers_animated.gif)

image credit:
- [Ashely McNamara's fantastic collection](https://github.com/ashleymcnamara/gophers)
- [Space Gophers](https://github.com/apiarian/space-gophers)

Using this client library, you could get information about things like the date of the next launch, videos / stats about 
previous launches, data about different rockets and capsules, live stats about the roadster/star man, etc.

## Supported Entities/Endpoints

- [Launches](https://godoc.org/github.com/orcaman/spacex#Launch)
- [Rockets](https://godoc.org/github.com/orcaman/spacex#Rocket)
- [Capsules](https://godoc.org/github.com/orcaman/spacex#Capsule)
- [History](https://godoc.org/github.com/orcaman/spacex#History)
- [Roadster / Object Info](https://godoc.org/github.com/orcaman/spacex#ObjectInfo)

Rate-limiting is supported natively (use `NewWithRateLimit`). Set a rate limit of 50 to comply with the [REST API's spec](https://github.com/r-spacex/SpaceX-API/blob/master/docs/home.md).

### TODO:
- Capsule Details
- Launchpads

## Usage Example

![Space Gophers Animated Gif](./images/space_gopher.png)

Get info about falcon9 rocket (check out the `client_test.go` file for more examples):

```go
c := New()
rocket, err := c.GetRocket("falcon9")
if err != nil {
    log.Fatal(err)
}
```

Result:

```json
{
  "rocketid": 2,
  "id": "falcon9",
  "name": "Falcon 9",
  "type": "rocket",
  "active": true,
  "stages": 2,
  "boosters": 0,
  "cost_per_launch": 50000000,
  "success_rate_pct": 97,
  "first_flight": "2010-06-04",
  "country": "United States",
  "company": "SpaceX",
  "height": {
    "meters": 70,
    "feet": 229.6
  },
  "diameter": {
    "meters": 3.7,
    "feet": 12
  },
  "mass": {
    "kg": 549054,
    "lb": 1207920
  },
  "payload_weights": [
    {
      "id": "leo",
      "name": "Low Earth Orbit",
      "kg": 22800,
      "lb": 50265
    },
    {
      "id": "gto",
      "name": "Geosynchronous Transfer Orbit",
      "kg": 8300,
      "lb": 18300
    },
    {
      "id": "mars",
      "name": "Mars Orbit",
      "kg": 4020,
      "lb": 8860
    }
  ],
  "first_stage": {
    "reusable": true,
    "engines": 9,
    "fuel_amount_tons": 385,
    "burn_time_sec": 162,
    "thrust_sea_level": {
      "kN": 7607,
      "lbf": 1710000
    },
    "thrust_vacuum": {
      "kN": 8227,
      "lbf": 1849500
    }
  },
  "second_stage": {
    "engines": 1,
    "fuel_amount_tons": 90,
    "burn_time_sec": 397,
    "thrust": {
      "kN": 934,
      "lbf": 210000
    },
    "payloads": {
      "option_1": "dragon",
      "option_2": "composite fairing",
      "composite_fairing": {
        "height": {
          "meters": 13.1,
          "feet": 43
        },
        "diameter": {
          "meters": 5.2,
          "feet": 17.1
        }
      }
    }
  },
  "engines": {
    "number": 9,
    "type": "merlin",
    "version": "1D+",
    "layout": "octaweb",
    "engine_loss_max": 2,
    "propellant_1": "liquid oxygen",
    "propellant_2": "RP-1 kerosene",
    "thrust_sea_level": {
      "kN": 845,
      "lbf": 190000
    },
    "thrust_vacuum": {
      "kN": 914,
      "lbf": 205500
    },
    "thrust_to_weight": 180.1
  },
  "landing_legs": {
    "number": 4,
    "material": "carbon fiber"
  },
  "wikipedia": "https://en.wikipedia.org/wiki/Falcon_9",
  "description": "Falcon 9 is a two-stage rocket designed and manufactured by SpaceX for the reliable and safe transport of satellites and the Dragon spacecraft into orbit."
}
```