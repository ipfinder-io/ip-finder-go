package main

import (
	"fmt"

	"github.com/ipfinder-io/ip-finder-go/ipfinder"
)

func main() {
	conf := ipfinder.New("YOUR_TOKEN_GOES_HERE", "")
	fire, err := conf.GetFirewall("as1", "nginx_deny")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(fire)
	}
}
