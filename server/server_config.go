package server

// ServerConfig contains configuration settings for running
// HTTP server
type ServerConfig struct {
	Port int    `xml:"port" json:"port"`
	Key  string `xml:"key" 	json:"key"`
}