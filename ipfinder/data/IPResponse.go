package data

type IPResponse struct {
    Status            string     `json:"status"`             
    StatusMessage     string     `json:"status_message"`     
    IP                string     `json:"ip"`                 
    Type              string     `json:"type"`               
    ContinentCode     string     `json:"continent_code"`     
    ContinentName     string     `json:"continent_name"`     
    CountryCode       string     `json:"country_code"`       
    CountryName       string     `json:"country_name"`       
    CountryNativeName string     `json:"country_native_name"`
    RegionName        string     `json:"region_name"`        
    City              string     `json:"city"`               
    Latitude          string     `json:"latitude"`           
    Longitude         string     `json:"longitude"`          
    Location          Location   `json:"location"`           
    TimeZone          TimeZone   `json:"time_zone"`          
    Currency          Currency   `json:"currency"`           
    Languages         Languages  `json:"languages"`          
    Connection        Connection `json:"connection"`         
    Security          Security   `json:"security"`           
    Header            Header     `json:"header"`             
}


type Connection struct {
    Asn          string `json:"asn"`         
    Organization string `json:"organization"`
    Domain       string `json:"domain"`      
    Type         string `json:"type"`        
}


type Currency struct {
    Name         string `json:"name"`         
    Symbol       string `json:"symbol"`       
    SymbolNative string `json:"symbol_native"`
}


type Header struct {
    TotalUserAgent int64           `json:"total_user_agent"`
    ListUserAgent  []ListUserAgent `json:"list_user_agent"` 
}


type ListUserAgent struct {
    UserAgent   string  `json:"user_agent"`  
    BrowserName *string `json:"browser_name"`
    OSName      *string `json:"os_name"`     
    DeviceType  string  `json:"device_type"` 
}


type Languages struct {
    Code       string `json:"code"`       
    Name       string `json:"name"`       
    NameNative string `json:"name_native"`
}


type Location struct {
    Capital                 string `json:"capital"`                   
    CountryFlag             string `json:"country_flag"`              
    FlagPNG                 string `json:"flag_png"`                  
    CountryFlagEmoji        string `json:"country_flag_emoji"`        
    CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
    CallingCode             string `json:"calling_code"`              
    ToplevelDomain          string `json:"toplevel_domain"`           
    Alpha3Code              string `json:"alpha3_code"`               
    Population              string `json:"population"`                
    IsEu                    bool   `json:"is_eu"`                     
    IsAr                    bool   `json:"is_ar"`                     
}


type Security struct {
    IsProxy     bool   `json:"is_proxy"`    
    ProxyType   bool   `json:"proxy_type"`  
    IsTor       bool   `json:"is_tor"`      
    IsSpam      bool   `json:"is_spam"`     
    ThreatLevel string `json:"threat_level"`
}


type TimeZone struct {
    ID          string `json:"id"`          
    UTC         string `json:"UTC"`         
    GmtOffset   int64  `json:"gmt_offset"`  
    CurrentTime string `json:"current_time"`
}