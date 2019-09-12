package main

import (
	"fmt"
)

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	auth, err := ipfinder.Authentication()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(auth)
	}
}
