// package validation
package validation

import (
    "regexp"
)
func ipvalidate(ip string)(bool) {

    match, _ := regexp.MatchString("^(as|AS)(\\d+)$", ip)
    if match == false {
        return false
    } else {
        return true
    }

}