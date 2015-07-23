package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
)

func Test_LoadXMLConfig_ShouldLoadValuesFromXMLFile(t *testing.T) {
	configContent := "<config><port>1234</port></config>"
	createConfigFile("config.xml", configContent)
	server := Server{}
	server.LoadXMLConfig("config.xml")
	assert.Equal(t, 1234, server.Config.Port)
	deleteConfigFile("config.xml")
}

func Test_LoadXMLConfig_IfConfigDoesntContainValuesTheDefaultValuesShouldBeSet(t *testing.T) {
	configContent := "<config></config>"
	createConfigFile("config.xml", configContent)
	server := Server{}
	server.LoadXMLConfig("config.xml")
	assert.Equal(t, 0, server.Config.Port)
	deleteConfigFile("config.xml")
}

func Test_LoadJSONConfig_ShouldLoadValuesFromJSONFile(t *testing.T) {
	configContent := "{\"port\":1234}"
	createConfigFile("config.json", configContent)
	server := Server{}
	server.LoadJSONConfig("config.json")
	assert.Equal(t, 1234, server.Config.Port)
	deleteConfigFile("config.json")
}

func Test_LoadJSONConfig_IfConfigDoesntContainValuesTheDefaultValuesShouldBeSet(t *testing.T) {
	configContent := "{}"
	createConfigFile("config.json", configContent)
	server := Server{}
	server.LoadJSONConfig("config.json")
	assert.Equal(t, 0, server.Config.Port)
	deleteConfigFile("config.json")
}

func Test_LoadYAMLConfig_ShouldLoadValuesFromJSONConfig(t *testing.T) {
	configContent := "port: 1234"
	createConfigFile("config.yaml", configContent)
	server := Server{}
	server.LoadYAMLConfig("config.yaml")
	assert.Equal(t, 1234, server.Config.Port)
	deleteConfigFile("config.yaml")
}

func Test_LoadYAMLConfig_IfConfigDoesntContainValuesTheDefaultValuesShouldBeSet(t *testing.T) {
	configContent := ""
	createConfigFile("config.yaml", configContent)
	server := Server{}
	server.LoadYAMLConfig("config.yaml")
	assert.Equal(t, 0, server.Config.Port)
	deleteConfigFile("config.yaml")
}

func Test_readConfig_ShouldReturnFileContent(t *testing.T) {
	server := Server{}
	fileContent := "file content string"
	createConfigFile("config.cfg", fileContent)
	actualFileContent := server.readConfig("config.cfg")
	assert.Equal(t, []byte("file content string"), actualFileContent)
	deleteConfigFile("config.cfg")
}

func Test_port_ShouldReturnCorrectPort(t *testing.T) {
	server := Server{
		Config: ServerConfig{
			Port: 1234,
		},
	}
	assert.Equal(t, ":1234", server.port())
}

func createConfigFile(filename string, content string) {
	bytesContent := []byte(content)
	err := ioutil.WriteFile(filename, bytesContent, 0777)
	if err != nil {
		logger.Fatal(err)
	}
}

func deleteConfigFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		logger.Fatal(err)
	}
}