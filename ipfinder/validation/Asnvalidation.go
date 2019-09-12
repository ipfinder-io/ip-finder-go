package validation

import (
    "regexp"
)

 // Asnvalidate Helper method for validating an argument if it is asn number
func Asnvalidate(asn string)(bool) {

    match, _ := regexp.MatchString("^(as|AS)(\\d+)$", asn)
    if match == false {
        return false
    } 
    return true
    
}
