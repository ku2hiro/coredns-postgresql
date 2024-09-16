package coredns_postgresql

import (
	"context"

	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

const name = "postgresql"

// ServeDNS implements the plugin.Handler interface.
func (handler Postgresql) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	clog.Info("Start postgresql !!!")
	clog.Info(state.Name())
	clog.Info(state.Type())
	clog.Info(state.IP())
	clog.Info("End postgresql !!!")

	response := createResponse(state)
	w.WriteMsg(response)

	return 0, nil
}

// Name implements the Handler interface.
func (handler Postgresql) Name() string { return name }
