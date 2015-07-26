package main

import "github.com/doojin/file-explorer/server"

func main() {
	server := new(server.Server)
	server.LoadXMLConfig("conf.xml")
	server.Start()
}


