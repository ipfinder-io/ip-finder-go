package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	domain, err := conf.GetDomainHistory("google.dz")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(domain)
	}
}
