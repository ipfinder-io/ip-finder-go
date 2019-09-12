<img src='https://camo.githubusercontent.com/46886c3e689a0d4a3f6c0733d1cab5d9f9a3926d/68747470733a2f2f697066696e6465722e696f2f6173736574732f696d616765732f6c6f676f732f6c6f676f2e706e67' height='60' alt='IP Finder'></a>
#  IPFinder Go Client Library

The official Go client library for the [IPFinder.io](https://ipfinder.io) get details for :
-  IP address details (city, region, country, postal code, latitude and more ..)
-  ASN details (Organization name, registry,domain,comany_type, and more .. )
-  Firewall by supported formats details (apache_allow,  nginx_deny, CIDR , and more ..)
-  IP Address Ranges by the Organization name  details (list_asn, list_prefixes , and more ..)
-  service status details (queriesPerDay, queriesLeft, key_type, key_info)
- Get Domain IP (asn, organization,country_code ....)
- Get Domain IP history (total_ip, list_ip,organization,asn ....)
- Get list Domain By ASN, Country,Ranges (select_by , total_domain  , list_domain ....)

## Getting Started
singup for a free account at [https://ipfinder.io/auth/signup](https://ipfinder.io/auth/signup), for Free IPFinder API access token.

The free plan is limited to 4,000 requests a day, and doesn't include some of the data fields
To enable all the data fields and additional request volumes see [https://ipfinder.io/pricing](https://ipfinder.io/pricing).

## Documentation

See the [official documentation](https://ipfinder.io/docs).

## Installation
```shell
go get github.com/ipfinder-io/ip-finder-go
```

## With `free` TOKEN

```go 
// lookup your IP address information
func main() {
	ipfinder := NewIPFinder("", "") // "" == free token or use "free"
	auth, err := ipfinder.Authentication()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(auth)
	}
}
```

## Authentication

```go 
// lookup your IP address information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	auth, err := ipfinder.Authentication()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(auth)
	}
}
```

## Get IP address

```go 
// lookup IP address information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetAddressInfo("1.0.0.0")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}

```

## Get ASN
This API available as part of our Pro and enterprise [https://ipfinder.io/pricing](https://ipfinder.io/pricing).

```go 
// lookup Asn information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetAsn("as1")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Firewall
This API available as part of our  enterprise [https://ipfinder.io/pricing](https://ipfinder.io/pricing).
formats supported are :  `apache_allow`, `apache_deny`,`nginx_allow`,`nginx_deny`, `CIDR`, `linux_iptables`, `netmask`, `inverse_netmask`, `web_config_allow `, `web_config_deny`, `cisco_acl`, `peer_guardian_2`, `network_object`, `cisco_bit_bucket`, `juniper_junos`, `microtik`

```go 
// lookup Asn Firewall information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetFirewall("as1", "nginx_deny")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Get IP Address Ranges
This API available as part of our  enterprise [https://ipfinder.io/pricing](https://ipfinder.io/pricing).

```go 
// lookup Organization information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetRanges("Telecom Algeria")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Get service status

```go 
// lookup IP TOKEN information
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetStatus()

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Get Domain IP

```go 
// Get domain name
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "")
	ip, err := ipfinder.GetDomain("google.dz")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Get Domain IP history

```go 
// Get domain name IP list history
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetDomainHistory("google.dz")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Get list Domain By ASN, Country,Ranges

```go 
// list live domain by country DZ,US,TN,FR,MA
func main() {
	ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "") 
	ip, err := ipfinder.GetDomainBy("DZ")

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(ip)
	}
}
```

## Add proxy
```go 
ipfinder := NewIPFinder("YOUR_TOKEN_GOES_HERE", "https://ipfinder.yourdomain.com")
```

Sample codes under **examples/** folder.


## Contact

Contact Us With Additional Questions About Our API, if you would like more information about our API that isn’t available in our IP geolocation API developer documentation, simply [contact](https://ipfinder.io/contact) us at any time and we’ll be able to help you find what you need.

## License
----

[![GitHub license](https://img.shields.io/github/license/ipfinder-io/ip-finder-node.svg)](https://github.com/ipfinder-io/ip-finder-node)