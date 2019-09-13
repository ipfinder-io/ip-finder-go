package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.New("", "")
	auth, err := conf.Authentication()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(auth)
	}
}
