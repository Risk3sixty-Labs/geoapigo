package main

import (
	"fmt"

	geoapigo "github.com/Risk3sixty-Labs/geoapi-go"
)

func main() {
	res1, err := geoapigo.Get("8.8.8.8")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res1)
	// {IP:8.8.8.8 Range:[134742016 134774783] Country:US Region: Eu:0 Timezone:America/Chicago City: Ll:[37.751 -97.822] Metro:0 Area:1000}
}
