package main

import (
	"flag"
	"fmt"
)

const (
	DefaultAddress           = "127.0.0.1"
	DefaultPort              = "8000"
	DefaultTransportProtocol = "tcp"
)

type Configuration struct {
	address  string
	port     string
	protocol string
}

var conf *Configuration

func init() {
	conf = &Configuration{}

	flag.StringVar(&conf.address, "address", DefaultAddress, "Address for host")
	flag.StringVar(&conf.port, "port", DefaultPort, "Proxy port")
	flag.Parse()

	conf.protocol = DefaultTransportProtocol
}

func (conf Configuration) GetServeAddress() string {
	return fmt.Sprintf("%s:%s", conf.address, conf.port)
}

func (conf Configuration) GetTransportProtocol() string {
	return conf.protocol
}

func GetConfiguration() (*Configuration, error) {
	return conf, nil
}
