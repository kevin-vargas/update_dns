package main

import (
	"log"
	"update_dns/cloudflare"
	"update_dns/constants"
	"update_dns/ip"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var cf *cloudflare.CF = cloudflare.GetInstance()
	var myIp *ip.IP
	myIp = ip.New(constants.INTERVAL_NOTIFY, func() {
		cf.Update(myIp.GetPublicIp())
	})
	myIp.Start()
}
