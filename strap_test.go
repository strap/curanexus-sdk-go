package strap

import "testing"

const (
	token = "{ READ TOKEN for Project }"
)

func TestEndpoints(*testing.T) {
	strap := getStrap()

	endpoints := strap.endpoints()
	strap.debug.Log("End Points", endpoints)
}

func TestActivity(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "activity",
		Method: "GET",
		Params: Query{
			"guid": "user-guid",
		},
	}

	reports := []*Report{}

	// Get Activity Reports by Page
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Activity:", reports, res, err)

	// Get all Activity Reports
	res, err = strap.All(q, &reports)
	strap.debug.Log("All Activity:", reports, res, err)
}

func TestBehavior(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "behavior",
		Method: "GET",
		Params: Query{
			"guid":    "user-guid",
			"weekday": "monday",
		},
	}

	reports := []*Report{}

	// Get Activity Reports by Page
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Behavior:", reports, res, err)

}

func TestJob(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "job",
		Method: "GET",
		Params: Query{},
	}

	job := []*Job{}

	// Get All Jobs
	res, err := strap.Call(q, &job)
	strap.debug.Log("All Jobs:", job, res, err)

	q = Request{
		Name:   "job",
		Method: "GET",
		Params: Query{
			"id": "job-id",
		},
	}

	job = []*Job{}

	// Get Job by Id
	res, err = strap.Call(q, &job)
	strap.debug.Log("Jobs:", job, res, err)

}

func TestMonth(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "month",
		Method: "GET",
		Params: Query{},
	}

	reports := []*Report{}

	// Get Activity Reports by Page
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Month:", reports, res, err)

}

func TestReport(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "report",
		Method: "GET",
		Params: Query{
			"id": "some-report-id",
		},
	}

	reports := []*Report{}

	// Get single Reports by Id
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Report:", reports, res, err)
}

func TestSegmentation(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "segmentation",
		Method: "GET",
		Params: Query{},
	}

	segmentation := Segmentation{}

	// Get Project Segmentation
	res, err := strap.Call(q, &segmentation)
	strap.debug.Log("Trigger:", segmentation, res, err)

}

func TestToday(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "today",
		Method: "GET",
		Params: Query{},
	}

	reports := []*Report{}

	// Get Today's Reports by Page
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Today:", reports, res, err)

	// Get all of Today's Reports
	res, err = strap.All(q, &reports)
	strap.debug.Log("All Today:", reports, res, err)
}

func TestTrigger(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "trigger",
		Method: "GET",
		Params: Query{
			"id": "trigger-id",
		},
	}

	trigger := Trigger{}

	// Get Trigger by ID
	res, err := strap.Call(q, &trigger)
	strap.debug.Log("Trigger:", trigger, res, err)

}

func TestUser(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "user",
		Method: "GET",
		Params: Query{
			"guid": "user-guid",
		},
	}

	user := User{}

	// Get User
	res, err := strap.Call(q, &user)
	strap.debug.Log("User:", user, res, err)

}

func TestUsers(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "users",
		Method: "GET",
		Params: Query{},
	}

	users := []*User{}

	// Get All Users by Page
	res, err := strap.Call(q, &users)
	strap.debug.Log("Users:", users, res, err)

}

func TestWeek(*testing.T) {
	strap := getStrap()

	q := Request{
		Name:   "week",
		Method: "GET",
		Params: Query{},
	}

	reports := []*Report{}

	// Get Activity Reports by Page
	res, err := strap.Call(q, &reports)
	strap.debug.Log("Week:", reports, res, err)

}

func getStrap() *Strap {
	strap := New(token)
	strap.Discover()
	return strap
}
