package service

//APIHeader defines API header structure
type APIHeader struct {
	ProcessTime string `json:"process_time"`
	Message     string `json:"message"`

	ErrorCode   int    `json:"error_code"`
	ErrorType   string `json:"error_type"`
	ErrorReason string `json:"error_reason"`
}

//APIResponse defines API response structure
type APIResponse struct {
	Header  APIHeader   `json:"header"`
	Payload interface{} `json:"payload"`
}
