package validation // import "github.com/ipfinder-io/ip-finder-go/ipfinder/validation"

import (
	"regexp"
)

// Asnvalidate Helper method for validating an argument if it is asn number
// the AS number you want details for
func Asnvalidate(asn string) bool {

	match, _ := regexp.MatchString("^(as|AS)(\\d+)$", asn)

	return match

}
