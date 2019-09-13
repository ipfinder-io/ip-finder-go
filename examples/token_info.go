package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.New("YOUR_TOKEN_GOES_HERE", "")
	status, err := conf.GetStatus()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(status)
	}
}
