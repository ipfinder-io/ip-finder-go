package ipfinder

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	// "encoding/json"
	// "log"
)

// IPFinder init
type IPFinder struct {
	token   string
	baseURL string
}

const (

	// DEFAULT BASE URL.

	DEFAULTBASEURL = "https://api.ipfinder.io/v1/"

	// The free token.

	DEFAULTAPITOKEN = "free"

	// DEFAULT FORMAT.

	FORMAT = "json"

	// service status path

	STATUSPATH = "info"
	// IP Address Ranges path

	RANGESPATH = "ranges/"

	// Firewall path

	FIREWALLPATH = "firewall/"

	// Get Domain IP path

	DOMAINPATH = "domain/"

	// Get Domain IP history path

	DOMAINHPATH = "domainhistory/"

	// Domain By ASN, Country,Ranges path

	DOMAINBYPATH = "domainby/"
)

// NewIPFinder Constructor
// ADD YOUR TOKEN AND YOUR PROXY
func NewIPFinder(t string, b string) *IPFinder {

	p := new(IPFinder)
	if t == "" {
		p.token = DEFAULTAPITOKEN
	} else {
		p.token = t
	}
	if b == "" {
		p.baseURL = DEFAULTBASEURL
	} else {
		p.baseURL = b
	}
	return p
}

// Call to server
// get data
func (p *IPFinder) Call(path string, format string) (string, error) {

	var jsonStr = []byte(`{"token":"` + p.token + `", "format":"` + format + `"}`)
	req, err := http.Post(p.baseURL+path, "application/json", bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", "IPfinder go-client")

	if err != nil {
		return "", errors.New("Error")
	}
	body, _ := ioutil.ReadAll(req.Body)
	if req.StatusCode != 200 {
		return "", errors.New("\nURL:" + p.baseURL + path + " => response response:\n" + string(body))
	}
	return string(body), err
}

// Authentication Get details for an Your IP address.
// Your IP address data.
func (p *IPFinder) Authentication() (string, error) {
	return p.Call("", "FORMAT")
}

// GetAddressInfo Get details for an IP address.
// IP address data.
func (p *IPFinder) GetAddressInfo(path string) (string, error) {
	return p.Call(path, FORMAT)
}

// GetAsn Get details for an AS number.
// AS number data.
func (p *IPFinder) GetAsn(path string) (string, error) {
	return p.Call(path, FORMAT)
}

// GetStatus Get details for an API Token .
// The Token data.
func (p *IPFinder) GetStatus() (string, error) {
	return p.Call(STATUSPATH, FORMAT)
}

// GetRanges Get details for an Organization name.
// Organization name data.
func (p *IPFinder) GetRanges(path string) (string, error) {
	return p.Call(RANGESPATH+url.QueryEscape(path), FORMAT)
}

// GetFirewall Get Firewall data
// Firewall data.
func (p *IPFinder) GetFirewall(by string, formats string) (string, error) {
	return p.Call(FIREWALLPATH+by, formats)
}

// GetDomain Get Domain IP
// Domain to IP data.
func (p *IPFinder) GetDomain(path string) (string, error) {
	return p.Call(DOMAINPATH+path, FORMAT)
}

// GetDomainHistory Get Domain History IP
// Domain History data.
func (p *IPFinder) GetDomainHistory(path string) (string, error) {
	return p.Call(DOMAINHPATH+path, FORMAT)
}

// GetDomainBy Get list Domain By ASN, Country,Ranges
// list Domain By ASN, Country,Ranges data.
func (p *IPFinder) GetDomainBy(by string) (string, error) {
	return p.Call(DOMAINBYPATH+by, FORMAT)
}
