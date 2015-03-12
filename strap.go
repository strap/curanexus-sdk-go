package strap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	// Strap organizes the resources
	Strap struct {
		token     string
		resources map[string]*Resource
	}
)

const (
	discoveryURL = "https://api2.straphq.com/discover"
)

// New delivers a new strapSDK instances with the token setup.
func New(token string) *Strap {
	fmt.Println("Strap SDK setup and using token: " + token)
	return &Strap{token: token}
}

// Discover instructs strapSDK to perform the initial discovery
func (w *Strap) Discover() {
	requestEncode(discoveryURL, w.token, &w.resources)
	// fmt.Println(w.resources)
}

func (w *Strap) endpoints() map[string]*Resource {
	return w.resources
}

func (w *Strap) getActivity(params map[string]interface{}) ([]*Report, error) {

	tt := []*Report{}

	if w.resources["activity"] != nil {
		// Set the Token value
		w.resources["activity"].Token = w.token

		dd, err := w.resources["activity"].Call(params)

		if err == nil {
			json.NewDecoder(dd).Decode(&tt)
			return tt, nil
		}
		return tt, err
	}
	return tt, errors.New("Could not find resource.")
}

func (w *Strap) getReport(params map[string]interface{}) (Report, error) {

	tt := Report{}

	if w.resources["report"] != nil {
		// Set the Token value
		w.resources["report"].Token = w.token

		dd, err := w.resources["report"].Call(params)

		if err == nil {
			json.NewDecoder(dd).Decode(&tt)
			return tt, nil
		}
		return tt, err
	}
	return tt, errors.New("Could not find resource.")
}

func (w *Strap) getToday(params map[string]interface{}) ([]*Report, error) {

	tt := []*Report{}

	if w.resources["today"] != nil {
		// Set the Token value
		w.resources["today"].Token = w.token

		dd, err := w.resources["today"].Call(params)

		if err == nil {
			json.NewDecoder(dd).Decode(&tt)
			return tt, nil
		}
		return nil, err
	}
	return tt, errors.New("Could not find resource.")
}

func (w *Strap) getTrigger(params map[string]interface{}) ([]*Report, error) {

	tt := []*Report{}

	if w.resources["trigger"] != nil {
		// Set the Token value
		w.resources["trigger"].Token = w.token

		dd, err := w.resources["trigger"].Call(params)

		if err == nil {
			json.NewDecoder(dd).Decode(&tt)
			return tt, nil
		}
		return tt, err
	}
	return nil, errors.New("Could not find resource.")
}

func (w *Strap) getUsers(params map[string]interface{}) ([]*User, error) {

	tt := []*User{}

	if w.resources["users"] != nil {
		// Set the Token value
		w.resources["users"].Token = w.token

		dd, err := w.resources["users"].Call(params)

		if err == nil {
			json.NewDecoder(dd).Decode(&tt)
			return tt, nil
		}
		return tt, err
	}
	return nil, errors.New("Could not find resource.")
}

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
