package coredns_postgresql

import (
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"net"
)

func createResponse(state request.Request) *dns.Msg {
	response := new(dns.Msg)
	response.SetReply(state.Req)
	response.Authoritative = true

	switch state.Type() {
	case "A":
		rdata := "1.1.1.1"
		var rr dns.RR = new(dns.A)
		rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
		rr.(*dns.A).A = net.ParseIP(rdata).To4()
		response.Answer = []dns.RR{rr}
	}
	return response
}
