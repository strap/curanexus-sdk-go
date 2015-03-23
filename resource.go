package strap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type (
	// Resource lays out the structure for a resource in the
	// discovery route.
	Resource struct {
		Name        string   `json:"name"`
		Token       string   `json:"token"`
		Method      string   `json:"method"`
		URI         string   `json:"uri"`
		Description string   `json:"description"`
		Required    []string `json:"required,omitempty"`
		Optional    []string `json:"optional,omitempty"`
	}

	Report struct {
		ID        string `json:"id" bson:"_id"`
		Timestamp int    `json:"timestamp" bson:"timestamp"`
		Date      string `json:"date" bson:"date"`
		Type      string `json:"type" bson:"type"`
		GUID      string `json:"guid" bson:"guid"`

		Activity Activities `json:"activity" bson:"activity"`
		Food     Food       `json:"food" bson:"food"`
		Body     Body       `json:"body" bson:"body"`
		Sleep    Sleep      `json:"sleep" bson:"sleep"`

		Average    Averages      `json:"average,omitempty" bson:"average"`
		Components []*Activities `json:"components,omitempty" bson:"components"`
	}

	Averages struct {
		Calories     int    `json:"calories" bson:"calories"`
		Floors       int    `json:"floors" bson:"floors"`
		Steps        int    `json:"steps" bson:"steps"`
		ActiveMin    int    `json:"activeMinutes" bson:"activeMinutes"`
		nonActiveMin string `json:"nonactiveMinutes" bson:"nonactiveMinutes"`
		Updated      string `json:"updated" bson:"updated"`
	}

	Activities struct {
		Calories     int    `json:"calories" bson:"calories"`
		Floors       int    `json:"floors" bson:"floors"`
		Steps        int    `json:"steps" bson:"steps"`
		ActiveMin    int    `json:"activeMinutes" bson:"activeMinutes"`
		nonActiveMin string `json:"nonactiveMinutes" bson:"nonactiveMinutes"`
		Updated      string `json:"updated" bson:"updated"`
	}

	Body struct {
		BMI     string `json:"bmi" bson:"bmi"`
		BodyFat int    `json:"bodyFat" bson:"bodyFat"`
		Weight  string `json:"weight" bson:"weight"`
	}

	Food struct {
		Calories int    `json:"calories" bson:"calories"`
		Carbs    string `json:"carbs" bson:"carbs"`
		Fat      string `json:"fat" bson:"fat"`
		Fiber    string `json:"fiber" bson:"fiber"`
		Protein  string `json:"protein" bson:"protein"`
		Sodium   string `json:"sodium" bson:"sodium"`
		Water    int    `json:"water" bson:"water"`
	}

	Sleep struct {
		Asleep   int `json:"asleep" bson:"asleep"`
		Awake    int `json:"awake" bson:"awake"`
		Duration int `json:"duration" bson:"duration"`
		Start    int `json:"start" bson:"start"`
	}

	User struct {
		GUID     string `json:"guid" bson:"guid"`
		Platform string `json:"platform" bson:"platform"`
	}
)

// Call invokes an operation on the resource.
func (r *Resource) Get(params map[string]interface{}) (io.ReadCloser, error) {

	// Verify method is valid
	if r.Method == "" {
		return nil, errors.New("Invalid method")
	}

	// Pull out pieces
	route := r.URI
	method := r.Method

	// Match path parameters out of url
	regex := regexp.MustCompile("{([^{}]+)}")
	pathParams := regex.FindAllStringSubmatch(route, -1)

	// Handle each path parameter
	for _, p := range pathParams {
		replacer := p[0]
		param := p[1]
		if _, ok := params[param]; ok {
			// Replace uri with parameter
			route = strings.Replace(route, replacer, params[param].(string), -1)
			delete(params, param)
		} else {
			// GET calls can forego path parameters
			if method != "GET" {
				return nil, errors.New("Missing parameter: " + param)
			}
			route = strings.Replace(route, replacer, "", -1)
		}
	}

	// Build query string for GET calls
	var body []byte
	if method == "GET" {
		allowed := url.Values{}
		for _, name := range r.Optional {
			if _, ok := params[name]; ok {
				allowed.Add(name, params[name].(string))
			}
		}

		// Attach to route
		route = route + "?" + allowed.Encode()
	} else {
		body, _ = json.Marshal(params)
	}

	// Setup request
	req, _ := http.NewRequest(method, route, bytes.NewBuffer(body))
	req.Header.Set("X-Auth-Token", r.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	//fmt.Println(method, route, req)

	// Perform request
	res, err := client.Do(req)
	//fmt.Println(res, err)
	if res.StatusCode >= 400 {
		e := map[string]interface{}{}
		json.NewDecoder(res.Body).Decode(&e)
		return nil, errors.New(e["code"].(string))
	}

	return res.Body, err
}
