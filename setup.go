package coredns_postgresql

import (
	"database/sql"
	"fmt"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	_ "github.com/lib/pq"
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

func (handler Postgresql) dbConnect() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		handler.user, handler.password, handler.database)
	return sql.Open("postgres", dsn)
}
