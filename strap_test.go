package strap

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const (
	token = "{ READ TOKEN for Project }"
)

func TestEndpoints(*testing.T) {
	w := getStrap()

	r := w.endpoints()
	spew.Println("getActivity: %v", r)
}

func TestActivity(*testing.T) {
	w := getStrap()

	r, _ := w.getActivity(map[string]interface{}{"guid": "demo-guid"})
	spew.Println("getActivity: %v", r)
}

func TestReport(*testing.T) {
	w := getStrap()

	r, _ := w.getReport(map[string]interface{}{})
	spew.Println("get Report: %v", r)
}

func TestToday(*testing.T) {
	w := getStrap()

	r, _ := w.getToday(map[string]interface{}{})
	spew.Println("getToday: %v", r)
}

func TestTrigger(*testing.T) {
	w := getStrap()

	r, _ := w.getTrigger(map[string]interface{}{})
	spew.Println("getTrigger: %v", r)
}

func TestUsers(*testing.T) {
	w := getStrap()

	r, _ := w.getUsers(map[string]interface{}{})
	spew.Println("getUsers: %v", r)
}

func getStrap() *Strap {
	strap := New(token)
	strap.Discover()
	return strap
}
