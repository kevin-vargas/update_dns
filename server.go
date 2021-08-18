package main

import (
	"fmt"
	"log"
	"update_dns/ip"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var myIp *ip.IP
	myIp = ip.New(10, func() {
		fmt.Println(myIp.GetPublicIp())
		fmt.Println("hello world")
	})
	myIp.Start()
}
