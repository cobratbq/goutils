package net

import "net"

// LocalAddr is a list of private network address ranges, i.e. reserved addresses of local usage.
// TODO consider separating and naming ranges, then joining in slice by variables.
var PrivateNetworks = []*net.IPNet{
	{IP: []byte{0, 0, 0, 0}, Mask: []byte{255, 255, 255, 255}},
	{IP: []byte{127, 0, 0, 0}, Mask: []byte{255, 0, 0, 0}},
	{IP: []byte{10, 0, 0, 0}, Mask: []byte{255, 0, 0, 0}},
	{IP: []byte{172, 16, 0, 0}, Mask: []byte{255, 240, 0, 0}},
	{IP: []byte{192, 168, 0, 0}, Mask: []byte{255, 255, 0, 0}},
	{IP: []byte{169, 254, 0, 0}, Mask: []byte{255, 255, 0, 0}},
}
