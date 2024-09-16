package coredns_postgresql

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("postgresql", setup) }

func setup(c *caddy.Controller) error {
	c.Next()
	if c.NextArg() {
		return plugin.Error("postgresql", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Postgresql{
			user:     "coredns",
			password: "coredns",
			database: "coredns",
		}
	})
	return nil
}
