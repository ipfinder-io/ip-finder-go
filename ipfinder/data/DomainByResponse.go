package data

type DomainByResponse struct {
    Status        string         `json:"status"`        
    StatusMessage string         `json:"status_message"`
    SelectBy      string         `json:"select_by"`     
    TotalDomain   int64          `json:"total_domain"`  
    ListDomain    [][]ListDomain `json:"list_domain"`   
}


type ListDomain struct {
    DomainName   string `json:"domain_name"` 
    Address      string `json:"address"`     
    Country      string `json:"country"`     
    Asn          string `json:"asn"`         
    Organization string `json:"organization"`
}