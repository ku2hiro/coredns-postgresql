package coredns_postgresql

import (
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"net"
)

func newResponse(state request.Request, rdata string) dns.RR {
	var rr dns.RR
	switch state.Type() {
	case "A":
		rr = new(dns.A)
		rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
		rr.(*dns.A).A = net.ParseIP(rdata).To4()
	}
	return rr
}
