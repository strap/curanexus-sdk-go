package strap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/swhite24/go-debug"
)

type (
	// Strap organizes the resources
	Strap struct {
		token     string
		resources map[string]*[]*Resource
		debug     debugger.Debugger
	}

	Result struct {
		StatusCode int
		Header     http.Header
	}

	Request struct {
		Name   string
		Method string
		Params Query
	}

	Query map[string]string
)

const (
	discoveryURL = "https://api.curanexus.io/discover"
)

// New delivers a new strapSDK instances with the token setup.
func New(token string) *Strap {
	return &Strap{
		token: token,
		debug: debugger.NewDebugger("strap:strap"),
	}
}

// Discover instructs strapSDK to perform the initial discovery
func (w *Strap) Discover() {
	requestEncode(discoveryURL, w.token, &w.resources)

}

func (w *Strap) endpoints() map[string]*[]*Resource {
	return w.resources
}

func (w *Strap) Call(req Request, v interface{}) (*Result, error) {

	// Check the avialability
	resource, err := w.checkResource(&req)

	if err != nil {
		return nil, err
	}

	// Make sure the method is upper case
	req.Method = strings.ToUpper(req.Method)

	// Get the information
	data, res, err := resource.DoIt(req.Method, w.token, req.Params)

	w.debug.Log(data)

	if err == nil {
		json.NewDecoder(data).Decode(&v)
		return res, nil
	}

	return res, err

}

func (w *Strap) All(req Request, v interface{}) (*Result, error) {

	// Holder
	reports := []*Report{}

	// Kick start everything..
	res, err := w.Call(req, &reports)

	//Page values
	cur_page := 1
	if len(res.Header["X-Page"]) != 0 {
		cur_page, _ = strconv.Atoi(res.Header["X-Page"][0])
	}

	tot_page := 1
	if len(res.Header["X-Pages"]) != 0 {
		tot_page, _ = strconv.Atoi(res.Header["X-Pages"][0])
	}

	w.debug.Log("first reports", req, reports)

	for cur_page < tot_page {

		reports_temp := []*Report{}

		// bump the page
		cur_page++

		// Set the Next page on Req
		req.Params["page"] = strconv.Itoa(cur_page)

		w.debug.Log("next reports", cur_page, req, reports_temp)

		res, err = w.Call(req, &reports_temp)

		if err == nil {
			for _, r := range reports_temp {
				reports = append(reports, r)
			}
		}

		cur_page, _ = strconv.Atoi(res.Header["X-Page"][0])
		tot_page, _ = strconv.Atoi(res.Header["X-Pages"][0])
	}

	w.debug.Log("final reports", reports)

	reports_temp, _ := json.Marshal(reports)

	json.Unmarshal(reports_temp, v)

	return res, err

}

// Pull our the resource+method combination
func (w *Strap) checkResource(req *Request) (*Resource, error) {

	for key, r := range w.resources {
		for _, rr := range *r {
			if key == req.Name && rr.Method == strings.ToUpper(req.Method) {
				return rr, nil
			}
		}
	}
	return nil, errors.New("Invalid resource or method")
}

// Check Pagination
func (w *Strap) pagination(res *Resource) bool {

	if res.Optional == nil {
		return false
	}

	for _, val := range res.Optional {
		if val == "page" {
			return true
		}
	}
	return false
}

// Get the Discovery
func requestEncode(url string, token string, v interface{}) {

	// fmt.Println(token)
	// Setup base discovery request
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	req.Header.Set("X-Auth-Token", token)

	// Perform request
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	// Decode body into resources
	json.NewDecoder(res.Body).Decode(v)
}
