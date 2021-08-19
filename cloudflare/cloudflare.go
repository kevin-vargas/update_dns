package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *CF
)

type CF struct {
	Zone      string
	DnsRecord string
	Apikey    string
}

func new() *CF {
	zone := os.Getenv("ZONE")
	dnsRecord := os.Getenv("DNS_RECORD")
	apikey := os.Getenv("API_KEY")
	return &CF{
		Zone:      zone,
		DnsRecord: dnsRecord,
		Apikey:    apikey,
	}
}

func GetInstance() *CF {
	once.Do(func() {
		instance = new()
	})
	return instance
}

func (cf *CF) Update(ip string) {
	var url strings.Builder
	url.WriteString("https://api.cloudflare.com/client/v4/zones/")
	url.WriteString(cf.Zone)
	url.WriteString("/dns_records/")
	url.WriteString(cf.DnsRecord)

	var bearer strings.Builder
	bearer.WriteString("Bearer ")
	bearer.WriteString(cf.Apikey)
	requestMap := map[string]interface{}{
		"name":    "pi.fast.ar",
		"type":    "A",
		"content": ip,
		"proxied": false,
	}
	requestBody, err := json.Marshal(requestMap)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("PUT", url.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", bearer.String())
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	log.Println("Updating Ip for Dns Record")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println([]byte(body))
}
