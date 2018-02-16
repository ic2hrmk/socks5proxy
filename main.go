package main

import "go-socks5"

func runServer(conf *Configuration) {
	socks5Conf := &socks5.Config{}
	server, err := socks5.New(socks5Conf)
	if err != nil {
		panic(err)
	}

	server.ListenAndServe(conf.GetTransportProtocol(), conf.GetServeAddress())
	if err != nil {
		panic(err)
	}
}

func main()  {
	conf, _ := GetConfiguration()
	runServer(conf)
}