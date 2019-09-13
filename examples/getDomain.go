package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.New("YOUR_TOKEN_GOES_HERE", "")
	domain, err := conf.GetDomain("google.dz")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(domain)
	}
}
