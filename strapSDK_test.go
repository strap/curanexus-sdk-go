package strapSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	token = ""
)

func TestUsers(*testing.T) {
	w := getWaiter()

	r, _ := w.Send("users", map[string]interface{}{})
	defer r.Close()
	m := []map[string]interface{}{}
	json.NewDecoder(r).Decode(&m)
	fmt.Println(m)
}

func TestActivity(*testing.T) {
	w := getWaiter()

	r, _ := w.Send("activity", map[string]interface{}{"guid": "kirk"})
	defer r.Close()
	m := []map[string]interface{}{}
	json.NewDecoder(r).Decode(&m)
	fmt.Println(m)
}

func TestToday(*testing.T) {
	w := getWaiter()

	r, _ := w.Send("today", map[string]interface{}{})
	defer r.Close()
	m := []map[string]interface{}{}
	json.NewDecoder(r).Decode(&m)
	fmt.Println(m)
}

func getWaiter() *StrapSDK {
	strap := New(token)
	strap.Discover()
	return strap
}
