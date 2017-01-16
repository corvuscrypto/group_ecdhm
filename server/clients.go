package server

import (
	"net"

	"github.com/corvuscrypto/group_ecdhm"
)

//Client representation for the server to manage
type Client struct {
	Name       string
	Connection *net.TCPConn
	sharedKey  gecdhm.Point
}

//This will hold our clients
var clientMap map[string]*Client

func init() {
	//Allocate the make
	clientMap = make(map[string]*Client)
}
