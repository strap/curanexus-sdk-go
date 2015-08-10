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
	"time"

	"github.com/swhite24/go-debug"
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
		AvgFood    Food          `json:"avgfood" bson:"avgfood"`
		AvgBody    Body          `json:"avgbody" bson:"avgbody"`
		AvgSleep   Sleep         `json:"avgsleep" bson:"avgsleep"`
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
		Gender   string `json:"gender,omitempty" bson:"gender"`
		GUID     string `json:"guid" bson:"guid"`
		Platform string `json:"platform" bson:"platform"`
	}

	Segmentation map[string]interface{}

	Trigger struct {
		ID         string `json:"id" bson:"_id,omitempty"`
		Key        string `json:"key" bson:"key"`
		Range      string `json:"range" bson:"range"`
		ActionType string `json:"actionType" bson:"actionType"`
		ActionURL  string `json:"actionUrl" bson:"actionUrl"`
	}

	Job struct {
		ID              string       `json:"id" bson:"_id,omitempty"`
		CreatedAt       time.Time    `json:"createdAt" bson:"createdAt"`
		UpdatedAt       time.Time    `json:"updatedAt" bson:"updatedAt"`
		Name            string       `json:"name" bson:"name"`
		Description     string       `json:"description" bson:"description"`
		NotificationUrl string       `json:"notificationUrl" bson:"notificationUrl"`
		Status          string       `json:"status" bson:"status"`
		StartDate       string       `json:"startDate" bson:"startDate,omitempty"`
		EndDate         string       `json:"endDate" bson:"endDate,omitempty"`
		Guids           []string     `json:"guids" bson:"guids"`
		Log             []*LogRecord `json:"logs" bson:"logs,omitempty"`
	}

	LogRecord struct {
		Status    string    `json:"status" bson:"status"`
		UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	}
)

// Call invokes an operation on the resource.
func (r *Resource) DoIt(method string, token string, params Query) (io.ReadCloser, *Result, error) {

	debug := debugger.NewDebugger("strap-sdk:resources")

	// Verify method is valid
	if r.Method == "" {
		return nil, nil, errors.New("Invalid method")
	}

	// Pull out pieces
	route := r.URI

	// Match path parameters out of url
	regex := regexp.MustCompile("{([^{}]+)}")
	pathParams := regex.FindAllStringSubmatch(route, -1)

	// Handle each path parameter
	for _, p := range pathParams {
		replacer := p[0]
		param := p[1]
		if _, ok := params[param]; ok {
			// Replace uri with parameter
			route = strings.Replace(route, replacer, params[param], -1)
			delete(params, param)
		} else {
			// GET calls can forego path parameters
			if method != "GET" {
				return nil, nil, errors.New("Missing parameter: " + param)
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
				allowed.Add(name, params[name])
			}
		}

		// Attach to route
		route = route + "?" + allowed.Encode()
	} else {
		body, _ = json.Marshal(params)
	}

	debug.Log(method, route, token)

	// Setup request
	req, _ := http.NewRequest(method, route, bytes.NewBuffer(body))
	req.Header.Set("X-Auth-Token", token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// Perform request
	res, err := client.Do(req)

	// Results
	result := &Result{res.StatusCode, res.Header}

	debug.Log(result, res.Body, err)

	//fmt.Println(res, err)
	if res.StatusCode >= 400 {
		e := map[string]interface{}{}
		json.NewDecoder(res.Body).Decode(&e)
		return res.Body, result, errors.New("Error processing request")
	}

	return res.Body, result, err
}
