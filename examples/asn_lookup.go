package main

import (
	"fmt"
)

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetAsn("as1")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
