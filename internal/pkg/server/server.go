package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Server is the server struct for a json server
type Server struct {
	URL             string `json:"url"`
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	Distance        int    `json:"distance"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	Cc              string `json:"cc"`
	Sponsor         string `json:"sponsor"`
	ID              string `json:"id"`
	Preferred       int    `json:"preferred"`
	Host            string `json:"host"`
	ForcePingSelect int    `json:"force_ping_select"`
}

// GetAllServers parses the list of all recommended servers
func GetAllServers() (servers []Server) {
	res, err := http.Get("http://www.speedtest.net/api/js/servers?engine=js")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &servers)
	if err != nil {
		panic(err.Error())
	}

	return servers
}

// GetBestServer returns the first servert of GetAllServers
func GetBestServer() (s Server) {
	// TODO right now we are just picking the first server
	// eventually we need a better algorithm
	servers := GetAllServers()
	return servers[0]
}