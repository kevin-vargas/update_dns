package ip

import (
	"bytes"
	"log"
	"os/exec"
	"sync"
	"time"
)

const (
	PID = "dig"
)

var (
	ARGS = []string{"+short", "myip.opendns.com", "@resolver1.opendns.com"}
)

type IP struct {
	sync.Mutex
	public     []byte
	notify     func()
	interval   int
	cancelChan *chan bool
}

func New(interval int, notify func()) *IP {
	cancelChan := make(chan bool, 1)
	ip, err := getPublicIp()
	if err != nil {
		log.Fatal("No connection, can't get ip address")
	}
	return &IP{
		public:     ip,
		interval:   interval,
		notify:     notify,
		cancelChan: &cancelChan,
	}
}

func (ip *IP) Start() {
L:
	for {
		select {
		case <-*ip.cancelChan:
			break L
		case <-time.After(time.Duration(ip.interval) * time.Second):
			currentPublicIp, err := getPublicIp()
			if err == nil {
				log.Printf("Current Public Ip: %s\n", currentPublicIp)
				ip.Lock()
				res := bytes.Compare(ip.public, currentPublicIp)
				if res != 0 {
					log.Println("Public Ip Change making notification")
					ip.public = currentPublicIp
					ip.notify()
				}
				ip.Unlock()
			}
		}
	}
}

func (ip *IP) GetPublicIp() string {
	currentIp := ip.public
	return string(currentIp)
}

func getPublicIp() ([]byte, error) {
	ip, err := exec.Command(PID, ARGS...).Output()
	if err != nil {
		return nil, err
	}
	return ip, nil
}
