package main

import (
	"github.com/Clipteam/dockerdns/common"
	"github.com/Clipteam/dockerdns/dnsserver"
)

func main() {
	common.InitConfig()
	common.InitLog()
	dnsserver.StartDnsServer()
}
