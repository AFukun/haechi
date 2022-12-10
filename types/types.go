package types

import (
	"bytes"
	"net"
	"strconv"
)

type HaechiAddress struct {
	Ip   net.IP
	Port uint16
}

func BytesToIp(bt []byte) net.IP {
	ip_parts := bytes.Split(bt, []byte("."))
	temp_p1, _ := strconv.ParseUint(string(ip_parts[0]), 10, 64)
	p1 := uint8(temp_p1)
	temp_p2, _ := strconv.ParseUint(string(ip_parts[1]), 10, 64)
	p2 := uint8(temp_p2)
	temp_p3, _ := strconv.ParseUint(string(ip_parts[2]), 10, 64)
	p3 := uint8(temp_p3)
	temp_p4, _ := strconv.ParseUint(string(ip_parts[3]), 10, 64)
	p4 := uint8(temp_p4)
	return []byte{p1, p2, p3, p4}
}
