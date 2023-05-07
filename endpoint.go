package sshtunnel

import (
	"net"
	"strconv"
)

type Endpoint struct {
	Host string
	Port int
	User string
}

func (endpoint *Endpoint) String() string {
	return net.JoinHostPort(endpoint.Host, strconv.Itoa(endpoint.Port))
}
