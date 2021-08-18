package cloudflare

import (
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *CF
)

type CF struct {
	zone      string
	dnsRecord string
	apikey    string
}

func new() *CF {
	zone := os.Getenv("ZONE")
	dnsRecord := os.Getenv("DNS_RECORD")
	apikey := os.Getenv("API_KEY")
	return &CF{
		zone:      zone,
		dnsRecord: dnsRecord,
		apikey:    apikey,
	}
}

func GetInstance() *CF {
	once.Do(func() {
		instance = new()
	})
	return instance
}
