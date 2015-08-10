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

*Checkout the strap_test.go file for all SDK endpoint examples*

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

// List available endpoints
r := strap.endpoints()
// No Params

// Example usage of Strap
// Checkout the strap_test.go file for more information
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

```
