package data

type DomainResponse struct {
    Status            string `json:"status"`             
    StatusMessage     string `json:"status_message"`     
    IP                string `json:"ip"`                 
    Domain            string `json:"domain"`             
    DomainStatus      bool   `json:"domain_status"`      
    ContinentCode     string `json:"continent_code"`     
    ContinentName     string `json:"continent_name"`     
    CountryCode       string `json:"country_code"`       
    CountryName       string `json:"country_name"`       
    CountryNativeName string `json:"country_native_name"`
    RegionName        string `json:"region_name"`        
    City              string `json:"city"`               
    Latitude          string `json:"latitude"`           
    Longitude         string `json:"longitude"`          
    Asn               string `json:"asn"`                
    Organization      string `json:"organization"`       
}