package main

import (
	"fmt"
	"update_dns/ip"
)

func main() {
	var myIp *ip.IP
	myIp = ip.New(10, func() {
		fmt.Println(myIp.GetPublicIp())
		fmt.Println("hello world")
	})
	myIp.Start()

	fmt.Printf("Hello World")
}
