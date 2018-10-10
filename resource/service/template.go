package service

//APIHeader defines API header structure
type APIHeader struct {
	ProcessTime string `json:"process_time"`
	Message     string `json:"message"`
	Reason      string `json:"reason"`
	ErrorCode   string `json:"error_code"`
}

//APIResponse defines API response structure
type APIResponse struct {
	Header  APIHeader   `json:"header"`
	Payload interface{} `json:"payload"`
}
