package data

type DomainHTResponse struct {
    Status        string         `json:"status"`        
    StatusMessage string         `json:"status_message"`
    TotalDomain   int64          `json:"total_domain"`  
    ListDomain    [][]ListDomainht `json:"list_domain"`   
}


type ListDomainht struct {
    DomainName   string `json:"domain_name"` 
    Address      string `json:"address"`     
    Country      string `json:"country"`     
    Asn          string `json:"asn"`         
    Organization string `json:"organization"`
    LastSeenOn   string `json:"last_seen_on"`
}