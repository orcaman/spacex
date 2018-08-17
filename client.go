package spacex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"go.uber.org/ratelimit"
)

type Client struct {
	limiter ratelimit.Limiter
}

// New creates a new spacex client
func New() *Client {
	return &Client{}
}

func NewWithRateLimit(rl int) (*Client, error) {
	if rl < 0 {
		return nil, fmt.Errorf("rate limit must be greater than 0 (got %d)", rl)
	}
	return &Client{
		limiter: ratelimit.New(rl),
	}, nil
}

// GetLaunches returns launches according to launch type and filters
func (c *Client) GetLaunches(launchType LaunchType, filters map[string]interface{}) ([]*Launch, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}

	sURL := getLaunchesURL(launchType)
	u, err := url.Parse(sURL)
	if err != nil {
		return nil, err
	}
	for k, v := range filters {
		u.Query().Add(k, fmt.Sprintf("%s", v))
	}
	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target []*Launch
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetNextLaunch returns the next launch``
func (c *Client) GetNextLaunch() (*Launch, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}

	r, err := http.Get(launchesEndpoint + "/next")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *Launch
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	return target, err
}

// GetLatestLaunch returns the latest launch
func (c *Client) GetLatestLaunch() (*Launch, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}

	r, err := http.Get(launchesEndpoint + "/latest")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *Launch
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	return target, err
}

// GetAllRockets returns all rockets info
func (c *Client) GetAllRockets() ([]*Rocket, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(rocketsEndpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target []*Rocket
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetRocket returns rocket info by name
func (c *Client) GetRocket(name string) (*Rocket, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(rocketsEndpoint + "/" + name)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *Rocket
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetAllCapsules returns all Capsules info
func (c *Client) GetAllCapsules() ([]*Capsule, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(capsulesEndpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target []*Capsule
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetCapsule returns Capsule info by name
func (c *Client) GetCapsule(name string) (*Capsule, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(capsulesEndpoint + "/" + name)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *Capsule
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetInfo returns comapny info
func (c *Client) GetInfo() (*Info, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(infoEndpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *Info
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetObjectInfo returns object info by name
func (c *Client) GetObjectInfo(name string) (*ObjectInfo, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}
	r, err := http.Get(infoEndpoint + "/" + name)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target *ObjectInfo
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}

// GetHistory returns historical events according to filters
func (c *Client) GetHistory(filters map[string]interface{}) ([]*History, error) {
	if c.limiter != nil {
		c.limiter.Take()
	}

	u, err := url.Parse(historyEndpoint)
	if err != nil {
		return nil, err
	}
	for k, v := range filters {
		u.Query().Add(k, fmt.Sprintf("%s", v))
	}
	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var target []*History
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &target)
	if err != nil {
		return nil, err
	}

	return target, err
}
