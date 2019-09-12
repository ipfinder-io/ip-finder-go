package validation // import "github.com/ipfinder-io/ip-finder-go/ipfinder/validation"

// Tokenvalidate Helper method for validating IPFINDER API TOKEN
// Your API key or the string "free" for the free API
func Tokenvalidate(token string) bool {

	if len(token) <= 25 {
		return false
	}
	return true

}
