package net

// LocalAddr is a list of private network address ranges, i.e. reserved addresses of local usage.
// TODO consider separating and naming ranges, then joining in slice by variables.
var LocalAddrs = []string{
	"0.0.0.0",
	"127/8",
	"10/8",
	"172.16/12",
	"192.168/16",
	"169.254/16",
}
