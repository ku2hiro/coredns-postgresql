package coredns_postgresql

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("postgresql", setup) }

func setup(c *caddy.Controller) error {
	c.Next()
	if !c.NextArg() {
		return plugin.Error("postgresql", c.ArgErr())
	}

	if c.Val() != "dsn" {
		return plugin.Error("postgresql", c.ArgErr())
	}

	var handler Postgresql

	for c.NextBlock() {
		switch c.Val() {
		case "user":
			if !c.NextArg() {
				return plugin.Error("postgresql", c.ArgErr())
			}
			handler.user = c.Val()
		case "password":
			if !c.NextArg() {
				return plugin.Error("postgresql", c.ArgErr())
			}
			handler.password = c.Val()
		case "database":
			if !c.NextArg() {
				return plugin.Error("postgresql", c.ArgErr())
			}
			handler.database = c.Val()
		}
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return handler
	})
	return nil
}
