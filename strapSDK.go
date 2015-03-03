package strapSDK

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type (
	// StrapSDK organizes the resources
	StrapSDK struct {
		token     string
		resources map[string]*Resource
	}
)

const (
	discoveryURL = "https://api2.straphq.com/discover"
)

// New delivers a new strapSDK instances with the token setup.
func New(token string) *StrapSDK {
	fmt.Println("Strap SDK setup and using token: " + token)
	return &StrapSDK{token: token}
}

// Discover instructs strapSDK to perform the initial discovery
func (w *StrapSDK) Discover() {
	requestEncode(discoveryURL, w.token, &w.resources)
	// fmt.Println(w.resources)
}

// Send finds a root resource and calls it
func (w *StrapSDK) Send(name string, params map[string]interface{}) (io.ReadCloser, error) {
	if w.resources[name] != nil {
		// Set the Token value
		w.resources[name].Token = w.token

		// fmt.Println(w.resources)
		return w.resources[name].Call(params)
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
