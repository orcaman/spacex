package spacex

import (
	"encoding/json"
	"log"
	"os/exec"
	"reflect"
	"testing"
)

func TestLaunches(t *testing.T) {
	launchTypes := []LaunchType{
		LaunchTypeUpcoming,
		LaunchTypePast,
		LaunchTypeAll}
	c := New()

	for _, l := range launchTypes {
		e, err := getExpectedResponse(l)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := c.GetLaunches(l, nil)
		if err != nil {
			t.Fatal(err)
		}

		a, err := json.Marshal(actual)
		if err != nil {
			t.Fatal(err)
		}

		eq := reflect.DeepEqual(e, a)
		if !eq {
			t.Logf("testing launch type: %v", l)
			t.Fatalf("expected:\n%s\n\nbut got:\n%s", e, a)
		}
		t.Logf("testing launch type: %v PASS", l)
	}

	e, err := getExpectedNextLaunch()
	if err != nil {
		t.Fatal(err)
	}

	actual, err := c.GetNextLaunch()
	if err != nil {
		t.Fatal(err)
	}

	a, err := json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	eq := reflect.DeepEqual(e, a)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", e, a)
	}

	e, err = getExpectedLatestLaunch()
	if err != nil {
		t.Fatal(err)
	}

	actual, err = c.GetLatestLaunch()
	if err != nil {
		t.Fatal(err)
	}

	a, err = json.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}

	eq = reflect.DeepEqual(e, a)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", e, a)
	}
}

func TestRockets(t *testing.T) {
	expectedRockets, err := getExpectedRockets()
	if err != nil {
		log.Fatal(err)
	}

	c := New()
	rockets, err := c.GetAllRockets()
	if err != nil {
		log.Fatal(err)
	}

	actual, err := json.Marshal(rockets)
	if err != nil {
		log.Fatal(err)
	}

	eq := reflect.DeepEqual(actual, expectedRockets)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedRockets, actual)
	}

	name := "falcon9"
	rocket, err := c.GetRocket(name)
	if err != nil {
		log.Fatal(err)
	}

	expectedRocket, err := getExpectedRocket(name)
	if err != nil {
		log.Fatal(err)
	}

	actual, err = json.Marshal(rocket)
	if err != nil {
		log.Fatal(err)
	}

	eq = reflect.DeepEqual(actual, expectedRocket)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedRocket, actual)
	}
}
func TestCapsules(t *testing.T) {
	expectedCapsules, err := getExpectedCapsules()
	if err != nil {
		log.Fatal(err)
	}

	c := New()
	capsules, err := c.GetAllCapsules()
	if err != nil {
		log.Fatal(err)
	}

	actual, err := json.Marshal(capsules)
	if err != nil {
		log.Fatal(err)
	}

	eq := reflect.DeepEqual(actual, expectedCapsules)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedCapsules, actual)
	}

	name := "dragon2"
	capsule, err := c.GetCapsule(name)
	if err != nil {
		log.Fatal(err)
	}

	expectedCapsule, err := getExpectedCapsule(name)
	if err != nil {
		log.Fatal(err)
	}

	actual, err = json.Marshal(capsule)
	if err != nil {
		log.Fatal(err)
	}

	eq = reflect.DeepEqual(actual, expectedCapsule)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedCapsule, actual)
	}
}

func TestInfo(t *testing.T) {
	c := New()
	info, err := c.GetInfo()
	if err != nil {
		log.Fatal(err)
	}

	expectedInfo, err := getExpectedInfo()
	if err != nil {
		log.Fatal(err)
	}

	actual, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	eq := reflect.DeepEqual(actual, expectedInfo)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedInfo, actual)
	}
}

func TestObjectInfo(t *testing.T) {
	c := New()
	name := "roadster"
	info, err := c.GetObjectInfo(name)
	if err != nil {
		log.Fatal(err)
	}

	expectedObjectInfo, err := getExpectedObjectInfo(name)
	if err != nil {
		log.Fatal(err)
	}

	actual, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	eq := reflect.DeepEqual(actual, expectedObjectInfo)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedObjectInfo, actual)
	}
}

func TestHistory(t *testing.T) {
	c := New()
	info, err := c.GetHistory(nil)
	if err != nil {
		log.Fatal(err)
	}

	expectedHistory, err := getExpectedHistory()
	if err != nil {
		log.Fatal(err)
	}

	actual, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	eq := reflect.DeepEqual(actual, expectedHistory)
	if !eq {
		t.Fatalf("expected:\n%s\n\nbut got:\n%s", expectedHistory, actual)
	}
}

func getExpectedHistory() ([]byte, error) {
	data, err := exec.Command("curl", "-s", historyEndpoint).Output()
	if err != nil {
		return nil, err
	}
	var r *[]*History
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedInfo() ([]byte, error) {
	data, err := exec.Command("curl", "-s", infoEndpoint).Output()
	if err != nil {
		return nil, err
	}
	var r *Info
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedObjectInfo(name string) ([]byte, error) {
	data, err := exec.Command("curl", "-s", infoEndpoint+"/"+name).Output()
	if err != nil {
		return nil, err
	}
	var r *ObjectInfo
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedCapsules() ([]byte, error) {
	data, err := exec.Command("curl", "-s", capsulesEndpoint).Output()
	if err != nil {
		return nil, err
	}
	var r []*Capsule
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedCapsule(name string) ([]byte, error) {
	data, err := exec.Command("curl", "-s", capsulesEndpoint+"/"+name).Output()
	if err != nil {
		return nil, err
	}
	var r *Capsule
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedRockets() ([]byte, error) {
	data, err := exec.Command("curl", "-s", rocketsEndpoint).Output()
	if err != nil {
		return nil, err
	}
	var r []*Rocket
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedRocket(name string) ([]byte, error) {
	data, err := exec.Command("curl", "-s", rocketsEndpoint+"/"+name).Output()
	if err != nil {
		return nil, err
	}
	var r *Rocket
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedNextLaunch() ([]byte, error) {
	data, err := exec.Command("curl", "-s", "https://api.spacexdata.com/v2/launches/next").Output()
	if err != nil {
		return nil, err
	}
	var l *Launch
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedLatestLaunch() ([]byte, error) {
	data, err := exec.Command("curl", "-s", "https://api.spacexdata.com/v2/launches/latest").Output()
	if err != nil {
		return nil, err
	}
	var l *Launch
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getExpectedResponse(l LaunchType) ([]byte, error) {
	var data []byte
	var err error
	switch l {
	case LaunchTypePast:
		data, err = exec.Command("curl", "-s", "https://api.spacexdata.com/v2/launches").Output()
		if err != nil {
			return nil, err
		}
		var l []*Launch
		err := json.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}
		return json.Marshal(l)
	case LaunchTypeUpcoming:
		data, err = exec.Command("curl", "-s", "https://api.spacexdata.com/v2/launches/upcoming").Output()
		if err != nil {
			return nil, err
		}
		var l []*Launch
		err := json.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}
		return json.Marshal(l)
	case LaunchTypeAll:
		data, err = exec.Command("curl", "-s", "https://api.spacexdata.com/v2/launches/all").Output()
		if err != nil {
			return nil, err
		}
		var l []*Launch
		err := json.Unmarshal(data, &l)
		if err != nil {
			return nil, err
		}
		return json.Marshal(l)
	}
	return data, nil
}
