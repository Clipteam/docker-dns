package main

import (
	"github.com/sparrowhe/dockerdns/common"
	"github.com/sparrowhe/dockerdns/dnsserver"
)

func main() {
	common.InitConfig()
	common.InitLog()
	dnsserver.StartDnsServer()
}
