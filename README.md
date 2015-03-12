# Go > Strap Server-Side SDK

Strap SDK Go provides an easy to use, chainable API for interacting with our
API services.  Its purpose is to abstract away resource information from
our primary API, i.e. not having to manually track API information for
your custom API endpoint.

Strap SDK Go keys off of a global API discovery object using the read token for the API. 
The Strap SDK Go extracts the need for developers to know, manage, and integrate the API endpoints.

The a Project API discovery can be found here:

HEADERS: "X-Auth-Token": 
GET [https://api2.straphq.com/discover]([https://api2.straphq.com/discover)

Once the above has been fetched, `strapSDK` will fetch the API discover
endpoint for the project and build its API.

### Installation

```
git clone git@github.com:strap/strap-sdk-go.git
```

### Usage

Below is a basic use case.

```golang
// Setup strap, passing in the Read Token for the Project
func getStrap() *Strap {
	strap := New(token)
	strap.Discover()
	return strap
}

// Example usage of Strap
// Checkout the strap_test.go file for more information
func TestActivity(*testing.T) {
	strap := getStrap()

	r, _ := strap.getActivity(map[string]interface{}{"guid": "demo-guid"})
	// etc...
}

// List available endpoints
r := strap.endpoints()
// No Params

// Fetch a user's activity
// URL resource: "guid"
// Optional: "day", "count"
r, _ := strap.getActivity(map[string]interface{}{"guid": "brian-strap"})

// Fetch a report's data
// URL resource: "id"
// Optional: none
r, _ := strap.getReport(map[string]interface{}{})

// Fetch all user data for today
// URL resource: none
// Optional: "guid", "page"
r, _ := strap.getToday(map[string]interface{}{})

// Fetch trigger data
// URL resource: "id"
// Optional: "count"
r, _ := strap.getTrigger(map[string]interface{}{})

// Fetch a user list for the Project
// URL resource: none
// Optional: "platform", "count"
r, _ := strap.getUsers(map[string]interface{}{})

```
