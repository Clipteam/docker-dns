package dnsserver

import (
	"fmt"
	"strings"

	"github.com/Clipteam/dockerdns/common"
	"github.com/miekg/dns"
)

func HandleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	for _, extra := range r.Extra {
		switch extra.(type) {
		case *dns.OPT:
			for _, o := range extra.(*dns.OPT).Option {
				switch e := o.(type) {
				case *dns.EDNS0_SUBNET:
					if e.Address != nil {
						common.Logger.Warnf("EDNS0 subnet not found: %s\n", e.Address.String())
					}
				}
			}
		}
	}
	m := new(dns.Msg)
	m.SetReply(r)
	m.SetEdns0(4096, true)
	m.Compress = false
	switch r.Opcode {
	case dns.OpcodeQuery:
		ParseQuery(m)
	}
	err := w.WriteMsg(m)
	if err != nil {
		return
	}
}

func ParseQuery(m *dns.Msg) {
	m.RecursionAvailable = false
	m.RecursionDesired = false
	m.Authoritative = true
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			if strings.HasSuffix(q.Name, "clipd.") {
				ipaddr, _ := PickIpAddr(q.Name)
				if ipaddr != "" {
					rr, _ := dns.NewRR(fmt.Sprintf("%s %s IN A %s", q.Name, "600", ipaddr))
					m.Answer = append(m.Answer, rr)
				} else {
					return
				}
			} else {
				return
			}
		default:
			return
		}
	}
}

func StartDnsServer() {
	common.Logger.Infof("Starting DNS server at %d", common.Config.DNSPort)
	dns.HandleFunc(common.Config.Suffix, HandleDnsRequest)
	server := &dns.Server{Addr: fmt.Sprintf(":%d", common.Config.DNSPort), Net: "udp"}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
	defer server.Shutdown()
}
