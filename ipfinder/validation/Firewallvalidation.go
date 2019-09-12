package validation

import (
    "regexp"
)

func firewallvalidate(by string,format string)(bool) {

    matchby, _ := regexp.MatchString("^(as|AS)(\\d+)$", by)
    match, _   := regexp.MatchString("^(as|AS)(\\d+)$", format)
    if matchby == false && match == false {
        return false
    } else {
        return true
    }

}
