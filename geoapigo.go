package geoapigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Response is the structure of a response as is
// returned by the geoapi endpoint.
type Response struct {
	IP       string
	Range    []int
	Country  string
	Region   string
	Eu       string
	Timezone string
	City     string
	Ll       []float32
	Metro    int
	Area     int
}

// Get makes a GET request to the geoapi endpoint
// and parses the response
func Get(ip string) (Response, error) {
	var r Response

	endpoint := getDefaultEnv("GEOAPI_ENDPOINT", "https://geo.risk3sixty.com")
	resp, err := http.Get(fmt.Sprintf("%s/%s", endpoint, ip))
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	return r, err
}

// getDefaultEnv gets an environment variable if present, otherwise
// uses a default fallback provided.
func getDefaultEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
