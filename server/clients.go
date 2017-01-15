package server

import "net"

//Client representation for the server to manage
type Client struct {
	Name       string
	Connection *net.TCPConn
	verified   bool
}

//This will hold our clients
var clientMap map[string]*Client

func init() {
	//Allocate the make
	clientMap = make(map[string]*Client)
}
