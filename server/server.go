// Package server implements a simple HTTP server for displaying
// user interface for the file explorer service
package server

import (
	"io/ioutil"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/op/go-logging"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

var logger = logging.MustGetLogger("HTTP Server")

// A simple HTTP server
type Server struct {
	Config ServerConfig
}

// Start runs server with configuration from server config
func (server *Server) Start() {
	r := mux.NewRouter()
	http.Handle("/", r)
	logger.Info("HTTP server is starting using port: %v", server.Config.Port)
	http.ListenAndServe(server.port(), nil)
}

// LoadXMLConfig fills server config with values from XML file
func (server *Server) LoadXMLConfig(filename string) {
	fileContent := server.readConfig(filename)
	xml.Unmarshal(fileContent, &server.Config)
}

// LoadJSONConfig fills server config with values from JSON file
func (server *Server) LoadJSONConfig(filename string) {
	fileContent := server.readConfig(filename)
	json.Unmarshal(fileContent, &server.Config)
}

// LoadYAMLConfig fills server config with values from YAML file
func (server *Server) LoadYAMLConfig(filename string) {
	fileContent := server.readConfig(filename)
	yaml.Unmarshal(fileContent, &server.Config)
}

func (server *Server) port() string {
	return ":" + strconv.Itoa(server.Config.Port)
}

func (server *Server) readConfig(filename string) []byte {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Fatalf("Cannot read file %v: %v", filename, err)
	}
	return fileContent
}