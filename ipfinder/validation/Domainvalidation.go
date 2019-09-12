package validation // import "github.com/ipfinder-io/ip-finder-go/ipfinder/validation"

import (
	"regexp"
)

// Domainvalidate Helper method for validating domain name
// passing in a single website name domain name
func Domainvalidate(domain string) bool {

	match, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\\.[a-zA-Z]{2,}$", domain)

	return match

}
