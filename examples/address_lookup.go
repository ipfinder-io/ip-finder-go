package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := conf.GetAddressInfo("1.0.0.0")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
