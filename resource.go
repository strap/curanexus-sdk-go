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

		Activity ReportActivities `json:"activity" bson:"activity"`
		Food     ReportFood       `json:"food" bson:"food"`
		Body     ReportBody       `json:"body" bson:"body"`
		Sleep    ReportSleep      `json:"sleep" bson:"sleep"`

		Average    ReportAverages      `json:"average,omitempty" bson:"average"`
		AvgFood    ReportFood          `json:"avgfood" bson:"avgfood"`
		AvgBody    ReportBody          `json:"avgbody" bson:"avgbody"`
		AvgSleep   ReportSleep         `json:"avgsleep" bson:"avgsleep"`
		Components []*ReportActivities `json:"components,omitempty" bson:"components"`
	}

	ReportAverages struct {
		Calories     int    `json:"calories" bson:"calories"`
		Floors       int    `json:"floors" bson:"floors"`
		Steps        int    `json:"steps" bson:"steps"`
		ActiveMin    int    `json:"activeMinutes" bson:"activeMinutes"`
		nonActiveMin string `json:"nonactiveMinutes" bson:"nonactiveMinutes"`
		Updated      string `json:"updated" bson:"updated"`
	}

	ReportActivities struct {
		Calories     int    `json:"calories" bson:"calories"`
		Floors       int    `json:"floors" bson:"floors"`
		Steps        int    `json:"steps" bson:"steps"`
		ActiveMin    int    `json:"activeMinutes" bson:"activeMinutes"`
		nonActiveMin string `json:"nonactiveMinutes" bson:"nonactiveMinutes"`
		Updated      string `json:"updated" bson:"updated"`
	}

	ReportBody struct {
		BMI     string `json:"bmi" bson:"bmi"`
		BodyFat int    `json:"bodyFat" bson:"bodyFat"`
		Weight  string `json:"weight" bson:"weight"`
	}

	ReportFood struct {
		Calories int    `json:"calories" bson:"calories"`
		Carbs    string `json:"carbs" bson:"carbs"`
		Fat      string `json:"fat" bson:"fat"`
		Fiber    string `json:"fiber" bson:"fiber"`
		Protein  string `json:"protein" bson:"protein"`
		Sodium   string `json:"sodium" bson:"sodium"`
		Water    int    `json:"water" bson:"water"`
	}

	ReportSleep struct {
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

	// Trigger represents the Trigger structure
	Trigger struct {
		ID         string       `json:"id" bson:"_id,omitempty"`
		Active     bool         `json:"active" bson:"active"`
		Name       string       `json:"name" bson:"name"`
		Type       string       `json:"type" bson:"type"`
		Key        string       `json:"key" bson:"key"`
		Range      string       `json:"range" bson:"range"`
		Created    time.Time    `json:"created" bson:"created"`
		Conditions []*Condition `json:"conditions" bson:"conditions"`
		ActionType string       `json:"actionType" bson:"actionType"`
		ActionURL  string       `json:"actionUrl" bson:"actionUrl"`
	}

	// Condition is the entire state necessary to trigger an ActionType event
	Condition struct {
		Section    string `json:"section" bson:"section"`
		Field      string `json:"field" bson:"field"`
		Comparison string `json:"comparison" bson:"comparison"`
		Threshold  int    `json:"threshold" bson:"threshold"`
	}

	TriggerData struct {
		Key   string  `json:"key" bson:"key"`
		Users []*User `json:"users" bson:"users"`
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

	Segmentation map[string]interface{}

	Trend map[string]interface{}

	JobData map[string]interface{}

	WordCloud struct {
		GUID    string                 `json:"guid,omitempty"`
		Food    map[string]interface{} `json:"food,omitempty" bson:"food"`
		Foods   map[string]interface{} `json:"foods,omitempty" bson:"foods"`
		Workout map[string]interface{} `json:"workout,omitempty" bson:"workout"`
		Brand   map[string]interface{} `json:"brand,omitempty" bson:"brand"`
		Recipe  map[string]interface{} `json:"recipe,omitempty" bson:"recipe"`
	}

	Workout struct {
		ID        string `json:"id" bson:"_id"`
		Date      string `json:"date" bson:"date"`
		CreatedAt int64  `json:"createdAt" bson:"createdAt"`
		UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`

		Name        string `json:"name" bson:"name"`
		Description string `json:"description" bson:"description"`
		Type        string `json:"type" bson:"type"`
		StartTime   int64  `json:"startTime" bson:"startTime"`

		Country  string   `json:"country" bson:"country"`
		State    string   `json:"state" bson:"state"`
		City     string   `json:"city" bson:"city"`
		StartLoc []string `json:"startLoc" bson:"startLoc"`
		EndLoc   []string `json:"endLoc" bson:"endLoc"`

		Distance float64 `json:"distance" bson:"distance"`
		Steps    int     `json:"steps" bson:"steps"`
		Calories int     `json:"calories" bson:"calories"`

		ActiveMinutes    int     `json:"activeMinutes" bson:"activeMinutes"`
		NonactiveMinutes int     `json:"nonactiveMinutes" bson:"nonactiveMinutes"`
		MovingTime       int     `json:"movingTime" bson:"movingTime"`
		ElapsedTime      int     `json:"elapsedTime" bson:"elapsedTime"`
		AvgHeartRate     float64 `json:"avgHeartRate" bson:"avgHeartRate"`
		MaxHeartRate     float64 `json:"maxHeartRate" bson:"maxHeartRate"`
		AvgSpeed         float64 `json:"avgSpeed" bson:"avgSpeed"`
		MaxSpeed         float64 `json:"maxSpeed" bson:"maxSpeed"`
		AvgTemp          float64 `json:"avgTemp" bson:"avgTemp"`
	}

	Food struct {
		ID        string `json:"id" bson:"_id"`
		Date      string `json:"date" bson:"date"`
		CreatedAt int64  `json:"createdAt" bson:"createdAt"`
		UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`

		Name     string `json:"name" bson:"name"`
		Brand    string `json:"brand" bson:"brand"`
		Amount   string `json:"amount" bson:"amount"`
		Unit     string `json:"unit" bson:"unit"`
		MealType string `json:"mealType" bson:"mealType"`
		Barcode  string `json:"barcode" bson:"barcode"`

		Calories int     `json:"calories" bson:"calories"`
		Carbs    float64 `json:"carbs" bson:"carbs"`
		Fat      float64 `json:"fat" bson:"fat"`
		Fiber    float64 `json:"fiber" bson:"fiber"`
		Protein  float64 `json:"protein" bson:"protein"`
		Sodium   float64 `json:"sodium" bson:"sodium"`
		Water    float64 `json:"water" bson:"water"`

		TransFat       float64 `json:"transFat" bson:"transFat"`
		SaturatedFat   float64 `json:"saturatedFat" bson:"saturatedFat"`
		UnsaturatedFat float64 `json:"unsaturatedFat" bson:"unsaturatedFat"`
		MonoFat        float64 `json:"monoFat" bson:"monoFat"`
		PolyFat        float64 `json:"polyFat" bson:"polyFat"`
		VitaminC       float64 `json:"vitaminC" bson:"vitaminC"`
		VitaminA       float64 `json:"vitaminA" bson:"vitaminA"`
		Sugar          float64 `json:"sugar" bson:"sugar"`
		Potassium      float64 `json:"potassium" bson:"potassium"`
		Calcium        float64 `json:"calcium" bson:"calcium"`
		Iron           float64 `json:"iron" bson:"iron"`
		Cholesterol    float64 `json:"cholesterol" bson:"cholesterol"`
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
