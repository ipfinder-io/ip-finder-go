package validation


import (
    "regexp"
)

func domainvalidate(domain string)(bool) {

    match, _ := regexp.MatchString("^(?!\\-)(?:[a-zA-Z\\d\\-]{0,62}[a-zA-Z\\d]\\.){1,126}(?!\\d+)[a-zA-Z\\d]{1,63}$", domain)
    if match == false {
        return false
    } else {
        return true
    }

}

