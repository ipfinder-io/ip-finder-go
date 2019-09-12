package main

import {
	"fmt"
	"github.com/ipfinder-io/ip-finder-go"
}

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetAsn("as1")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}