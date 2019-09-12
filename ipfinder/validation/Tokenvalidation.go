package validation


func tokenvalidate(token string)(bool) {

    if len(token) <= 25 {
        return false
    } else {
        return true
    }

}

