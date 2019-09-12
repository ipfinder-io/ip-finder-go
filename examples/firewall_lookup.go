package main

import {
	"fmt"
	"github.com/ipfinder-io/ip-finder-go"
}

func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetFirewall("as1", "nginx_deny")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}