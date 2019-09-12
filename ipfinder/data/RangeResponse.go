package data

type RangeResponse struct {
    Status        string       `json:"status"`        
    StatusMessage string       `json:"status_message"`
    OrgName       string       `json:"org_name"`      
    ContinentCode string       `json:"continent_code"`
    ContinentName string       `json:"continent_name"`
    CountryName   string       `json:"country_name"`  
    Domain        string       `json:"domain"`        
    NumRanges     string       `json:"num_ranges"`    
    NumIpv4       string       `json:"num_ipv4"`      
    NumIpv6       string       `json:"num_ipv6"`      
    NumAsn        int64        `json:"num_asn"`       
    ListAsn       []ListAsn    `json:"list_asn"`      
    ListPrefixes  []ListPrefix `json:"list_prefixes"` 
}


type ListAsn struct {
    Asn         string `json:"asn"`         
    OrgID       string `json:"OrgId"`       
    TotalPrefix string `json:"Total_prefix"`
    TotalV4     string `json:"Total_v4"`    
    TotalV6     string `json:"Total_v6"`    
}


type ListPrefix struct {
    Asn    string `json:"asn"`   
    Prefix string `json:"prefix"`
}