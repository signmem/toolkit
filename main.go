package main

import (
	"github.com/signmem/toolkit/net"
	"fmt"
)

func test(){
	ipaddr, err := net.GetLinuxIPaddress()
	if err != nil {
		fmt.Println("can not get ipaddr")
	} else {
		fmt.Println(ipaddr)
	}
}

func main() {
	test()
}
