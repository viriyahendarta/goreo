package config

//Config holds all config
type Config struct {
	HTTPServer   HTTPServer `json:"http_server"`
	CoreDatabase Database   `json:"core_database"`
}

//HTTPServer holds data needed for serving http
type HTTPServer struct {
	Port int `json:"port"`
}

//Database holds data needed for connect to database
type Database struct {
	URL    string `json:"url"`
	Driver string `json:"driver"`
}
