package main

import (
	"fmt"
)

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetStatus()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
