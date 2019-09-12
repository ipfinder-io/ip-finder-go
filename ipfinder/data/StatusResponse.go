package data

type StatusResponse struct {
    APIKey          string `json:"apiKey"`         
    QueriesPerDay   int64  `json:"queriesPerDay"`  
    QueriesLeft     int64  `json:"queriesLeft"`    
    AsqueriesPerDay int64  `json:"asqueriesPerDay"`
    AsqueriesLeft   int64  `json:"asqueriesLeft"`  
    KeyType         string `json:"key_type"`       
    KeyInfo         string `json:"key_info"`       
    Status          string `json:"status"`         
}