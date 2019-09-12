package main

import (
	"fmt"
)

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetDomainBy("DZ")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
