package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	range, err := conf.GetRanges("Telecom Algeria")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(range)
	}
}
