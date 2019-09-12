package data

type AsnResponse struct {
    Status        string      `json:"status"`        
    StatusMessage string      `json:"status_message"`
    Asn           string      `json:"asn"`           
    OrgName       string      `json:"org_name"`      
    ContinentCode string      `json:"continent_code"`
    ContinentName string      `json:"continent_name"`
    CountryCode   string      `json:"country_code"`  
    CountryName   string      `json:"country_name"`  
    Allocated     string      `json:"allocated"`     
    Registry      string      `json:"registry"`      
    Domain        string      `json:"domain"`        
    NumIPSIpv4    string      `json:"num_ips_ipv4"`  
    NumIPSIpv6    string      `json:"num_ips_ipv6"`  
    AsName        string      `json:"as_name"`       
    OrgID         string      `json:"org_id"`        
    ComanyType    string      `json:"comany_type"`   
    Speed         Speed       `json:"speed"`         
    Peers         Peers       `json:"peers"`         
    Upstreams     Upstreams   `json:"upstreams"`     
    Downstreams   Downstreams `json:"downstreams"`   
    Prefixes      Prefixes    `json:"prefixes"`      
}


type Downstreams struct {
    TotalDownstreams int64    `json:"total_downstreams"`
    ListDownstreams  []string `json:"list_downstreams"` 
}


type Peers struct {
    TotalPeers int64    `json:"total_peers"`
    ListPeers  []string `json:"list_peers"` 
}


type Prefixes struct {
    TotalPrefixes int64    `json:"total_prefixes"`
    ListPrefixes  []string `json:"list_prefixes"` 
}


type Speed struct {
    Ping     string `json:"ping"`    
    Download string `json:"download"`
    Upload   string `json:"upload"`  
}


type Upstreams struct {
    TotalUpstreams int64    `json:"total_upstreams"`
    ListUpstreams  []string `json:"list_upstreams"` 
}