# strapSDK

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
git clone git+ssh://git@github.com:strap/strap-sdk-go.git
```

### Usage

Below is a basic use case.

```javascript
// Setup strapSDK, passing in the Read Token for the Project
func getStrapSDK() *StrapSDK {
	strap := New("token value")	// 
	strap.Discover()
	return strap
}

// Listen for ready before interacting with strapSDK
func TestUsers(*testing.T) {
	s := getStrapSDK()

	r, _ := s.Send("users", map[string]interface{}{})
	defer r.Close()
	m := []map[string]interface{}{}
	json.NewDecoder(r).Decode(&m)
	fmt.Println(m)
}
```
