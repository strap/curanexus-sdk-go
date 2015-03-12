package strap

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const (
	token = "{ READ TOKEN for Project }"
)

func TestEndpoints(*testing.T) {
	strap := getStrap()

	r := strap.endpoints()
	spew.Println("getActivity: %v", r)
}

func TestActivity(*testing.T) {
	strap := getStrap()

	r, _ := strap.getActivity(map[string]interface{}{"guid": "demo-guid"})
	spew.Println("getActivity: %v", r)
}

func TestReport(*testing.T) {
	strap := getStrap()

	r, _ := strap.getReport(map[string]interface{}{})
	spew.Println("get Report: %v", r)
}

func TestToday(*testing.T) {
	strap := getStrap()

	r, _ := strap.getToday(map[string]interface{}{})
	spew.Println("getToday: %v", r)
}

func TestTrigger(*testing.T) {
	strap := getStrap()

	r, _ := strap.getTrigger(map[string]interface{}{})
	spew.Println("getTrigger: %v", r)
}

func TestUsers(*testing.T) {
	strap := getStrap()

	r, _ := strap.getUsers(map[string]interface{}{})
	spew.Println("getUsers: %v", r)
}

func getStrap() *Strap {
	strap := New(token)
	strap.Discover()
	return strap
}
