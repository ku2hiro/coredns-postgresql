package coredns_postgresql

import (
	"context"
	"fmt"
	"net"

	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

const name = "postgresql"

type Postgresql struct{}

// ServeDNS implements the plugin.Handler interface.
func (wh Postgresql) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	clog.Info("Start postgresql !!!")
	clog.Info(state.Name())
	clog.Info(state.Type())
	clog.Info(state.IP())
	clog.Info("End postgresql !!!")

	a := new(dns.Msg)
	a.SetReply(r)
	a.Authoritative = true

	ip := "1.1.1.1"

	var rr dns.RR = new(dns.A)
	rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
	rr.(*dns.A).A = net.ParseIP(ip).To4()

	a.Extra = []dns.RR{rr}
	a.Answer = []dns.RR{rr}

	w.WriteMsg(a)

	return 0, nil
}

// Name implements the Handler interface.
func (wh Postgresql) Name() string { return name }
